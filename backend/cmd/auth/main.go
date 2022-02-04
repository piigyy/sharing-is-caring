package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/piigyy/sharing-is-caring/config"
	"github.com/piigyy/sharing-is-caring/internal/auth/repository"
	authServer "github.com/piigyy/sharing-is-caring/internal/auth/server"
	"github.com/piigyy/sharing-is-caring/internal/auth/service"
	"github.com/piigyy/sharing-is-caring/pkg/database"
	"github.com/piigyy/sharing-is-caring/pkg/server"
	"github.com/piigyy/sharing-is-caring/pkg/token"
)

func main() {
	fmt.Println("auth service")
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		done()
		if r := recover(); r != nil {
			panic(fmt.Sprintf("application panic: %v", r))
		}
	}()

	cfg, cfgErr := config.ReadConfig()
	if cfgErr != nil {
		panic(cfgErr)
	}

	mongoURI := fmt.Sprintf(cfg.Mongo.URI, cfg.Mongo.DB.Auth)
	mongoClient, mongoClientErr := database.NewMongoClient(ctx, mongoURI)
	if mongoClientErr != nil {
		panic(mongoClientErr)
	}

	authDB := mongoClient.Database(cfg.Mongo.DB.Auth)
	userCollection := authDB.Collection("user")

	authRepository := repository.NewUserMongoDB(userCollection)
	tokenCreator := token.NewJWTToken(cfg.JWTSecret)
	authService := service.NewAuthService(authRepository, tokenCreator)
	authServer := authServer.NewHTTPServer(cfg, authService)

	srv, srvErr := server.New(cfg.Port.Auth)
	if srvErr != nil {
		panic(srvErr)
	}

	log.Printf("server is listening on %s\n", cfg.Port.Auth)
	if err := srv.ServeHTTPHandler(ctx, authServer.CORS(authServer.HandlerLogging(authServer.Routes(ctx)))); err != nil {
		panic(err)
	}
}
