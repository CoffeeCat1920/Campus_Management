package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createAdminToken() (string, error) {
	claims := AdminTokenClaims{
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

func verifyAdminToken(tokenString string) (*AdminTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AdminTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AdminTokenClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func authAdmin(r *http.Request) error {
	cookie, err := r.Cookie("admin_token")
	if err != nil {
		return err
	}

	tokenString := cookie.Value
	_, err = verifyAdminToken(tokenString)
	if err != nil {
		return err
	}

	return nil
}

func (a *authService) AdminAuthData(r *http.Request) error {
	return authAdmin(r)
}

func (a *authService) AuthAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authAdmin(r); err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func (a *authService) SetAdminToken(w http.ResponseWriter) error {
	token, err := createAdminToken()
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "admin_token",
		Value:    token,
		Path:     "/",
		Domain:   "",
		HttpOnly: true,

		SameSite: http.SameSiteLaxMode,
	})

	return nil
}

func (a *authService) ClearAdminToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "admin_token",
		Value:    "",
		Path:     "/",
		Domain:   "",
		HttpOnly: true,
		Expires:  time.Unix(0, 0), // Set to a time in the past
		MaxAge:   -1,
		SameSite: http.SameSiteLaxMode,
	})
}
