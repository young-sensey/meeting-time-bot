package main

import (
	"github.com/go-co-op/gocron"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := gocron.NewScheduler(time.UTC)
	//s.Cron("* * * * *").Do(meetingTime) // каждая минута
	s.Cron("0 0 * * *").Do(meetingTime) // в 12:00
	s.StartBlocking()
}

func meetingTime() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	chatId, _ := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)

	msg := tgbotapi.NewMessage(chatId, "meeting time")
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
