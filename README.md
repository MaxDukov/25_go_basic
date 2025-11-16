Инструкция по запуску:

1. Создайте бота в Telegram:

Напишите @BotFather в Telegram

Используйте команду /newbot

Следуйте инструкциям и получите токен

2. Установите зависимости:

bash
go mod tidy

3. Запустите бота:

bash
export TELEGRAM_BOT_TOKEN="ваш_токен_бота"
go run main.go