package handler

import (
	"awesomeProject8/questions"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update, questionService *questions.Service) {
	if update.Message != nil && update.Message.Text == "/start" {
		sendQuestion(bot, update.Message.Chat.ID, 0, questionService)
	}

	if update.CallbackQuery != nil {
		processAnswer(bot, *update.CallbackQuery, questionService)
	}
}

func sendQuestion(bot *tgbotapi.BotAPI, chatID int64, index int, questionService *questions.Service) {
	// логика отправки вопроса
}

func processAnswer(bot *tgbotapi.BotAPI, callbackQuery tgbotapi.CallbackQuery, questionService *questions.Service) {
	// логика обработки ответа
}
