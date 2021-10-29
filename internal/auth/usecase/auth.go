package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	exam "example.com/internal"
	"example.com/models"
	"github.com/dgrijalva/jwt-go/v4"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *models.User `json:"user"`
}

type UserUseCase struct {
	userRepo       exam.UserRepository
	hashSalt       string
	signInKey      []byte
	expireDuration time.Duration
}

func NewUserUseCase(
	userRepo exam.UserRepository,
	hash string,
	signInKey []byte,
	tokenTtl time.Duration) *UserUseCase {
	return &UserUseCase{
		userRepo:       userRepo,
		hashSalt:       hash,
		signInKey:      signInKey,
		expireDuration: time.Second * tokenTtl,
	}
}

func (useCase *UserUseCase) Login(ctx context.Context, authParams models.LoginParam) (string, error) {
	authParams.Password = getEncodedPassword(authParams.Password, useCase.hashSalt)
	user, err := useCase.userRepo.GetUser(ctx, authParams)
	if err != nil {
		return "", err
	}
	claims := AuthClaims{
		User: &user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(useCase.expireDuration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(useCase.signInKey)
}

func (useCase *UserUseCase) Register(ctx context.Context, authParams models.LoginParam) error {
	pwd := getEncodedPassword(authParams.Password, useCase.hashSalt)
	user := models.User{
		Username: authParams.Username,
		Password: pwd,
	}
	return useCase.userRepo.CreateUser(ctx, user)
}

func (useCase *UserUseCase) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return useCase.signInKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, err
}

func getEncodedPassword(pwd string, hash string) string {
	out := sha1.New()
	out.Write([]byte(pwd))
	out.Write([]byte(hash))
	return fmt.Sprintf("%x", out.Sum(nil))
}
