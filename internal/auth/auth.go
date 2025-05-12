package auth

import (
	"net/http"
	"os"
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

type AdminTokenClaims struct {
	jwt.RegisteredClaims
}

type AuthService interface {
	SetUserToken(w http.ResponseWriter, user *modals.User) error
	AuthStudent(next http.HandlerFunc) http.HandlerFunc
	AuthLibrarian(next http.HandlerFunc) http.HandlerFunc
	StudentAuthData(r *http.Request) (*UserTokenClaims, error)
	LibrarianAuthData(r *http.Request) (*UserTokenClaims, error)
	ClearUserToken(w http.ResponseWriter)

	SetAdminToken(w http.ResponseWriter) error
	AuthAdmin(next http.HandlerFunc) http.HandlerFunc
	AdminAuthData(r *http.Request) error
	ClearAdminToken(w http.ResponseWriter)
}

type authService struct{}

func NewAuthService() AuthService {
	return &authService{}
}
