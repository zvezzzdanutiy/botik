package main

import (
	"log"
	"net/http"
	"newbot/internal/domain/AnekdotProviders"
	"newbot/internal/domain/TelegramBot"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Can't read env")
	}
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))
	JokeProvider := AnekdotProviders.New(
		&http.Client{},
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	go e.Start(":" + port)
	tgBot, err := TelegramBot.New(os.Getenv("BOT_TOKEN"), JokeProvider)
	if err != nil {
		// Обработайте ошибку
		log.Fatalf("Failed to create Telegram bot: %v", err)
	}
	tgBot.StartBot()
}
