package intellect

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// состояние интелекта
type Params struct {
	id_user int64
	sost    []int
}

// Сообщение при нестандартном вводе
var errmsg = `Я тебя не понимаю`

// Приветственное сообщение
var help = `Привет! Меня зовут Движ. 
Я здесь для того, чтобы рассказать тебе про мой любимый отряд — сводный студенческий педагогический отряд «Движники».  

Про что ты хочешь узнать подробнее?`

// Создание интелекта
func Create(id_user int64) *Params {
	return &Params{id_user: id_user, sost: []int{0, 0}}
}

// Клавиатура кнопок главного меню
func (p *Params) baseKeyboard() *tgbotapi.ReplyKeyboardMarkup {
	k := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Что такое отряды?"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Что значит сводный отряд?"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Что такое педагогический отряд?"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Как твои дела, Движ?"),
		))
	return &k
}

// Возвращает информацию по поределённому отряду 3-ий уровень дерева 1 ветка
func (p *Params) getInfo(msg string) (string, *tgbotapi.ReplyKeyboardMarkup, error) {
	switch {
	case msg == "ССО":
		return `Я состою в педагогическом отряде, поэтому мало знаю про строительные. 
		
		Знаю, что они работают на стройках по всей России и иногда даже за её пределами. 
		Строительные отряды покоряют такие грандиозные объекты, как стройки проекта Мирный атом, Космодром Восточный. 
		Именно со Строительных отрядов началась история РСО. 
		Прежде чем выехать на целину, каждый отряд получает полезную строительную специальность, а потом работает в соответствии с нею`, nil, nil
	case msg == "СПО":
		return ``, nil, nil
	case msg == "ССхО":
		return `Я состою в педагогическом отряде, поэтому мало знаю про сельскохозяйственные. 
		
		Знаю, что сельскохозяйственные отряды работаю в полях, садах и на фермах по всей России. 
		Они работают и с растениями и с животными.`, nil, nil
	case msg == "СОП":
		return `Я состою в педагогическом отряде, поэтому мало знаю про отряды проводников. 
		
		Знаю, что они работают в поездах по всей России. 
		Без проводников не могла бы существовать железная дорога.`, nil, nil
	case msg == "серво":
		return `Я состою в педагогическом отряде, поэтому мало знаю про сервисные. 
		
		Знаю, что они работают по всей России. 
		Некоторые из них работают баристами, другие спасателями, третьи горничными, всего и не перечислишь.`, nil, nil
	case msg == "СМО":
		return `Я состою в педагогическом отряде, поэтому мало знаю про медицинские. 
		
		Знаю, что попасть туда могут только студенты медицинских вузов или имеющие медицинское образование. 
		Они работают в поликлинках и больницах по всей России.`, nil, nil
	case msg == "ТОП":
		return `Трудовые отряды подростков созданы для того, чтобы несовершеннолетние ребята, жаждущие трудиться, могли направить свою энергию. 
		Чаще всего ТОПы прикрепляются к какому-то линейному студенческому отряду и работают по его направлению. 
		У Сводного СПО  «Движники» есть такой младший брат - ТОП «Блеск»`, nil, nil
	case msg == "спец":
		return ``, nil, nil
	case msg == "А что там можно было спросить раньше?":
		p.sost[0] = 1
		p.sost[1] = 0

		return help, p.baseKeyboard(), nil
	}

	return errmsg, nil, nil
}

// Возвращает информацию о знамёнах 3-ий уровень дерева 2 ветка
func (p *Params) getInfoForZnam(msg string) (string, *tgbotapi.ReplyKeyboardMarkup, error) {
	if msg == "Какие знамёна есть у «Движников»?" {
		answ := `В 2020 году «Движники» получили знамя лучшего студенческого отряда Москвы по комиссарской деятельности

		В 2021 году «Движники» получили знамя лучшего студенческого отряда Москвы по совокупности показателей
		
		А в 2022 году «Движники» смогли взять два знамени - лучшего студенческого педагогического отряда Москвы и Лучшего студенческого отряда Москвы по совокупности показателей`

		return answ, nil, nil
	}
	if msg == "А что там можно было спросить раньше?" {
		p.sost[0] = 1
		p.sost[1] = 0

		return help, p.baseKeyboard(), nil
	}

	return help, nil, nil
}

