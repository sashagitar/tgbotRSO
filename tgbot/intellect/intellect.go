package intellect

// переделать ответ + кнопки в формат массива
// получить информацию - сохранение введённого сообщения

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// состояние интелекта
type Params struct {
	id_user  int64
	sost     []int
	tree     treebr
	pred     string
	predSost int
}

type answer struct {
	text     string
	keyboard tgbotapi.ReplyKeyboardMarkup
	photo    tgbotapi.PhotoConfig
	toSost   int
	predSost int
	predMsg  string
}

type treebr struct {
	br1 map[string]answer
	br2 map[string]answer
	br3 map[string]answer
	br4 map[string]answer
	br5 map[string]answer
	br6 map[string]answer
}

// Сообщение при нестандартном вводе
var errmsg = answer{
	text: `К сожалению, ответа на этот вопрос я пока не знаю. Попробуй задать его своим кураторам или командному составу. Уверен, они с радостью помогут тебе!`,
	keyboard: tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Что такое отряды?"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Что значит сводный отряд?"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Что такое педагогический отряд?"),
		)),
}

// Создание интелекта
func Create(id_user int64) *Params {
	p := &Params{
		id_user: id_user,
		sost:    []int{0, 0},
	}
	tree := treebr{
		br1: map[string]answer{
			"Что такое отряды?": answer{
				text: `В России существует организация, благодаря которой тысячи студентов могут найти себе работу на лето, компанию единомышленников, пространство для творчества и многое другое. 
Эта организация называется Российские студенческие отряды (РСО). 
Большая организация состоит из множества маленьких ячеек - линейных отрядов.`, //RSD = Russian students detouchments - РСО
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Расскажи больше про РСО"),
					),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Расскажи историю РСО"),
					),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Какие есть отряды?"),
					)),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/Dvij3.png")),
				toSost:   2,
				predSost: 0,
			},
			"Что значит сводный отряд?": answer{
				text: `Большинство существующих отрядов прикреплены к какой-либо образовательной организации - их штабу. 
Однако кроме них существуют отряды без штаба. Такие отряды называются сводными и именно таким отрядом являются Движники.`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Какие есть отряды?"),
					),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Расскажи больше про твой отряд"),
					),
				),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/DvijJumps.png")),
				toSost:   3,
				predSost: 0,
			},
			"Что такое педагогический отряд?": answer{
				text: `Педагогические отряды занимают самое большое место в моём электронном сердечке.
		
Участники педагогических отрядов работают с детьми. В основном в лагерях в самых разных уголках России. Вожатые, организаторы, сопровождающие. 
Конечно же этим деятельность педагогических отрядов не ограничивается, но это, можно сказать, основные их задачи.`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССО"),
						tgbotapi.NewKeyboardButton("СПО"),
						tgbotapi.NewKeyboardButton("ССхО"),
						tgbotapi.NewKeyboardButton("СОП")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССервО"),
						tgbotapi.NewKeyboardButton("СМО"),
						tgbotapi.NewKeyboardButton("ТОП"),
						tgbotapi.NewKeyboardButton("ПСО")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Расскажи ещё раз, что такое отряды?"),
					)),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/DvijMic.png")),
				toSost:   5,
				predSost: 0,
			},
		},
		br2: map[string]answer{
			"Какие есть отряды?": answer{
				text: `В Российских студенческих отрядах существует 8 основных направлений:
- Строительные отряды (ССО)
- Педагогические отряды (СПО)
- Сельскохозяйственные отряды (ССхО)
- Отряды проводников (СОП)
- Сервисные отряды (Серво)
- Медицинские отряды (СМО)
- Трудовые отряды подростков (ТОП)
- Профильные студенческие отряды(ПСО)
Каждое из направлений занимается чем-то своим.
						
Я могу рассказать про каждое направление немного подробнее 
Напиши их краткое название (то что в скобочках)`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССО"),
						tgbotapi.NewKeyboardButton("СПО"),
						tgbotapi.NewKeyboardButton("ССхО"),
						tgbotapi.NewKeyboardButton("СОП")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССервО"),
						tgbotapi.NewKeyboardButton("СМО"),
						tgbotapi.NewKeyboardButton("ТОП"),
						tgbotapi.NewKeyboardButton("ПСО")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Расскажи ещё раз, что такое отряды?"),
					)),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/DvijReads.png")),
				toSost:   5,
				predSost: 1,
				predMsg:  "Что такое отряды?",
			},
			"Расскажи историю РСО": answer{
				text: `История Российских студенческих отрядов берёт своё начало летом 1959 года, когда 339 студентов физического факультета МГУ отправились покорять целину в Казахстане, где построили 16 объектов.
					
С тех пор организация начала стремительно развиваться. Уже через 5 лет в 1964 году в строительных отрядах было уже 30 тысяч молодых энтузиастов, представителей 9 союзных республик, 41 города, 178 высших учебных заведений. Было построено 3860 объектов, организовано более 3 тысяч концертов, прочитано 5 тысяч лекций для сельских тружеников.
		
На пике своего развития в Советский период Российские студенческие отряды стали домом для  более 860 тысяч человек. А всего с 1959 по 1991 года частью РСО побывали более 13 миллионов человек. 
		
Движение возродилось в 2004 году. В 2015 году указом президента РФ В. Путина был установлен профессиональный праздник - День Российских Студенческих Отрядов, отмечаемый 17 февраля.`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Круть!")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Расскажи ещё раз, что такое отряды?"),
					)),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/DvijBook.png")),
				toSost:   0,
				predSost: 1,
				predMsg:  "Что такое отряды?",
			},
			"Расскажи больше про РСО": answer{
				text: `Полное название организации - Молодёжная общероссийская общественная организация Российские студенческие отряды (МООО «РСО»). Это крупнейшая молодёжная организация России, объединяющая более 150 тысяч молодых людей из 76 субъектов страны.
		
РСО предоставляет студентом огромное количество возможностей для самореализации. Ежегодно проводятся конкурсы профмастерства, а лучшие отряды получат знамёна (кстати, у «Движников» тоже есть знамёна). Кроме профессиональных навыков можно испытать себя и в творчестве - на творческих конкурсах - и в спорте - на спартакиадах.`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Какие знамёна есть у Движников?")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("А что вообще такое знамя?"))),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/DvijMistery.png")),
				toSost:   3,
				predSost: 1,
				predMsg:  "Что такое отряды?",
			},
		},
		br3: map[string]answer{ //слишком много текста, он столько не хочет отправлять. Я убрал что посчитал лишним и всё заработало
			"Расскажи больше про твой отряд": answer{
				text: `Сводный студенческий педагогический отряд был создан в 2017 году студентами, выезжающими вожатыми в лагерь «Бугорок». Днём рождения отряда считается 5 марта.
Первым командиром СПО «Движники» стала Белокрылова Наталья Витальевна. С тех пор каждый год бойцы отряда вместе выбирают комсостав.
				
Движники становятся организаторами на многих массовых мерпориятиях, в которых принимает участие РСО.
Например, многие Движники принимают участие в организации Школы вожатых, в организации форумов. 
		
Каждую зиму бойцы отряда выезжают участниками Всероссийской патриотической акции Снежный десант РСО.
На данный момент наши бойцы состоят в отрядах Снежного десанта «Клюква», «Пламя», «Буря», «Белое золото»
				
И всё это - только верхушка айсберга. Гораздо больше ты узнаешь, когда станешь частью нашей семьи. А пока можешь изучить нашу страничку в социальной сети Вконтакте https://vk.com/spodvijnik1 или в мессенджере телеграм t.me/dvizhniki.`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("В каких лагерях работают Движники?")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Какие знамёна есть у Движников?")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("А что значит комсостав?"))),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/DvijLoves.png")),
				toSost:   4,
				predSost: 1,
				predMsg:  "Что значит сводный отряд?",
			},
			"Какие знамёна есть у Движников?": answer{
				text: `В 2020 году «Движники» получили знамя лучшего студенческого отряда Москвы по комиссарской деятельности
		
В 2021 году «Движники» получили знамя лучшего студенческого отряда Москвы по совокупности показателей
					
А в 2022 году «Движники» смогли взять два знамени - лучшего студенческого педагогического отряда Москвы и Лучшего студенческого отряда Москвы по совокупности показателей`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("А что вообще такое знамя?"))),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/Znamea.jpg")),
				toSost:   4,
				predSost: 1,
				predMsg:  "Расскажи больше про РСО",
			},
			"Какие есть отряды?": answer{
				text: `В Российских студенческих отрядах существует 8 основных направлений:
- Строительные отряды (ССО)
- Педагогические отряды (СПО)
- Сельскохозяйственные отряды (ССхО)
- Отряды проводников (СОП)
- Сервисные отряды (Серво)
- Медицинские отряды (СМО)
- Трудовые отряды подростков (ТОП)
- Профильные студенческие отряды(ПСО)
Каждое из направлений занимается чем-то своим.
						
Я могу рассказать про каждое направление немного подробнее 
Напиши их краткое название (то что в скобочках)`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССО"),
						tgbotapi.NewKeyboardButton("СПО"),
						tgbotapi.NewKeyboardButton("ССхО"),
						tgbotapi.NewKeyboardButton("СОП")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССервО"),
						tgbotapi.NewKeyboardButton("СМО"),
						tgbotapi.NewKeyboardButton("ТОП"),
						tgbotapi.NewKeyboardButton("ПСО")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Расскажи ещё раз, что такое отряды?"),
					)),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/DvijReads.png")),
				toSost:   5,
				predSost: 1,
				predMsg:  "Что значит сводный отряд?",
			},
			"А что вообще такое знамя?": answer{
				text: `Знамя - это полотнище на древке со шпилем и подставкой, торжественная награда, символизирующая большую заслугу в общественной организации. В МООО "РСО" знамя - это главная награда лучшим отрядам.

На древке знамени располагаются набойки, показывающие, какой отряд в какой год был удостоен этого знака почёта.
				
Знамя лучшего отряда МосРСО по совокупности показателей (т.е. знамя лучшего отряда Москвы) - награда, которая выдается отряду, набравшему по результатам всех конкурсных мероприятий регионального отделения и сданных отчетов наибольшее количество баллов.
				
Знамя лучшего отряда МосРСО по комиссарской деятельности забирает ЛСО, который набрал максимальное количество баллов за отчетную часть.
				
Знамя лучшего педагогического отряда получает СПО, победивший в региональном конкурсе профессионального вожатского мастерства`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Круть!")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Расскажи ещё раз, что такое отряды?"),
					)),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/DvijReads.png")),
				toSost:   0,
				predSost: 2,
				predMsg:  "Расскажи больше про РСО",
			},
		},
		//Было бы круто сделать гиперссылки. Задача на будущее
		br4: map[string]answer{
			"А что значит комсостав?": answer{
				text: `В каждом отряде есть свой командный состав.
Командир – руководитель студенческого отряда. Командир отряда организует работу студенческого отряда, несёт персональную ответственность перед соответствующим Штабом, за производственную, общественную и финансово-хозяйственную деятельность студенческого отряда, обеспечивает здоровые и безопасные условия труда, внутриотрядную дисциплину.
				
Командир нашего отряда: Шелепа Марина (https://vk.com/m.sh13).
				
Комиссар – заместитель командира, который несет ответственность за организацию внутриотрядной жизни, следит за настроением внутренним эмоциональным климатом отряда.
				
Комиссар нашего отряда: Казмалы Даша (https://vk.com/dari_kaz)
				
Методист - заместитель командира по производственной деятельности студенческого отряда. Должен обладать необходимым профессиональным опытом работы по профилю деятельности отряда.
				
Методист нашего отряда: Пашечко Женя (https://vk.com/idpashechko)
				
	Также в отряде есть должности руководителя PR-службы, медика и флагоносца.`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("А что насчёт медика, руководителя PR и флагоносца?"))),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/DvijWeCanDoIt.png")),
				toSost:   5,
				predSost: 3,
				predMsg:  "Расскажи больше про твой отряд",
			},
			"В каких лагерях работают Движники?": answer{
				text: `Движники работают во многих лагерях. Этим летом, например, мы выезжали в «Бугорок», «Салют», «Планета Английского», «SportZania», «Химик» `,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Круть!")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Расскажи ещё раз про твой отряд!"))),
				photo:  tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/DvijWeCanDoIt.png")),
				toSost: 0,
			},
			"Какие знамёна есть у Движников?": answer{
				text: `В 2020 году «Движники» получили знамя лучшего студенческого отряда Москвы по комиссарской деятельности
		
В 2021 году «Движники» получили знамя лучшего студенческого отряда Москвы по совокупности показателей
					
А в 2022 году «Движники» смогли взять два знамени - лучшего студенческого педагогического отряда Москвы и Лучшего студенческого отряда Москвы по совокупности показателей`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Круть!")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Расскажи ещё раз про твой отряд!"))),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/Znamea.jpg")),
				toSost:   0,
				predSost: 3,
				predMsg:  "Расскажи больше про твой отряд",
			},
			"А что вообще такое знамя?": answer{
				text: `Знамя - это полотнище на древке со шпилем и подставкой, торжественная награда, символизирующая большую заслугу в общественной организации. В МООО "РСО" знамя - это главная награда лучшим отрядам.

На древке знамени располагаются набойки, показывающие, какой отряд в какой год был удостоен этого знака почёта.
				
Знамя лучшего отряда МосРСО по совокупности показателей (т.е. знамя лучшего отряда Москвы) - награда, которая выдается отряду, набравшему по результатам всех конкурсных мероприятий регионального отделения и сданных отчетов наибольшее количество баллов.
				
Знамя лучшего отряда МосРСО по комиссарской деятельности забирает ЛСО, который набрал максимальное количество баллов за отчетную часть.
				
Знамя лучшего педагогического отряда получает СПО, победивший в региональном конкурсе профессионального вожатского мастерства`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Круть!")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Расскажи ещё раз про твой отряд!"))),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/DvijReads.png")),
				toSost:   0,
				predSost: 3,
				predMsg:  "Какие знамёна есть у Движников?",
			},
		},

		br5: map[string]answer{
			"А что насчёт медика, руководителя PR и флагоносца?": answer{
				text: `Медики особенно необходимы во время целины и отрядных мероприятий. Медиком становится человек, разбирающийся в медицине, способный оказать помощь в случае необходимости. У медика хранится отрядная аптечка. Отличить отрядного медика можно по нашивке - у него на ней изображён медицинский крест.
				
Медиком нашего отряда является Олеся Белкина.
				
На руководителе PR отдела лежит ответственность за ведение социальных сетей отряда, за освещение деятельности отряда, за поддержание стиля отряда. 

Руководителем PR отдела Движников является Лена Кириленко. 

Флагоносец отряда ответственен за сохранение флага и флагштока. Флагоносец является лицом отряда на крупных мероприятиях, гордо держа флаг. 

Флагоносцем отряда является Дима Горобец`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Круть!")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Расскажи ещё раз про твой отряд!"))),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/DvijTV.png")),
				toSost:   0,
				predSost: 3,
			},
			"ССО": answer{
				text: `Я состою в педагогическом отряде, поэтому мало знаю про строительные. 
					
Знаю, что они работают на стройках по всей России и иногда даже за её пределами. 
Строительные отряды покоряют такие грандиозные объекты, как стройки проекта Мирный атом, Космодром Восточный. 
Именно со Строительных отрядов началась история РСО. 
Прежде чем выехать на целину, каждый отряд получает полезную строительную специальность, а потом работает в соответствии с нею`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССО"),
						tgbotapi.NewKeyboardButton("СПО"),
						tgbotapi.NewKeyboardButton("ССхО"),
						tgbotapi.NewKeyboardButton("СОП")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССервО"),
						tgbotapi.NewKeyboardButton("СМО"),
						tgbotapi.NewKeyboardButton("ТОП"),
						tgbotapi.NewKeyboardButton("ПСО")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Расскажи ещё раз, что такое отряды?"),
					)),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/Dvij3.png")),
				toSost:   5,
				predSost: 4,
			},
			"СПО": answer{
				text: `Педагогические отряды занимают самое большое место в моём электронном сердечке.
		
Участники педагогических отрядов работают с детьми. В основном в лагерях в самых разных уголках России. Вожатые, организаторы, сопровождающие. 
Конечно же этим деятельность педагогических отрядов не ограничивается, но это, можно сказать, основные их задачи.`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССО"),
						tgbotapi.NewKeyboardButton("СПО"),
						tgbotapi.NewKeyboardButton("ССхО"),
						tgbotapi.NewKeyboardButton("СОП")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССервО"),
						tgbotapi.NewKeyboardButton("СМО"),
						tgbotapi.NewKeyboardButton("ТОП"),
						tgbotapi.NewKeyboardButton("ПСО")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Расскажи ещё раз, что такое отряды?"),
					)),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/DvijMic.png")),
				toSost:   5,
				predSost: 4,
			},
			"СОП": answer{
				text: `Я состою в педагогическом отряде, поэтому мало знаю про отряды проводников. 
					
Знаю, что они работают в поездах по всей России. 
Без проводников не могла бы существовать железная дорога.`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССО"),
						tgbotapi.NewKeyboardButton("СПО"),
						tgbotapi.NewKeyboardButton("ССхО"),
						tgbotapi.NewKeyboardButton("СОП")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССервО"),
						tgbotapi.NewKeyboardButton("СМО"),
						tgbotapi.NewKeyboardButton("ТОП"),
						tgbotapi.NewKeyboardButton("ПСО")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Расскажи ещё раз, что такое отряды?"),
					)),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/Dvij5.png")),
				toSost:   5,
				predSost: 4,
			},
			"ССервО": answer{
				text: `Я состою в педагогическом отряде, поэтому мало знаю про сервисные. 
					
Знаю, что они работают по всей России. 
Некоторые из них работают бариста, другие спасателями, третьи горничными, поворами, официантами, всего и не перечислишь.`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССО"),
						tgbotapi.NewKeyboardButton("СПО"),
						tgbotapi.NewKeyboardButton("ССхО"),
						tgbotapi.NewKeyboardButton("СОП")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССервО"),
						tgbotapi.NewKeyboardButton("СМО"),
						tgbotapi.NewKeyboardButton("ТОП"),
						tgbotapi.NewKeyboardButton("ПСО")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Расскажи ещё раз, что такое отряды?"),
					)),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/DvijEats.png")),
				toSost:   5,
				predSost: 4,
			},
			"ССхО": answer{
				text: `Я состою в педагогическом отряде, поэтому мало знаю про сельскохозяйственные. 
					
Знаю, что сельскохозяйственные отряды работаю в полях, садах и на фермах по всей России. 
Они работают и с растениями, и с животными.`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССО"),
						tgbotapi.NewKeyboardButton("СПО"),
						tgbotapi.NewKeyboardButton("ССхО"),
						tgbotapi.NewKeyboardButton("СОП")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССервО"),
						tgbotapi.NewKeyboardButton("СМО"),
						tgbotapi.NewKeyboardButton("ТОП"),
						tgbotapi.NewKeyboardButton("ПСО")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Расскажи ещё раз, что такое отряды?"),
					)),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/DvijGlasses.png")),
				toSost:   5,
				predSost: 4,
			},
			"СМО": answer{
				text: `Я состою в педагогическом отряде, поэтому мало знаю про медицинские. 
					
Знаю, что попасть туда могут только студенты медицинских вузов или имеющие медицинское образование. 
Они работают в поликлинках и больницах по всей России, медиками на различных трудовых проектах и много где ещё.`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССО"),
						tgbotapi.NewKeyboardButton("СПО"),
						tgbotapi.NewKeyboardButton("ССхО"),
						tgbotapi.NewKeyboardButton("СОП")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССервО"),
						tgbotapi.NewKeyboardButton("СМО"),
						tgbotapi.NewKeyboardButton("ТОП"),
						tgbotapi.NewKeyboardButton("ПСО")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Расскажи ещё раз, что такое отряды?"),
					)),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/DvijSickness.png")),
				toSost:   5,
				predSost: 4,
			},
			"ТОП": answer{
				text: `Трудовые отряды подростков созданы для того, чтобы несовершеннолетние ребята, жаждущие трудиться, могли направить свою энергию. 
Чаще всего ТОПы прикрепляются к какому-то линейному студенческому отряду и работают по его направлению. 
У Сводного СПО  «Движники» есть такой младший брат - ТОП «Блеск»`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССО"),
						tgbotapi.NewKeyboardButton("СПО"),
						tgbotapi.NewKeyboardButton("ССхО"),
						tgbotapi.NewKeyboardButton("СОП")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССервО"),
						tgbotapi.NewKeyboardButton("СМО"),
						tgbotapi.NewKeyboardButton("ТОП"),
						tgbotapi.NewKeyboardButton("ПСО")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Расскажи ещё раз, что такое отряды?"),
					)),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/DvijLaugh.png")),
				toSost:   5,
				predSost: 4,
			},
			"ПСО": answer{
				text: `К профильным студенческим отрядам относятся отряды археологов, энергетические отряды, путийные отряды. 
					
Каждое направление профильных отрядов занимается чем-то особенным. Раньше медицинские отряды тоже были профильным направлением, но со временем количество бойцов СМО увеличилось и их вывели в отдельное направление.`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССО"),
						tgbotapi.NewKeyboardButton("СПО"),
						tgbotapi.NewKeyboardButton("ССхО"),
						tgbotapi.NewKeyboardButton("СОП")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("ССервО"),
						tgbotapi.NewKeyboardButton("СМО"),
						tgbotapi.NewKeyboardButton("ТОП"),
						tgbotapi.NewKeyboardButton("ПСО")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Расскажи ещё раз, что такое отряды?"),
					)),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/Dvij6.png")),
				toSost:   5,
				predSost: 4,
			},
		},
		br6: map[string]answer{
			"Всё плохо, давай сначала": answer{

				text: `Привет! Меня зовут Движ. 
Я здесь для того, чтобы рассказать тебе про мой любимый отряд — Сводный студенческий педагогический отряд «Движники».  
		
Про что ты хочешь узнать подробнее?`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Что такое отряды?"),
					),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Что значит сводный отряд?"),
					),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Что такое педагогический отряд?"),
					)),
				photo:  tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/DvijHello.png")),
				toSost: 1,
			},
			"Круть!": answer{

				text: `Это, может быть и круто, но знай, что ты ещё круче!`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Спасибо за комплимент, Движ!"))),
				photo:  tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/DvijWink.png")),
				toSost: 1,
			},
			"Спасибо за комплимент, Движ!": answer{

				text: `Ну а теперь я готов рассказать тебе ещё немного про мою любимую организацию. Что тебе интересно?`,
				keyboard: tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Что такое отряды?"),
					),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Что значит сводный отряд?"),
					),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Что такое педагогический отряд?"),
					)),
				photo:    tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("tgbot/intellect/photos/DvijHides.png")),
				toSost:   1,
				predSost: 0,
			},
		},
	}
	p.tree = tree
	return p
}

