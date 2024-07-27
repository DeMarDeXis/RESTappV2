package service

import (
	"crypto/sha1"
	"fmt"

	gorestapiv2 "github.com/DeMarDeXis/RESTV1"
	"github.com/DeMarDeXis/RESTV1/pkg/storage"
)

const salt = "hjtrfgj32325asfsdfg"

type AuthService struct {
	storage storage.Authorization
}

func NewAuthService(storage storage.Authorization) *AuthService {
	return &AuthService{
		storage: storage,
	}
}

func (s *AuthService) CreateUser(user gorestapiv2.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.storage.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
