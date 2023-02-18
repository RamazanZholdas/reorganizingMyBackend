package cart

import (
	"github.com/RamazanZholdas/KeyboardistSV2/internal/app"
	"github.com/RamazanZholdas/KeyboardistSV2/internal/jwt"
	"github.com/RamazanZholdas/KeyboardistSV2/internal/models"
	"github.com/RamazanZholdas/KeyboardistSV2/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllFromUsersCart(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	claims, err := jwt.ExtractTokenClaimsFromCookie(cookie)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid token"})
	}

	var user models.User
	err = app.GetMongoInstance().FindOne("users", bson.M{"email": claims.Issuer}, &user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User not found"})
	}

	if len(user.Cart) == 0 {
		return c.JSON(fiber.Map{"message": "Cart is empty"})
	}

	var products []models.Product
	for _, id := range user.Cart {
		var product models.Product
		err = app.GetMongoInstance().FindOne("products", bson.M{"_id": id}, &product)
		if err != nil {
			utils.LogWarning("This product not found:"+product.Name, err)
			continue // Skip if product not found
		}
		products = append(products, product)
	}

	return c.JSON(products)
}