// основа диалога
func (p *Params) GetAnswer(command string, msg string) (*string, *tgbotapi.ReplyKeyboardMarkup, *tgbotapi.PhotoConfig, error) {
	if command == "back" {
		p.sost[0] = p.predSost
		msg = p.pred
	}

	if command == "start" {
		p.sost[0] = 0
		msg = "Всё плохо, давай сначала"
	}

	if msg == "Круть!" {
		p.sost[0] = 0
		msg = "Круть!"
	}

	if msg == "Расскажи ещё раз, что такое отряды?" {
		p.sost[0] = 1
		msg = "Что такое отряды?"
	}

	if msg == "Расскажи ещё раз про твой отряд!" {
		p.sost[0] = 3
		msg = "Расскажи больше про твой отряд"
	}

	if msg == "Спасибо за комплимент, Движ!" {
		p.sost[0] = 0
		msg = "Спасибо за комплимент, Движ!"
	}

	if p.sost[0] == 5 {
		if ans, ok := p.tree.br5[msg]; ok {
			p.predSost = ans.predSost
			p.pred = ans.predMsg
			p.sost[0] = ans.toSost
			return &ans.text, &ans.keyboard, &ans.photo, nil
		}
	}

	if p.sost[0] == 4 {
		if ans, ok := p.tree.br4[msg]; ok {
			p.predSost = ans.predSost
			p.pred = ans.predMsg
			p.sost[0] = ans.toSost
			return &ans.text, &ans.keyboard, &ans.photo, nil
		}
	}

	if p.sost[0] == 3 {
		if ans, ok := p.tree.br3[msg]; ok {
			p.predSost = ans.predSost
			p.pred = ans.predMsg
			p.sost[0] = ans.toSost
			return &ans.text, &ans.keyboard, &ans.photo, nil
		}
	}

	// если уровень дерева 2 идём в древо диалогов второго уровня
	if p.sost[0] == 2 {
		if ans, ok := p.tree.br2[msg]; ok {
			p.predSost = ans.predSost
			p.pred = ans.predMsg
			p.sost[0] = ans.toSost
			return &ans.text, &ans.keyboard, &ans.photo, nil
		}
	}

	// Главное меню
	if p.sost[0] == 1 {
		if ans, ok := p.tree.br1[msg]; ok {
			p.predSost = ans.predSost
			p.pred = ans.predMsg
			p.sost[0] = ans.toSost
			return &ans.text, &ans.keyboard, &ans.photo, nil
		}
	}
	// Начало диалога
	if p.sost[0] == 0 {
		if ans, ok := p.tree.br6[msg]; ok {
			p.predSost = 0
			p.sost[0] = ans.toSost
			return &ans.text, &ans.keyboard, &ans.photo, nil
		}

	}

	return &errmsg.text, nil, nil, nil
}
