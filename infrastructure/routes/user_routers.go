package routes

// func RoutesUser(fiber *fiber.App, conf config.Config) {
// 	db := config.InitDatabase(conf)

// 	repo := repository.NewUserRepository(db)
// 	srv := service.NewUserService(conf, repo)
// 	ctrl := controllers.NewUserController(srv)

// 	app := fiber.Group("/users")

// 	app.Post("/register", ctrl.Register)
// 	app.Post("/login", ctrl.Login)
// 	app.Put("/:userId", middlewares.JwtMiddleware(), ctrl.UpdateUser)
// 	app.Delete("/:userId", middlewares.JwtMiddleware(), ctrl.DeleteUser)
// }
