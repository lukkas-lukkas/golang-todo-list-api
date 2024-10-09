package register_user

import (
	"github.com/lukkas-lukkas/golang-todo-list-api/app/domain/helpers"
	"github.com/lukkas-lukkas/golang-todo-list-api/app/domain/user"
)

type RegisterUserService struct {
	repository user.Repository
	encryptor  helpers.Encryptor
}

func NewRegisterUserService(repo user.Repository, encryptor helpers.Encryptor) *RegisterUserService {
	return &RegisterUserService{repo, encryptor}
}

func (s RegisterUserService) CreateUser(dto UserDto) (user.User, error) {
	pass, err := s.encryptor.Encrypt(dto.Password)
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
