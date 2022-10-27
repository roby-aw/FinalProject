package api

import (
	"final-project/api/controllers"
	"final-project/infrastructure/middlewares"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	CommentController *controllers.CommentController
	PhotosController  *controllers.PhotoController
	SocmedController  *controllers.SosmedController
	UserController    *controllers.Usercontrollers
}

func RegistrationPath(e *fiber.App, Controller Controller) {
	commentRouter := e.Group("/comments")

	commentRouter.Post("/", middlewares.JwtMiddleware(), Controller.CommentController.CreateComment)
	commentRouter.Get("/", middlewares.JwtMiddleware(), Controller.CommentController.GetComment)
	commentRouter.Put("/:commentId", middlewares.JwtMiddleware(), Controller.CommentController.UpdateComment)
	commentRouter.Delete("/:commentId", middlewares.JwtMiddleware(), Controller.CommentController.DeleteComment)

	RoutesPhotos := e.Group("/photos")

	RoutesPhotos.Post("/", middlewares.JwtMiddleware(), Controller.PhotosController.CreatePhoto)
	RoutesPhotos.Get("/", middlewares.JwtMiddleware(), Controller.PhotosController.GetAllPhoto)
	RoutesPhotos.Put("/:photoId", middlewares.JwtMiddleware(), Controller.PhotosController.UpdatePhoto)
	RoutesPhotos.Delete("/:photoId", middlewares.JwtMiddleware(), Controller.PhotosController.DeletePhoto)

	routesSocmed := e.Group("/socialmedias")

	routesSocmed.Post("/", middlewares.JwtMiddleware(), Controller.SocmedController.CreateSosmed)
	routesSocmed.Get("/", middlewares.JwtMiddleware(), Controller.SocmedController.GetSosmed)
	routesSocmed.Put("/:socialMediaId", middlewares.JwtMiddleware(), Controller.SocmedController.UpdateSosmed)
	routesSocmed.Delete("/:socialMediaId", middlewares.JwtMiddleware(), Controller.SocmedController.DeleteSosmed)

	routesUser := e.Group("/users")

	routesUser.Post("/register", Controller.UserController.Register)
	routesUser.Post("/login", Controller.UserController.Login)
	routesUser.Put("/:userId", middlewares.JwtMiddleware(), Controller.UserController.UpdateUser)
	routesUser.Delete("/:userId", middlewares.JwtMiddleware(), Controller.UserController.DeleteUser)
}
