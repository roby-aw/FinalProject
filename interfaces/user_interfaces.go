package interfaces

import (
	"final-project/dto"
	"final-project/entity"

	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	Register(user *dto.Register) (entity.User, error)
	Login(email, password string) (dto.Login, error)
	UpdateUser(id uint, user *dto.UpdateUser) (dto.UpdateUser, error)
	DeleteUser(id uint) error
}

type UserService interface {
	Register(user *dto.Register) (dto.ResponseRegister, error)
	Login(email, password string) (string, error)
	UpdateUser(id uint, user *dto.UpdateUser) (dto.UpdateUser, error)
	DeleteUser(id uint) error
}

type UserController interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}
