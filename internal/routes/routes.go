package routes

import (
	"github.com/RamazanZholdas/KeyboardistSV2/internal/controllers/authorization"
	"github.com/RamazanZholdas/KeyboardistSV2/internal/controllers/cart"
	"github.com/RamazanZholdas/KeyboardistSV2/internal/controllers/news"
	"github.com/RamazanZholdas/KeyboardistSV2/internal/controllers/product"
	servicemaster "github.com/RamazanZholdas/KeyboardistSV2/internal/controllers/service_master"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	//authorization
	app.Post("/register", authorization.Register)
	app.Post("/login", authorization.Login)
	app.Post("/logout", authorization.Logout)

	//products
	app.Get("/getAllItems", product.GetAllProducts)
	app.Get("/getProduct/:order", product.GetProduct)
	app.Post("/insertProduct", product.InsertProduct)

	//cart
	app.Get("/getCart", cart.GetAllFromUsersCart)
	app.Post("/addToCart/:order", cart.InsertToCart)

	//news
	app.Get("/getNews", news.GetAllNews)
	app.Get("/getNews/:order", news.GetNews)
	app.Post("/addNews", news.InsertNews)

	//service_master
	app.Get("/getAllServiceMasters", servicemaster.GetAllServiceMasters)
	app.Get("/getServiceMaster/:order", servicemaster.GetServiceMaster)
	app.Post("/insertServiceMaster", servicemaster.InsertServiceMaster)
}
