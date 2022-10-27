package controllers

import (
	"final-project/dto"
	"final-project/interfaces"
	"final-project/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Usercontrollers struct {
	UserService interfaces.UserService
}

func NewUserController(service interfaces.UserService) *Usercontrollers {
	return &Usercontrollers{
		UserService: service,
	}
}

func (h *Usercontrollers) Register(c *fiber.Ctx) error {
	user := dto.Register{}

	_ = c.BodyParser(&user)

	err := validator.New().Struct(user)
	if err != nil {
		return utils.HandleErrorValidator(err, c)
	}

	result, err := h.UserService.Register(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    400,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"code":    201,
		"data":    result,
	})
}

func (h *Usercontrollers) Login(c *fiber.Ctx) error {
	user := dto.Login{}

	_ = c.BodyParser(&user)

	err := validator.New().Struct(user)
	if err != nil {
		return utils.HandleErrorValidator(err, c)
	}

	result, err := h.UserService.Login(user.Email, user.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"code":    400,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successfully",
		"code":    200,
		"data":    result,
	})
}

func (h *Usercontrollers) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("userId")
	uid, _ := strconv.Atoi(id)
	user := dto.UpdateUser{}

	_ = c.BodyParser(&user)

	err := validator.New().Struct(user)
	if err != nil {
		return utils.HandleErrorValidator(err, c)
	}

	result, err := h.UserService.UpdateUser(uint(uid), &user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
			"code":    404,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
		"code":    200,
		"data":    result,
	})
}

func (h *Usercontrollers) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("userId")
	uid, _ := strconv.Atoi(id)

	err := h.UserService.DeleteUser(uint(uid))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
			"code":    404,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Your account has been successfully deleted",
		"code":    200,
	})
}
