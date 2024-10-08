package register_user

import (
	"github.com/lukkas-lukkas/golang-todo-list-api/internal/domain/helpers"
	"github.com/lukkas-lukkas/golang-todo-list-api/internal/domain/user"
)

type RegisterUserService struct {
	repository     user.Repository
	encryptService helpers.EncryptService
}

func NewRegisterUserService(repo user.Repository, encrypt helpers.EncryptService) *RegisterUserService {
	return &RegisterUserService{repo, encrypt}
}

func (s RegisterUserService) CreateUser(dto UserDto) (user.User, error) {
	pass, err := s.encryptService.Encrypt(dto.Password)
	if err != nil {
		return user.User{}, err
	}

	newUser := user.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: pass,
	}

	created, err := s.repository.Create(newUser)
	if err != nil {
		return user.User{}, err
	}

	return created, nil
}
