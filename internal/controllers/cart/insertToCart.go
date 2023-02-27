package cart

import (
	"strconv"

	"github.com/RamazanZholdas/KeyboardistSV2/internal/app"
	"github.com/RamazanZholdas/KeyboardistSV2/internal/jwt"
	"github.com/RamazanZholdas/KeyboardistSV2/internal/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)
/*
	TODO:
		- add option id
*/
func InsertToCart(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	claims, err := jwt.ExtractTokenClaimsFromCookie(cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	var user models.User
	err = app.GetMongoInstance().FindOne("users", bson.M{"email": claims.Issuer}, &user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	order := c.Params("order")
	number, err := strconv.Atoi(order)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "The order must be a valid number.",
		})
	}

	var product models.Product
	err = app.GetMongoInstance().FindOne("products", bson.M{"order": number}, &product)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Product not found"})
	}

	user.Cart = append(user.Cart, product.ID)

	err = app.GetMongoInstance().UpdateOne("users", bson.M{"email": claims.Issuer}, bson.M{"$set": bson.M{"cart": user.Cart}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update user"})
	}

	return c.JSON(fiber.Map{"message": "Product added to cart"})
}
