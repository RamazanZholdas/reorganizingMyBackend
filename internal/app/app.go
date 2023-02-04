package app

import (
	"log"
	"os"

	"github.com/RamazanZholdas/KeyboardistSV2/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type App struct {
	Fiber *fiber.App
}

var (
	allCollections = []string{
		os.Getenv("COLLECTION_USERS"),
		os.Getenv("COLLECTION_PRODUCTS"),
		os.Getenv("COLLECTION_NEWS"),
		os.Getenv("COLLECTION_WIKI"),
		os.Getenv("COLLECTION_SERVICE_MASTERS"),
		os.Getenv("COLLECTION_PURCHASE_HISTORY"),
	}
	MongoInstance = &database.MongoDB{}
)

func New(mongoURI, dbName string) *App {
	err := MongoInstance.Connect(mongoURI, dbName)
	if err != nil {
		log.Fatalln("Error connecting to MongoDB:", err)
	} else if err.Error() != "database already exists" {
		log.Println("Database aleady exists")
	}

	errors := MongoInstance.CreateCollections(allCollections)
	if len(errors) > 0 {
		for _, err := range errors {
			if err.Error() != "collection already exists" {
				log.Fatalln("Error creating collection:", err)
			}
		}
		log.Println("We have some collections that already exists: ", errors)
	}

	fiber := fiber.New()

	fiber.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	fiber.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	Setup(fiber)

	return &App{
		Fiber: fiber,
	}
}

func (a *App) Close() {
	MongoInstance.Disconnect()
	a.Fiber.Shutdown()
}

func GetMongoInstance() *database.MongoDB {
	return MongoInstance
}
