package routes

// func RoutesPhoto(fiber *fiber.App, conf config.Config) {
// 	db := config.InitDatabase(conf)

// 	repo := repository.NewPhotoRepository(db)
// 	svc := service.NewPhotoService(repo, conf)
// 	ctrl := controllers.NewPhotoController(svc)

// 	app := fiber.Group("/photos")

// 	app.Post("/", middlewares.JwtMiddleware(), ctrl.CreatePhoto)
// 	app.Get("/", middlewares.JwtMiddleware(), ctrl.GetAllPhoto)
// 	app.Put("/:photoId", middlewares.JwtMiddleware(), ctrl.UpdatePhoto)
// 	app.Delete("/:photoId", middlewares.JwtMiddleware(), ctrl.DeletePhoto)
// }
