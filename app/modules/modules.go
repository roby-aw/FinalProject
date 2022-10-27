package modules

import (
	"final-project/api"
	"final-project/api/controllers"
	"final-project/config"
	"final-project/repository"
	"final-project/service"
	"final-project/utils"
)

func RegistrationModules(dbCon *utils.DatabaseConnection, conf *config.AppConfig) api.Controller {
	commentPermitRepository := repository.NewCommentRepository(dbCon.Postgres)
	commentPermitService := service.NewCommentService(commentPermitRepository, *conf)
	commentPermitController := controllers.NewCommentController(commentPermitService)

	photosPermitRepository := repository.NewPhotoRepository(dbCon.Postgres)
	photosPermitService := service.NewPhotoService(photosPermitRepository, *conf)
	photosPermitController := controllers.NewPhotoController(photosPermitService)

	sosmedPermitRepository := repository.NewSosmedRepository(dbCon.Postgres)
	sosmedPermitService := service.NewSosmedService(sosmedPermitRepository, *conf)
	sosmedPermitController := controllers.NewSosmedController(sosmedPermitService)

	userPermitRepository := repository.NewUserRepository(dbCon.Postgres)
	userPermitService := service.NewUserService(*conf, userPermitRepository)
	userPermitController := controllers.NewUserController(userPermitService)

	controller := api.Controller{
		CommentController: commentPermitController,
		PhotosController:  photosPermitController,
		SocmedController:  sosmedPermitController,
		UserController:    userPermitController,
	}

	return controller
}
