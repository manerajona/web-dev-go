package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	guuid "github.com/google/uuid"
)

const (
	key        = "01070709010805000600040102090804010507060506080807070708000501080701090607010506030004030307080409030701090900050808030501000506"
	expiration = 10 // seg
	CookieName = "jwt-token"
)

type UserClaims struct {
	jwt.StandardClaims
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("token has expired")
	}

	if u.ID == "" {
		return fmt.Errorf("invalid session Name")
	}

	return nil
}

func createToken(username string) (token string, err error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiration * time.Second).Unix(),
		},
		ID:       guuid.New().String(),
		Username: username,
	}

	token, err = jwt.NewWithClaims(jwt.SigningMethodHS512, &claims).SignedString([]byte(key))
	if err != nil {
		return "", fmt.Errorf("error in createToken when signing token: %w", err)
	}
	return
}

func parseToken(signedToken string) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(signedToken, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("invalid signing algorithm")
		}
		return []byte(key), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error while parsing token: %w", err)
	}

	if !t.Valid {
		return nil, fmt.Errorf("error while parsing token, token is not valid")
	}

	return t.Claims.(*UserClaims), nil
}
