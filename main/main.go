package main

import (
	"awesomeProject8/handler"
	"awesomeProject8/questions"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	token := "7894387284:AAFOhRNGkiJOzUJCFS0_aCH9kiqfBbbIWeM"
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("Не получилось создать бота", err)
	}
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	questionService := questions.NewService()
	for update := range updates {
		handler.HandleUpdate(bot, update, questionService)
	}
}
