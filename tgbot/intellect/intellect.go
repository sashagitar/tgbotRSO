package intellect

// переделать ответ + кнопки в формат массива
// получить информацию - сохранение введённого сообщения

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// состояние интелекта
type Params struct {
	id_user int64
	sost    []int
	logica  logical
}

type answer struct {
	text     string
	keyboard tgbotapi.ReplyKeyboardMarkup
	photo    *tgbotapi.PhotoConfig
}

type logical struct {
	br01 map[string]answer
	br02 map[string]answer
	br03 map[string]answer
	br04 map[string]answer
	br05 map[string]answer
	br06 map[string]answer
	br07 map[string]answer
	br08 map[string]answer
	br09 map[string]answer
	br10 map[string]answer
	br11 map[string]answer
	br12 map[string]answer
	br13 map[string]answer
	br14 map[string]answer
	br15 map[string]answer
	br16 map[string]answer
	br17 map[string]answer
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
	photo: &tgbotapi.NewPhoto(update.Message.From.ID, tgbotapi.FilePath("photos/DvijFace1.png"))}

// Сообщения-ответы:
var help = answer{

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
	photo:/*&tgbotapi.NewPhoto(update.Message.From.ID, tgbotapi.FilePath("maxresdefault.jpg"))*/ nil}
var SSO = answer{
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
			tgbotapi.NewKeyboardButton("ПСО"),
		)),
	photo: nil}

var SPO = answer{
	text: `Педагогческие отряды занимают самое большое место в моём электронном сердечке.

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
			tgbotapi.NewKeyboardButton("ПСО"),
		)),
	photo: nil}
var SOP = answer{
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
			tgbotapi.NewKeyboardButton("ПСО"),
		)),
	photo: nil}
var SShO = answer{
	text: `Я состою в педагогическом отряде, поэтому мало знаю про сельскохозяйственные. 
		
 Знаю, что сельскохозяйственные отряды работаю в полях, садах и на фермах по всей России. 
 Они работают и с растениями и с животными.`,
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
			tgbotapi.NewKeyboardButton("ПСО"),
		)),
	photo: nil}
var SServO = answer{
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
			tgbotapi.NewKeyboardButton("ПСО"),
		)),
	photo: nil}
var SMO = answer{
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
			tgbotapi.NewKeyboardButton("ПСО"),
		)),
	photo: nil}
var TOP = answer{
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
			tgbotapi.NewKeyboardButton("ПСО"),
		)),
	photo: nil}
var PSO = answer{
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
			tgbotapi.NewKeyboardButton("ПСО"),
		)),
	photo: nil}
var Znamea = answer{
	text: `В 2020 году «Движники» получили знамя лучшего студенческого отряда Москвы по комиссарской деятельности

 В 2021 году «Движники» получили знамя лучшего студенческого отряда Москвы по совокупности показателей
		
 А в 2022 году «Движники» смогли взять два знамени - лучшего студенческого педагогического отряда Москвы и Лучшего студенческого отряда Москвы по совокупности показателей`,
	keyboard: tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Круть!"))),
	photo: nil}
var History = answer{
	text: `История Российских студенческих отрядов берёт своё начало летом 1959 года, когда 339 студентов физического факультета МГУ отправились покорять целину в Казахстане, где построили 16 объектов.
		
 С тех пор организация начала стремительно развиваться. Уже через 5 лет в 1964 году в строительных отрядах было уже 30 тысяч молодых энтузиастов, представителей 9 союзных республик, 41 города, 178 высших учебных заведений. Было построено 3860 объектов, организовано более 3 тысяч концертов, прочитано 5 тысяч лекций для сельских тружеников.

 На пике своего развития в Советский период Российские студенческие отряды стали домом для  более 860 тысяч человек. А всего с 1959 по 1991 года частью РСО побывали более 13 миллионов человек. 

 Движение возродилось в 2004 году. В 2015 году указом президента РФ В. Путина был установлен профессиональный праздник - День Российских Студенческих Отрядов, отмечаемый 17 февраля.`,
	keyboard: tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Круть!"))),
	photo: nil}
var Dvijniki = answer{
	text: `Сводный студенческий педагогический отряд был создан в 2017 году студентами, выезжающими вожатыми в лагерь «Бугорок». Днём рождения отряда считается 5 марта.
 Первым командиром СПО «Движники» стала Белокрылова Наталья Витальевна.
		
 Сегодня в отряде состоит 25 человек.
 Командир отряда - Шелепа Марина.
 Комиссар отряда - Казмалы Даша.
 Методист отряда - Пашечко Женя. 
 Медик отряда - Белкина Олеся.
 Руководитель PR отдела - Кириленко Лена.
	
 Движники становятся организаторами на многих массовых мерпориятиях, в которых принимает участие РСО.
 Например, многие Движники принимают участие в организации Школы вожатых, в организации форумов. 

 Каждую зиму бойцы отряда выезжают участниками Всероссийской патриотической акции Снежный десант РСО.
 На данный момент наши бойцы состоят в отрядах Снежного десанта «Клюква», «Пламя», «Буря», «Белое золото»
	
 И всё это - только верхушка айсберга. Гораздо больше ты узнаешь, когда станешь частью нашей семьи. А пока можешь изучить нашу страничку в социальной сети Вконтакте https://vk.com/spodvijnik1 или в мессенджере телеграм t.me/dvizhniki.`,
	keyboard: tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("В каких лагерях работают Движники?")),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Какие знамёна есть у Движников?"))),
	photo: nil}
