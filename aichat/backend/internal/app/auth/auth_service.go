package auth

import (
	"aichat/internal/app/dto"
	"aichat/internal/app/models"
	"crypto/sha256"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type AuthService interface {
	Login(request *dto.LoginRequest) (*dto.LoginResponse, error)
	GenerateToken() (string, error)
}

type authService struct {
	db Repository
}

func NewService(authRepo Repository) AuthService {
	return &authService{
		db: authRepo,
	}
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	// Encode password to sha256
	password := encodeSha256(req.Password)
	req.Password = password

	// Check username and password in db
	user, err := s.db.GetUser(req)
	if err != nil {
		return &dto.LoginResponse{}, err
	}
	if user == (models.User{}) {
		return &dto.LoginResponse{}, errors.New("Access Denied")
	}

	// If all good, generate the token
	token, err := s.GenerateToken()
	if err != nil {
		return nil, err
	}
	return &dto.LoginResponse{
		Token: token,
	}, nil
}

func (s *authService) GenerateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "user123"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Set your secret key
	secretKey := []byte("your_secret_key")

	// Sign the token with your secret key
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func encodeSha256(val string) string {
	s := val
	h := sha256.New()
	h.Write([]byte(s))

	hashed := h.Sum(nil)

	return fmt.Sprintf("%x", hashed)
}
