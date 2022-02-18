package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/piigyy/sharing-is-caring/config"
	"github.com/piigyy/sharing-is-caring/internal/gateway/model"
	"github.com/piigyy/sharing-is-caring/internal/gateway/repository"
	authServer "github.com/piigyy/sharing-is-caring/internal/gateway/server"
	"github.com/piigyy/sharing-is-caring/internal/gateway/service"
	"github.com/piigyy/sharing-is-caring/pkg/database"
	"github.com/piigyy/sharing-is-caring/pkg/logger"
	"github.com/piigyy/sharing-is-caring/pkg/middleware"
	"github.com/piigyy/sharing-is-caring/pkg/server"
	"github.com/piigyy/sharing-is-caring/pkg/token"
)

func main() {
	const caller = "main"
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		done()
		if r := recover(); r != nil {
			logger.Fatal(ctx, caller, "application panic: %v", r)
		}
	}()
	logger.Info(ctx, caller, "starting service gateway")

	var cfg model.Config
	err := config.ReadConfigFromFile("auth", &cfg)
	if err != nil {
		logger.Fatal(ctx, caller, "config.ReadConfigFromFile return an error: %v", err)
	}

	mongoURI := fmt.Sprintf(cfg.Mongo.URI, cfg.Mongo.DB.Auth)
	mongoClient, mongoClientErr := database.NewMongoClient(ctx, mongoURI)
	if mongoClientErr != nil {
		logger.Fatal(ctx, caller, "database.NewMongoClient return an error: %v", mongoClientErr)
	}

	authDB := mongoClient.Database(cfg.Mongo.DB.Auth)
	userCollection := authDB.Collection("user")

	authRepository := repository.NewUserMongoDB(userCollection)
	tokenService := token.NewJWTToken(cfg.JWTSecret)
	authService := service.NewAuthService(authRepository, tokenService)
	middleware := middleware.NewMiddleware(tokenService)
	authServer := authServer.NewHTTPServer(&cfg, authService, middleware)

	srv, srvErr := server.New(cfg.Port)
	if srvErr != nil {
		logger.Fatal(ctx, caller, "server.New return an error: %v", srvErr)
	}

	logger.Info(ctx, caller, "serving gateway serving on port: %s", cfg.Port)
	if err := srv.ServeHTTPHandler(ctx, authServer.CORS(authServer.HandlerLogging(authServer.Routes(ctx)))); err != nil {
		logger.Fatal(ctx, caller, "srv.ServeHTTPHandler return an error: %v", srvErr)
	}
}
