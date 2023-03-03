package tgbot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sashagitar/tgbotDime/tgbot/intellect"
)

// Интерфейс интелекта
type com interface {
	GetAnswer(command string, msg string) (*string, *tgbotapi.ReplyKeyboardMarkup, *tgbotapi.PhotoConfig, error)
}

// адаптер интелекта
type userIntellect struct {
	intellect com
}

// Бот
type Bot struct {
	users_bot_tupoi map[int]*userIntellect
	bot             *tgbotapi.BotAPI
}

// Создание бота
func Create(api string) (*Bot, error) {
	// Подключение бота
	bot, err := tgbotapi.NewBotAPI(api)
	if err != nil {
		log.Panic(err)
	}
	// Создание объекта бота
	b := Bot{}
	b.bot = bot

	// Подключение интелекта
	b.users_bot_tupoi = make(map[int]*userIntellect, 0)

	log.Printf("Authorized on account %s\n", bot.Self.UserName)

	return &b, nil
}

// Отправка сообщения пользователю
func (b *Bot) sendAnswer(id_user int, answer *string, buttonKeybord *tgbotapi.ReplyKeyboardMarkup, photo *tgbotapi.PhotoConfig) {
	if photo != nil {
		photo.Caption = *answer
		photo.ReplyMarkup = buttonKeybord
		if _, err := b.bot.Send(*photo); err != nil {
			log.Fatalln(err)
		}

	} else {
		msg := tgbotapi.NewMessage(int64(id_user), *answer)
		msg.ReplyMarkup = buttonKeybord
		b.bot.Send(msg)
	}

}

// Запуск бота
func (b *Bot) Run() {
	// Запуск ожидания сообщений от пользователя
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message == nil { // If we got a message
			continue
		}

		// Получение id  пользователя,команды и сообщения
		command := update.Message.Command()
		msg_user := update.Message.Text
		id_user := int(update.Message.From.ID)
		name_user := update.Message.From.UserName

		// Если пользователь написал впервые - создаём для него отдеьный объект интелекта
		if b.users_bot_tupoi[id_user] == nil {
			b.users_bot_tupoi[id_user] = &userIntellect{intellect: intellect.Create(int64(id_user))}
		}

		// Отправляем сообщение интелекту
		ans_wer, butKeyboard, photo, err := b.users_bot_tupoi[id_user].intellect.GetAnswer(command, msg_user)

		log.Printf("[%d] name %s, com %s, msg %s, %s", id_user, name_user, command, msg_user, err)

		b.sendAnswer(id_user, ans_wer, butKeyboard, photo)
	}
}
