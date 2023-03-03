package main

import (
	"fmt"

	"github.com/sashagitar/tgbotDime/tgbot"
)

const api_tg_my = "5710143526:AAH6SoHEK_eFQJm_CNCHciYSaMgJl0sa4MQ"
const api_tg = "5710143526:AAH6SoHEK_eFQJm_CNCHciYSaMgJl0sa4MQ"

func main() {
	fmt.Printf("%s\n", api_tg_my)

	// Создание бота
	bot, err := tgbot.Create(api_tg_my)
	if err != nil {
		panic(err)
	}
	// Запуск бота
	bot.Run()
}
