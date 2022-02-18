package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/piigyy/sharing-is-caring/config"
	"github.com/piigyy/sharing-is-caring/internal/payment/model"
	"github.com/piigyy/sharing-is-caring/internal/payment/proto"
	"github.com/piigyy/sharing-is-caring/internal/payment/repository"
	paymentService "github.com/piigyy/sharing-is-caring/internal/payment/service"
	"github.com/piigyy/sharing-is-caring/pkg/database"
	"github.com/piigyy/sharing-is-caring/pkg/logger"
	"github.com/piigyy/sharing-is-caring/pkg/server"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	const caller = "main"
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		done()
		if r := recover(); r != nil {
			log.Fatalf("application panic: %v", r)
		}
	}()
	logger.Info(ctx, caller, "starting payment service")

	var cfg model.Config
	if err := config.ReadConfigFromFile("payment", &cfg); err != nil {
		logger.Fatal(ctx, caller, "config.ReadConfigFromFile return an error: %v", err)
	}

	mongoURI := fmt.Sprintf(cfg.Database.URI, cfg.Database.DB)
	mongoClient, err := database.NewMongoClient(ctx, mongoURI)
	if err != nil {
		logger.Fatal(ctx, caller, "database.NewMongoClient return an error: %v", err)
	}

	mongoDB := mongoClient.Database(cfg.Database.DB)

	logger.Info(ctx, caller, "starting payment service on port %s", cfg.Port)
	err = initService(ctx, cfg, mongoDB)
	done()

	if err != nil {
		logger.Fatal(ctx, caller, "initService return an error: %v", err)
	}

	logger.Info(ctx, caller, "payment server successfully shutdown")
}

func initService(ctx context.Context, cfg model.Config, mongoDB *mongo.Database) error {
	paymentRepository := repository.NewPaymentRepository(mongoDB)
	paymentService := paymentService.NewPayment(
		cfg.Payment.URL,
		cfg.Payment.Key,
		paymentRepository,
	)

	creds, credsErr := credentials.NewServerTLSFromFile(cfg.Certfile, cfg.Keyfile)
	if credsErr != nil {
		return credsErr
	}

	gRPC := grpc.NewServer(grpc.Creds(creds))
	proto.RegisterPaymentServiceServer(gRPC, paymentService)

	srv, err := server.New(cfg.Port)
	if err != nil {
		return fmt.Errorf("failed to create new server handler: %w", err)
	}

	return srv.ServeGRPC(ctx, gRPC)
}
