package service

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"userBar/back/models"
	"userBar/back/pkg/repository"
)

const (
	salt = "drinkingeveryday1mdrinking"
	signKey = "0maevaMOuSHinde1rU"
)

type AuthService struct {
	rep repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(rep repository.Authorization) *AuthService  {
	return &AuthService{rep: rep}
}

func (a *AuthService) CreateUser(user models.User) (int, error)  {
	user.Password = generatePasswordHash(user.Password)
	return a.rep.CreateUser(user)
}

func generatePasswordHash(password string) string  {
	hash := sha1.New()
	hash.Write([]byte(password))
	sha := base64.URLEncoding.EncodeToString(hash.Sum([]byte(salt)))
	return fmt.Sprintf("%s", sha)
}

func (a *AuthService) GenerateToken(name, password string) (string, error)  {
	user, err := a.rep.GetUser(name, generatePasswordHash(password))
	if err != nil{
		fmt.Println(1)
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
		ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
		IssuedAt: time.Now().Unix(),
	},
	user.Id,
	})
	return token.SignedString([]byte(signKey))
}

func (a *AuthService) ParseToken(accessToken string) (int, error){
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, errors.New("invalid signing method")
		}
		return []byte(signKey), nil
	})
	if err != nil{
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("error in token claims - its not of type tokenClaims")
	}
	return claims.UserId, nil
}