package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/piigyy/sharing-is-caring/config"
	"github.com/piigyy/sharing-is-caring/internal/auth/model"
	"github.com/rs/cors"
)

type httpServer struct {
	cfg         *config.Config
	authService model.ServiceReaderWriter
}

func NewHTTPServer(cfg *config.Config, authService model.ServiceReaderWriter) *httpServer {
	log.Println("creating new http server")
	return &httpServer{
		cfg:         cfg,
		authService: authService,
	}
}

func (s *httpServer) Routes(ctx context.Context) http.Handler {
	log.Println("initializing auth server routes")

	mux := mux.NewRouter()
	mux.Use(s.APIMiddleware())

	mux.HandleFunc("/api/v1/login", s.Login).Methods(http.MethodPost)
	mux.HandleFunc("/api/v1/registrations", s.RegisterUser).Methods(http.MethodPost)

	return mux
}

func (s *httpServer) CORS(mux http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
	})
	return c.Handler(mux)
}

func (s *httpServer) HandlerLogging(mux http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, mux)
}

func (s *httpServer) APIMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Date", time.Now().Format(time.RFC1123))
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Header().Set("Connection", "keep-alive")

			next.ServeHTTP(w, r)
		})
	}
}
