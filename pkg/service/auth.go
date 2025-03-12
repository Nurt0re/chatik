package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/Nurt0re/chatik"
	"github.com/Nurt0re/chatik/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "ksaiujhdf49857"
	tokenTTL   = 12 * time.Hour
	signingKey = "8937w4bfncy5wyg"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}


func (s *AuthService) CreateUser(user chatik.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err:= jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token)(interface{}, error){
		if _, ok:= token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err!=nil{
		return 0 ,err
	}

	claims, ok:= token.Claims.(*tokenClaims); 
	if !ok{
		return 0 , errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil




}

func (s *AuthService) GenerateToken(username, password string) (string, error) {

	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		int(user.ID),
	})

	return token.SignedString([]byte(signingKey))

}
