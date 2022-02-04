package token

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrInvalidSigningMethod error = errors.New("invalid signing method")
	ErrInvalidToken         error = errors.New("invalid token")
)

type Claims struct {
	jwt.RegisteredClaims
	ID    string
	Email string
	Name  string
}

type JWTToken struct {
	key string
}

func NewJWTToken(key string) *JWTToken {
	return &JWTToken{key: key}
}

func (s *JWTToken) GenerateAccessToken(ctx context.Context, ID, email, name string) (accessToken string, err error) {
	expiresAt := jwt.NewNumericDate(time.Now().Add(12 * time.Hour))
	jwtTimeNow := jwt.NewNumericDate(time.Now())

	log.Println("creating token claims")
	jwtClaims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "sharing-is-caring-auth",
			ExpiresAt: expiresAt,
			NotBefore: jwtTimeNow,
			IssuedAt:  jwtTimeNow,
		},
		ID:    ID,
		Email: email,
		Name:  name,
	}

	log.Println("signing token")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	signerToken, signedTokenErr := token.SignedString([]byte(s.key))
	if signedTokenErr != nil {
		log.Printf("signing token error: %v\n", signedTokenErr)
		err = signedTokenErr
		return
	}

	return signerToken, nil
}

func (s *JWTToken) VerifyToken(accessToken string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		}
		return []byte(s.key), nil
	})

	if err != nil {
		return nil, err
	}

	return jwtToken, nil
}

func (s *JWTToken) ValidateToken(signedToken *jwt.Token) error {
	if _, ok := signedToken.Claims.(Claims); !ok && !signedToken.Valid {
		return ErrInvalidToken
	}

	return nil
}
