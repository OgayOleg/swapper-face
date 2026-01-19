package main

import (
	"faceSwapper/internal/adapter"
	"faceSwapper/internal/usecase"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Init and Load godotenv
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	telegramKey := os.Getenv("TELEGRAM_KEY")
	wModelKey := os.Getenv("WMODEL_KEY")

	// Telegram Bot API
	bot, err := tgbotapi.NewBotAPI(telegramKey)
	if err != nil {
		log.Fatal(err)
	}

	// Wmodel adapter
	wmodel := adapter.New(wModelKey)

	// usecase (service)
	usecase := usecase.New(bot, wmodel)
	usecase.Start()
}