var Camps = answer{
	text: `Движники работают во многих лагерях. Этим летом, например, мы выезжали в «Бугорок», «Салют», «Планета Английского», «SportZania», «Химик» `,
	keyboard: tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Круть!"))),
	photo: nil}
var Detouchments = answer{
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
			tgbotapi.NewKeyboardButton("ПСО"),
		)),
	photo: nil}
var RSD = answer{
	text: `В России существует организация, благодаря которой тысячи студентов могут найти себе работу на лето, компанию единомышленников, пространство для творечества и многое другое. 
 Эта организация называется Российские студенческие отряды (РСО). 
 Большая организация состоит из множества маленьких ячеек - линейных отрядов.`, //RSD = Russian students detouchments - РСО
	keyboard: tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Расскажи больше про отряды"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Расскажи историю РСО"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Какие направления есть в РСО?"),
		)),
	photo: nil}
var Svodniy = answer{
	text: `Большинство существующих отрядов прикреплены к какой-либо образовательной организации - их штабу. 
 Однако кроме них существуют отряды без штаба. Такие отряды называются сводными и именно таким отрядом являются Движники.`,
	keyboard: tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Какие ещё есть отряды?"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Расскажи больше про твой отряд"),
		),
	),
	photo: nil}
var MoreRSD = answer{
	text: `Полное название организации - Молодёжная общероссийская общественная организация Российские студенческие отряды (МООО «РСО»). Это крупнейшая молодёжная организация России, объединяющая более 150 тысяч молодых людей из 76 субъектов страны.

 РСО предоставляет студентом огромное количество возможностей для самореализации. Ежегодно проводятся конкурсы профмастерства, а лучшие отряды получат знамёна (кстати, у «Движников» тоже есть знамёна). Кроме профессиональных навыков можно испытать себя и в творчестве - на творческих конкурсах - и в спорте - на спартакиадах.`,
	keyboard: tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Какие знамёна есть у Движников?"))),
	photo: nil}

// Создание интелекта
func Create(id_user int64) *Params {
	p := &Params{id_user: id_user, sost: []int{0, 0}}
	log := logical{
		br01: make(map[string]answer, 0),
		br02: make(map[string]answer, 0),
		br03: make(map[string]answer, 0),
		br04: make(map[string]answer, 0),
		br05: make(map[string]answer, 0),
		br06: make(map[string]answer, 0),
		br07: make(map[string]answer, 0),
		br08: make(map[string]answer, 0),
		br09: make(map[string]answer, 0),
		br10: make(map[string]answer, 0),
		br11: make(map[string]answer, 0),
		br12: make(map[string]answer, 0),
		br13: make(map[string]answer, 0),
		br14: make(map[string]answer, 0),
		br15: make(map[string]answer, 0),
		br16: make(map[string]answer, 0),
		br17: make(map[string]answer, 0),
	}
	log.br01["Что такое отряды?"] = RSD
	log.br02["Что значит Сводный отряд?"] = Svodniy
	log.br03["Что такое педагогический отряд?"] = SPO
	log.br04["Какие есть отряды?"] = Detouchments
	log.br05["В каких лагерях работают Движники?"] = Camps
	log.br06["Расскажи больше про РСО"] = MoreRSD
	log.br07["Какие знамёна есть у Движников?"] = Znamea
	log.br08["ССО"] = SSO
	log.br09["СПО"] = SPO
	log.br10["СОП"] = SOP
	log.br11["ССервО"] = SServO
	log.br12["ССхО"] = SShO
	log.br13["СОП"] = SOP
	log.br14["ТОП"] = TOP
	log.br15["ПСО"] = PSO
	log.br16["Расскажи больше про твой отряд"] = Dvijniki
	log.br17["Расскажи историю РСО"] = History

	return p
}

/*
 photo := tgbotapi.NewPhoto(p.id_user, tgbotapi.FilePath("имя картинки"))
 вертуть третьим аргументом
*/
// основа диалога
func (p *Params) GetAnswer(command string, msg string) (answer, *tgbotapi.ReplyKeyboardMarkup, *tgbotapi.PhotoConfig, error) {
	//Возрождаем бота, если он застрял
	if msg == "Всё плохо, давай сначала" {

	}
	if msg == "Кто тебя создал?" {

	}

	return errmsg, nil, nil, nil
}
