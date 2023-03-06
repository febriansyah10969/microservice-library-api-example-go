package middleware

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	ValidateOnlineToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	KeyOnline []byte
}

func NewJWTService(keyOnline []byte) *jwtService {
	return &jwtService{keyOnline}
}

func (js *jwtService) ValidateOnlineToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(js.KeyOnline), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return token, nil
}
