package routes

// func RoutesSosmed(fiber *fiber.App, conf config.Config) {
// 	db := config.InitDatabase(conf)

// 	repo := repository.NewSosmedRepository(db)
// 	svc := service.NewSosmedService(repo, conf)
// 	ctrl := controllers.NewSosmedController(svc)

// 	app := fiber.Group("/socialmedias")

// 	app.Post("/", middlewares.JwtMiddleware(), ctrl.CreateSosmed)
// 	app.Get("/", middlewares.JwtMiddleware(), ctrl.GetSosmed)
// 	app.Put("/:socialMediaId", middlewares.JwtMiddleware(), ctrl.UpdateSosmed)
// 	app.Delete("/:socialMediaId", middlewares.JwtMiddleware(), ctrl.DeleteSosmed)
// }
