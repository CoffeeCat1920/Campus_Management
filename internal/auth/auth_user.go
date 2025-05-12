package auth

import (
	"fmt"
	"net/http"
	"time"
	modals "what/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

func createUserToken(user *modals.User) (string, error) {

	fmt.Printf("\nUser to have it's token set : %v\n", user)

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

func (a *authService) SetUserToken(w http.ResponseWriter, user *modals.User) error {
	token, err := createUserToken(user)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "user_token",
		Value:    token,
		Path:     "/",
		Domain:   "",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return err
}

func verifyUserToken(tokenString string) (*UserTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserTokenClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func authStudent(r *http.Request) error {
	cookie, err := r.Cookie("user_token")
	if err != nil {
		return err
	}

	tokenString := cookie.Value
	claims, err := verifyUserToken(tokenString)
	if err != nil {
		return err
	}

	if claims.Type != 0 {
		return fmt.Errorf("unauthorized: not a student")
	}

	return nil
}

func (a *authService) StudentAuthData(r *http.Request) (*UserTokenClaims, error) {
	cookie, err := r.Cookie("user_token")
	if err != nil {
		return nil, err
	}

	tokenString := cookie.Value
	claims, err := verifyUserToken(tokenString)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\n%v", claims.Type)

	if claims.Type != 0 {
		return nil, fmt.Errorf("unauthorized: not a student \n")
	}

	return claims, nil
}

func (a *authService) LibrarianAuthData(r *http.Request) (*UserTokenClaims, error) {
	cookie, err := r.Cookie("user_token")
	if err != nil {
		return nil, err
	}

	tokenString := cookie.Value
	claims, err := verifyUserToken(tokenString)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\n%v", claims.Type)

	if claims.Type != 1 {
		return nil, fmt.Errorf("unauthorized: not a librarian \n")
	}

	return claims, nil
}

func (a *authService) AuthStudent(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authStudent(r); err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func authLibrarian(r *http.Request) error {
	cookie, err := r.Cookie("user_token")
	if err != nil {
		return err
	}

	tokenString := cookie.Value
	claims, err := verifyUserToken(tokenString)
	if err != nil {
		return err
	}

	if claims.Type != 1 {
		return fmt.Errorf("unauthorized: not an admin")
	}

	return nil
}

func (s *authService) AuthLibrarian(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authLibrarian(r); err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func (a *authService) ClearUserToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Domain:   "",
		HttpOnly: true,
		Expires:  time.Unix(0, 0), // Set to a time in the past
		MaxAge:   -1,
		SameSite: http.SameSiteLaxMode,
	})
}
