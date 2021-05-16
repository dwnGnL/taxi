package jwtImpl

import (
	"errors"
	"fmt"
	"taxi/pkg/auth"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// AuthClaims ...
type AuthClaims struct {
	jwt.StandardClaims
	ID int
}

// JWT ...
type JWT struct {
	signingKey     []byte
	expireDuration int64
}

// NewJWT ...
func NewJWT(signingKey string, expireDuration int64) (*JWT, error) {
	if signingKey == "" {
		return nil, errors.New("empty signingKey")
	}

	return &JWT{
		signingKey:     []byte(signingKey),
		expireDuration: expireDuration,
	}, nil
}

// NewJWT ...
func (t *JWT) NewJWT(Id int) (*string, error) {
	fmt.Println("(t *JWT) NewJWT(Id int) start ")
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, AuthClaims{
		ID: Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(time.Minute * 5)).Unix(),
		},
	}).SignedString(t.signingKey)
	fmt.Println("err new jwt token =>", err)
	fmt.Println("(t *JWT) NewJWT(Id int) end token =>", token)
	return &token, err

}

// ParseToken ...
func (a *JWT) ParseToken(accessToken *string) (int, error) {
	token, err := jwt.ParseWithClaims(*accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return a.signingKey, nil
	})
	fmt.Println("err parsing token =>", err)
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.ID, nil
	}
	fmt.Println("error parsing ErrInvalidAccessToken")
	return 0, auth.ErrInvalidAccessToken
}

// func (m *JWT) NewRefreshToken() (string, error) {
// 	b := make([]byte, 32)

// 	s := rand.NewSource(time.Now().Unix())
// 	r := rand.New(s)

// 	_, err := r.Read(b)
// 	if err != nil {
// 		return "", err
// 	}

// 	return fmt.Sprintf("%x", b), nil
// }
