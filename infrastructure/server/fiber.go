package server

import (
	"final-project/api"
	"final-project/app/modules"
	"final-project/config"
	"final-project/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitServer() *fiber.App {
	app := fiber.New()
	config := config.GetConfig()
	dbCon := utils.NewConnectionDatabase(config)

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	controller := modules.RegistrationModules(dbCon, config)
	api.RegistrationPath(app, controller)

	return app
}
