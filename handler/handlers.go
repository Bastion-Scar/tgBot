package handler

import (
	"awesomeProject8/questions"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

var userProgress = make(map[int64]int)
var userCorrectAnswers = make(map[int64]int)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update, questionService *questions.Service) {
	if update.Message != nil && update.Message.Text == "/start" {
		log.Println("Нажата кнопка start")
		chatID := update.Message.Chat.ID
		userProgress[chatID] = 0
		userCorrectAnswers[chatID] = 0
		sendQuestion(bot, chatID, questionService)
	}

	if update.CallbackQuery != nil {
		processAnswer(bot, *update.CallbackQuery, questionService)
	}
}
func sendQuestion(bot *tgbotapi.BotAPI, chatID int64, questionService *questions.Service) {
	// Проверяем текущий билет
	ticketIndex := userProgress[chatID]
	if ticketIndex >= len(questionService.Tickets) {
		msg := tgbotapi.NewMessage(chatID, "Вы завершили все билеты этой Z бабки! /start нажмите для перезапуска. Создатель бота - @SlaughterContinues (Бот не логирует ваши данные, я не вижу кто это проходит)")
		bot.Send(msg)
		return
	}

	// Получаем текущий билет и вопрос
	ticket := questionService.Tickets[ticketIndex]
	qIndex := userCorrectAnswers[chatID]

	// Если все вопросы в текущем билете пройдены, переходим к следующему билету
	if qIndex >= len(ticket.Questions) {
		userProgress[chatID]++
		sendQuestion(bot, chatID, questionService) // переходим к следующему билету
		return
	}

	// Получаем вопрос по индексу
	q := ticket.Questions[qIndex]

	// 🔢 Формируем текст вопроса + варианты
	messageText := q.Text + "\n\n"
	for i, option := range q.Options {
		messageText += strconv.Itoa(i+1) + ". " + option + "\n"
	}

	// 🎛 Кнопки: просто 1, 2, 3, 4
	var buttons []tgbotapi.InlineKeyboardButton
	for i := range q.Options {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(i+1), strconv.Itoa(i)))
	}
	keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons)

	// 📤 Отправка вопроса с вариантами ответов
	msg := tgbotapi.NewMessage(chatID, messageText)
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}

func processAnswer(bot *tgbotapi.BotAPI, callbackQuery tgbotapi.CallbackQuery, questionService *questions.Service) {
	chatID := callbackQuery.Message.Chat.ID
	ticketIndex := userProgress[chatID]
	ticket := questionService.Tickets[ticketIndex]

	qIndex := userCorrectAnswers[chatID]
	q := ticket.Questions[qIndex]

	userAnswerIndex, _ := strconv.Atoi(callbackQuery.Data)
	correct := userAnswerIndex == q.Answer

	var reply string
	if correct {
		reply = "✅ Верно"
		userCorrectAnswers[chatID]++ // Увеличиваем индекс правильных ответов
	} else {
		reply = "❌ Неверно. Правильный ответ: " + q.Options[q.Answer]
		userCorrectAnswers[chatID]++ // Увеличиваем индекс правильных ответов
	}

	msg := tgbotapi.NewMessage(chatID, reply)
	bot.Send(msg)

	// Убираем старые кнопки
	edit := tgbotapi.NewEditMessageReplyMarkup(
		chatID,
		callbackQuery.Message.MessageID,
		tgbotapi.InlineKeyboardMarkup{},
	)
	bot.Send(edit)

	// Если все вопросы в билете завершены, переходим к следующему билету
	if userCorrectAnswers[chatID] >= len(ticket.Questions) {
		userProgress[chatID]++         // Переход к следующему билету
		userCorrectAnswers[chatID] = 0 // Сброс правильных ответов
	}

	// Отправляем следующий вопрос
	sendQuestion(bot, chatID, questionService)
}
