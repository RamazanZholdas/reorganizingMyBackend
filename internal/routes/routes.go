package routes

import (
	"github.com/RamazanZholdas/KeyboardistSV2/internal/controllers/authorization"
	"github.com/RamazanZholdas/KeyboardistSV2/internal/controllers/product"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	//authorization
	app.Post("/register", authorization.Register)
	app.Post("/login", authorization.Login)
	app.Post("/logout", authorization.Logout)

	//products
	app.Get("/getAllProducts", product.GetAllProducts)
	app.Get("/getProduct/:order", product.GetProduct)
	app.Post("/insertProduct", product.InsertProduct)

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
