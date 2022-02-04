package middleware

import (
	"context"
	"fmt"
	"net/http"
	"regexp"

	"github.com/golang-jwt/jwt/v4"
	"github.com/piigyy/sharing-is-caring/pkg/presenter"
	"github.com/piigyy/sharing-is-caring/pkg/server"
	"github.com/piigyy/sharing-is-caring/pkg/token"
)

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(`(^bearer\s|^Bearer\s)`)
}

type (
	Auth interface {
		Authotization() server.Adapter
	}
)

type mdlw struct {
	tokenVerificator token.TokenVerificator
}

func NewMiddleware(tokenVeridicator token.TokenVerificator) *mdlw {
	return &mdlw{
		tokenVerificator: tokenVeridicator,
	}
}

func (m *mdlw) Authotization() server.Adapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			authorizationHeader := r.Header.Get("Authorization")
			tokenRaw := re.ReplaceAllString(authorizationHeader, "")

			signedToken, err := m.tokenVerificator.VerifyToken(tokenRaw)
			if err != nil {
				presenter.ErrResponse(rw, http.StatusUnauthorized, err)
			}

			if err := m.tokenVerificator.ValidateToken(signedToken); err != nil {
				presenter.ErrResponse(rw, http.StatusUnauthorized, err)
			}

			claims := signedToken.Claims.(jwt.MapClaims)

			credentials := map[string]string{
				"id":    claims["ID"].(string),
				"email": claims["Email"].(string),
				"name":  claims["Name"].(string),
			}

			fmt.Printf("credentials: %+v\n", credentials)
			ctx := context.WithValue(r.Context(), "credentials", credentials)
			next.ServeHTTP(rw, r.WithContext(ctx))
		})
	}
}
