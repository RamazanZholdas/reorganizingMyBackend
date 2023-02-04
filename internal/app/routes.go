package app

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}

/*
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Get("/user", controllers.User)
	app.Post("/logout", controllers.Logout)
	app.Get("/getAllItems", controllers.GetAllItems)
	app.Post("/insertProduct", controllers.InsertProduct)
	app.Post("/addToCart/:order", controllers.AddToCart)
	app.Get("/getCart", controllers.GetAllFromCart)
	app.Get("/getKeyboard/:order", controllers.GetKeyboard)
	app.Get("/getNews", controllers.GetNews)
	app.Get("/getNews/:order", controllers.GetNewsById)
	app.Post("/addNews", controllers.AddNews)
*/
