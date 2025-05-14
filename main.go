package main

import (
	"awesomeProject8/handler"
	"awesomeProject8/questions"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("7894387284:AAFOhRNGkiJOzUJCFS0_aCH9kiqfBbbIWeM")
	if err != nil {
		log.Panic(err) // Если не удалось создать бота, выводим ошибку
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u) // Канал для получения обновлений

	questionService := questions.NewService() // Создаем сервис для работы с вопросами

	// Бесконечный цикл для обработки обновлений
	for update := range updates {
		handler.HandleUpdate(bot, update, questionService) // Перехватываем обновления и передаем в обработчик
	}
}
