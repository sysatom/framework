package service

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sysatom/framework/internal/repository"
	"github.com/sysatom/framework/pkg/config"
	"github.com/sysatom/framework/pkg/types"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(_ config.Type, userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Login(ctx context.Context, username, password string) (id uint64, token string, err error) {
	user, err := s.userRepo.GetMerchantAccount(ctx, username)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return
	}

	// gen jwt token
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  types.Uint64ID(user.ID),
		"nbf": time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
	})

	token, err = t.SignedString([]byte(config.App.JWTSecret))
	if err != nil {
		return
	}

	return
}
