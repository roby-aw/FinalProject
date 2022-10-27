package routes

// func RoutesComment(fiber *fiber.App, conf config.Config) {
// 	db := config.InitDatabase(conf)

// 	repo := repository.NewCommentRepository(db)
// 	svc := service.NewCommentService(repo, conf)
// 	ctrl := controllers.NewCommentController(svc)

// 	app := fiber.Group("/comments")

// 	app.Post("/", middlewares.JwtMiddleware(), ctrl.CreateComment)
// 	app.Get("/", middlewares.JwtMiddleware(), ctrl.GetComment)
// 	app.Put("/:commentId", middlewares.JwtMiddleware(), ctrl.UpdateComment)
// 	app.Delete("/:commentId", middlewares.JwtMiddleware(), ctrl.DeleteComment)
// }
