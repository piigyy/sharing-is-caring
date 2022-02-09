package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/piigyy/sharing-is-caring/internal/auth/model"
	"github.com/piigyy/sharing-is-caring/pkg/middleware"
	"github.com/piigyy/sharing-is-caring/pkg/server"
	"github.com/rs/cors"
)

type httpServer struct {
	cfg         *model.Config
	authService model.ServiceReaderWriter
	middleware  middleware.Auth
}

func NewHTTPServer(
	cfg *model.Config,
	authService model.ServiceReaderWriter,
	middleware middleware.Auth,
) *httpServer {
	log.Println("creating new http server")
	return &httpServer{
		cfg:         cfg,
		authService: authService,
		middleware:  middleware,
	}
}

func (s *httpServer) Routes(ctx context.Context) http.Handler {
	log.Println("initializing auth server routes")

	mux := mux.NewRouter()
	mux.Use(s.APIMiddleware())

	mux.HandleFunc("/api/v1/login", s.Login).Methods(http.MethodPost)
	mux.HandleFunc("/api/v1/registrations", s.RegisterUser).Methods(http.MethodPost)

	mux.HandleFunc("/api/v1/users", server.Adapt(
		http.HandlerFunc(s.GetUserDetail),
		s.middleware.Authotization(),
	).ServeHTTP).Methods(http.MethodGet)
	mux.HandleFunc("/api/v1/users", server.Adapt(
		http.HandlerFunc(s.UpdatePasword),
		s.middleware.Authotization(),
	).ServeHTTP).Methods(http.MethodPatch)

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
