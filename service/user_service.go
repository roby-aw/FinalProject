package service

import (
	"final-project/config"
	"final-project/dto"
	"final-project/interfaces"
	"final-project/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Conf config.AppConfig
	repo interfaces.UserRepository
}

func NewUserService(conf config.AppConfig, repo interfaces.UserRepository) interfaces.UserService {
	return &UserService{
		Conf: conf,
		repo: repo,
	}
}

func (s *UserService) Register(user *dto.Register) (dto.ResponseRegister, error) {
	// bcrypt password
	pass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)

	// set password
	user.Password = string(pass)

	res, err := s.repo.Register(user)
	var response dto.ResponseRegister
	if err != nil {
		return response, err
	}
	response.ID = res.ID
	response.Username = res.Username
	response.Email = res.Email
	response.Age = res.Age
	return response, nil
}

func (s *UserService) Login(email, password string) (string, error) {
	login, err := s.repo.Login(email, password)
	if err != nil {
		return "", err
	}
	_ = bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(password))

	token, _ := utils.GenerateAccessToken(login.ID, login.Email)

	return token, nil
}

func (s *UserService) UpdateUser(id uint, user *dto.UpdateUser) (dto.UpdateUser, error) {
	return s.repo.UpdateUser(id, user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