// Воторой уровень дерева базовые вопросы и ответы
func (p *Params) lvl2(msg string) (string, *tgbotapi.ReplyKeyboardMarkup, error) {
	switch {
	case msg == "Расскажи больше про РСО":
		butKeyboard := tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Какие знамёна есть у «Движников»?"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("А что там можно было спросить раньше?"),
			),
		)
		answ := `Полное название организации - Молодёная 
		общероссийская общественная организация Российские студенческие отряды (МООО «РСО»). 
		Это крупнейшая молодёжная организация России, объединяющая более 150 тысяч молодых людей из 76 субъектов страны.
		
		РСО предоставляет студентом огромное количество возможностей для самореализации. 
		Ежегодно проводятся конкурсы профмастерства, а лучшие отряды получат знамёна (кстати, у «Движников» тоже есть знамёна). 
		Кроме профессиональных навыков можно испытать себя и в творчестве - на творческих конкурсах - и в спорте - на спартакиадах.`
		p.sost[0] = 3
		p.sost[1] = 2

		return answ, &butKeyboard, nil

	case msg == "Расскажи историю РСО":
		answ := ``

		return answ, nil, nil

	case msg == "Расскажи больше про твой отряд":
		answ := ``

		return answ, nil, nil

	case msg == "В каких лагерях работают Движники?":
		answ := ``

		return answ, nil, nil

	case msg == "Какие есть отряды?" || msg == "А какие ещё бывают отряды?":
		butKeyboard := tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("ССО"),
				tgbotapi.NewKeyboardButton("СПО"),
				tgbotapi.NewKeyboardButton("ССхО"),
				tgbotapi.NewKeyboardButton("СОП"),
				tgbotapi.NewKeyboardButton("Серво"),
				tgbotapi.NewKeyboardButton("СМО"),
				tgbotapi.NewKeyboardButton("ТОП"),
				tgbotapi.NewKeyboardButton("Спец"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("А что там можно было спросить раньше?"),
			),
		)
		answ := `В Российских студенческих отрядах существует 8 основных направлений:
			- Строительные отряды (ССО)
			- Педагогические отряды (СПО)
			- Сельскохозяйственные отряды (ССхО)
			- Отряды проводников (СОП)
			- - Сервисные отряды (Серво)
			- Медицинские отряды (СМО)
			- Трудовые отряды подростков (ТОП)
			- - Специализированные отряды (Спец)
			Каждое из направлений занимается чем-то своим.
			
			Я могу рассказать про каждое направление подробнее 
			Напиши их краткое название (то что в скобочках)`
		p.sost[0] = 3
		p.sost[1] = 1

		return answ, &butKeyboard, nil

	case msg == "А что там можно было спросить раньше?":
		p.sost[0] = 1
		p.sost[1] = 0

		return help, p.baseKeyboard(), nil
	}

	return errmsg, nil, nil
}

// Основна диалога
func (p *Params) GetAnswer(command string, msg string) (string, *tgbotapi.ReplyKeyboardMarkup, error) {
	// если уровень дерева 3 идём в нужную ветку
	if p.sost[0] == 3 {
		if p.sost[1] == 1 {
			return p.getInfo(msg)
		}
		if p.sost[1] == 2 {
			return p.getInfoForZnam(msg)
		}
	}
	// если уровень дерева 2 идём в древо диалогов второго уровня
	if p.sost[0] == 2 {
		return p.lvl2(msg)
	}

	// Главное меню
	if p.sost[0] == 1 {
		switch msg {
		// сравниваем ответ
		case "Что такое отряды?":
			// Создаём клавиатуру
			butKeyboard := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Расскажи больше про РСО"),
				),
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Расскажи историю РСО"),
				),
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Какие есть отряды?"),
				),
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("А что там можно было спросить раньше?"),
				),
			)
			// Записываем ответ
			answ := `В России существует организация, благодаря которой тысячи студентов могут найти себе работу на лето, компанию единомышленников, пространство для творечества и многое другое. 
				Эта организация называется Российские студенческие отряды (РСО). 
				Большая организация состоит из множества маленьких ячеек - линейных отрядов.`
			// меняем состояние дерева
			p.sost[0] = 2
			p.sost[1] = 1

			// возвращаем ответ и клавиатуру
			return answ, &butKeyboard, nil

		case "Что значит сводный отряд?":
			butKeyboard := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("А какие ещё бывают отряды?"),
				),
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Расскажи больше про твой отряд"),
				),
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("А что там можно было спросить раньше?"),
				),
			)
			answ := `Большинство существующих отрядов прикреплены к какой-либо образовательной организации - их штабу. 
				Однако кроме них существуют отряды без штаба. Такие отряды называются сводными и именно таким отрядом являются Движники.`
			p.sost[0] = 2
			p.sost[1] = 2

			return answ, &butKeyboard, nil

		case "Что такое педагогический отряд?":
			butKeyboard := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("А какие ещё бывают отряды?"),
				),
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("В каких лагерях работают Движники?"),
				),
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("А что там можно было спросить раньше?"),
				),
			)
			answ := `Существуют разные направления отрядов. Каждое направление занимается каким-то своим делом. 
				Участники педагогических отрядов работают с детьми. В основном в лагерях в самых разных уголках России. Вожатые, организаторы, сопровождающие. 
				Конечно же этим деятельность педагогических отрядов не ограничивается, но это, можно сказать, основные их задачи.`
			p.sost[0] = 2
			p.sost[1] = 3

			return answ, &butKeyboard, nil

		case "Как твои дела, Движ?":
			butKeyboard := tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Получить информацию"),
				),
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("А что там можно было спросить раньше?"),
				),
			)
			answ := `Спасибо, что спросил. 
				Мои дела прекрасны, а общение с тобой делает их ещё лучше. 
				Кстати, раз мы перешли к более тёплому общению, расскажи мне о себе`
			p.sost[0] = 2
			p.sost[1] = 4

			return answ, &butKeyboard, nil
		}

	}
	// Начало диалога // Главное меню
	if p.sost[0] == 0 {
		p.sost[0] = 1

		return help, p.baseKeyboard(), nil

	}
	return errmsg, nil, nil
}
