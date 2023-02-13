package main

import (
	"fmt"

	"github.com/sashagitar/tgbotDime/tgbot"
)

const api_tg_my = " "
const api_tg = "5710143526:AAH6SoHEK_eFQJm_CNCHciYSaMgJl0sa4MQ"

func main() {
	fmt.Printf("%s\n", api_tg)

	// Создание бота
	bot, err := tgbot.Create(api_tg)
	if err != nil {
		panic(err)
	}
	// Запуск бота
	bot.Run()
}
