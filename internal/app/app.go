package app

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/RamazanZholdas/KeyboardistSV2/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type App struct {
	Fiber *fiber.App
}

var (
	MongoInstance = &database.MongoDB{}
)

func Intitialize(mongoURI, dbName string) (*App, error) {
	err := MongoInstance.Connect(mongoURI, dbName)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	var allCollections = []string{
		os.Getenv("COLLECTION_USERS"),
		os.Getenv("COLLECTION_PRODUCTS"),
		os.Getenv("COLLECTION_NEWS"),
		os.Getenv("COLLECTION_WIKI"),
		os.Getenv("COLLECTION_SERVICE_MASTERS"),
		os.Getenv("COLLECTION_PURCHASE_HISTORY"),
	}

	for _, collectionName := range allCollections {
		err := MongoInstance.CreateCollection(collectionName)
		if err != nil {
			if !strings.Contains(err.Error(), "already exists") {
				return nil, fmt.Errorf("error creating collection: %v", err)
			}
			log.Println("We have some collections that already exists: ", err)
		}
	}

	file, err := os.OpenFile("./../../logs/FiberLogs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	fiber := fiber.New()

	fiber.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	fiber.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
		Output:     file,
		TimeFormat: time.RFC3339Nano,
		TimeZone:   "Local",
	}))

	return &App{
		Fiber: fiber,
	}, nil
}

func (a *App) Close() {
	MongoInstance.Disconnect()
	a.Fiber.Shutdown()
}

func GetMongoInstance() *database.MongoDB {
	return MongoInstance
}
