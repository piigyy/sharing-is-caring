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
	"github.com/piigyy/sharing-is-caring/pkg/server"
	"google.golang.org/grpc"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		done()
		if r := recover(); r != nil {
			log.Fatalf("application panic: %v", r)
		}
	}()
	log.Println("statring payment service")

	var cfg model.Config
	if err := config.ReadConfigFromFile("payment", &cfg); err != nil {
		log.Panicf("err config.ReadConfigFromFile: %v\n", err)
	}

	mongoURI := fmt.Sprintf(cfg.Database.URI, cfg.Database.DB)
	mongoClient, err := database.NewMongoClient(ctx, mongoURI)
	if err != nil {
		panic(err)
	}

	mongoDB := mongoClient.Database(cfg.Database.DB)
	paymentRepository := repository.NewPaymentRepository(mongoDB)

	paymentService := paymentService.NewPayment(
		cfg.Payment.URL,
		cfg.Payment.Key,
		paymentRepository,
	)
	GRPCSrv := grpc.NewServer()
	proto.RegisterPaymentServiceServer(GRPCSrv, paymentService)

	srv, err := server.New(cfg.Port)
	if err != nil {
		log.Panicf("err server.New: %v\n", err)
	}

	log.Printf("starting payment service on port %s", cfg.Port)
	srv.ServeGRPC(ctx, GRPCSrv)
}
