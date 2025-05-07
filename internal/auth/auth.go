package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"
	modals "what/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

type UserTokenClaims struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Type int    `json:"type"`
	jwt.RegisteredClaims
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func createToken(user *modals.User) (string, error) {

	claims := UserTokenClaims{
		UUID: user.UUID.String(),
		Name: user.Name,
		Type: user.Type,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))

	if err != nil {
		fmt.Print(err)
		return "", err
	}

	return signedToken, nil
}

func auth(r *http.Request) error {
	cookie, err := r.Cookie("token")
	if err != nil {
		return err
	}

	tokenString := cookie.Value
	_, err = verifyToken(tokenString)
	if err != nil {
		return err
	}

	return nil
}

func SetToken(w http.ResponseWriter, user *modals.User) error {

	token, err := createToken(user)

	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Domain:   "",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return err
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth(r); err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
