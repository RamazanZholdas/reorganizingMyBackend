package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/RamazanZholdas/KeyboardistSV2/internal/app"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	app := app.New(os.Getenv("MONGO_URI"), os.Getenv("DATABASE_NAME"))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		app.Close()
	}()

	if err := app.Fiber.Listen(":8000"); err != nil {
		log.Panic(err)
	}

	fmt.Println("Running cleanup tasks...")
}
