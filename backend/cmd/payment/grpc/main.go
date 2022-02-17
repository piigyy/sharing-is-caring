package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/piigyy/sharing-is-caring/config"
	"github.com/piigyy/sharing-is-caring/internal/payment/model"
	"github.com/piigyy/sharing-is-caring/internal/payment/proto"
	paymentService "github.com/piigyy/sharing-is-caring/internal/payment/service"
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

	paymentService := paymentService.NewPayment(cfg.Payment.URL, cfg.Payment.Key)
	GRPCSrv := grpc.NewServer()
	proto.RegisterPaymentServiceServer(GRPCSrv, paymentService)

	log.Printf("config: %+v\n", cfg)
	srv, err := server.New(cfg.Port)
	if err != nil {
		log.Panicf("err server.New: %v\n", err)
	}

	log.Println("starting payment service")
	srv.ServeGRPC(ctx, GRPCSrv)
}
