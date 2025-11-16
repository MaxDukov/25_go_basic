package main

// TODO
// 1. Добавить кнопку "Желания"
// 2. Написать функцию сохранения "Желаний" на диск и чтения их при старте бота, просто в файл
// 3. Написать функцию сохранения "Желаний" на диск и чтения их при старте бота, в sqlite
// 4. Написать логику разного приветствия для разных людей, хранить желания разных людей отдельно.

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	version = "1.0.0"
)

func main() {
	// Получаем токен бота из переменной окружения
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable is required")
	}

	// Создаем экземпляр бота
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Настраиваем обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Обрабатываем входящие сообщения
	for update := range updates {
		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}

		if update.Message != nil {
			handleMessage(bot, update.Message)
		}

		if update.CallbackQuery != nil {
			handleCallback(bot, update.CallbackQuery)
		}
	}
}

func handleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	// Создаем клавиатуру с кнопками
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("👋 Скажи привет", "say_hello"),
			tgbotapi.NewInlineKeyboardButtonData("ℹ️ Покажи версию", "show_version"),
		),
	)

	msg := tgbotapi.NewMessage(message.Chat.ID, "Выберите действие:")
	msg.ReplyMarkup = keyboard

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

func handleCallback(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	var responseText string

	switch callback.Data {
	case "say_hello":
		responseText = "Привет! 👋\nРад вас видеть!"
	case "show_version":
		responseText = "Версия бота: " + version
	default:
		responseText = "Неизвестная команда"
	}

	// Отправляем ответ на callback
	callbackConfig := tgbotapi.NewCallback(callback.ID, "")
	if _, err := bot.Request(callbackConfig); err != nil {
		log.Printf("Error sending callback response: %v", err)
	}

	// Отправляем сообщение с ответом
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, responseText)
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
