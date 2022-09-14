package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/chocomilk/app/settings"
	"github.com/gofiber/fiber/v2"
)

func main() {

	current_time := time.Now()

	fmt.Printf("%02d-%02d-%d %02d:%02d:%02d",
		current_time.Day(), current_time.Month(), current_time.Year(),
		current_time.Hour(), current_time.Minute(), current_time.Second())

	log.Println("Welcome to Chocomilk!")

	app := settings.SetupApp()
	InitializeServices(app)

	go func() {
		log.Println("Server started on port 8000")
		err := app.Listen(":8000")
		if err != nil {
			log.Fatal(err)
		}
	}()

	GracefulShutdown(app)

}

func InitializeServices(app *fiber.App) {

	api := app.Group("api/chocomilk")

	for _, service := range settings.Services {
		service.InitializeRouter(api)
		service.InitializeValidator()
		service.MakeMigrations()
	}

}

func GracefulShutdown(app *fiber.App) {

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Received terminate,graceful shutdown", sig)

	// database, err := connection.DB()
	// if err != nil {
	// 	fmt.Println("DB Closing ERROR :", err)
	// }
	// database.Close()
	// fmt.Println("DB has been closed")

	app.Shutdown()

}
