package stringConst

const SailNumRequest = "Введите номер паруса"
const RuleNumRequest = "Введите номер нарушенного правила (10-13, 15-23, 31, 42)"
const ProtestStatusRequest = "Введите статус протеста (1 - ожидает рассмотрения, 2 - рассмотрен)"
const CommentRequest = "Введите комментарий"
const ProtesteePointsRequest = "Введите количество баллов опротестованной яхты"

const BlowCntRequest = "Введите количество выбрасываемых результатов"
const NameRequest = "Введите ФИО"
const RatingNameRequest = "Введите название регаты"
const RaceNumberRequest = "Введите номер гонки"
const LoginRequest = "Введите логин"

// #nosec G101: Это не жестко закодированный пароль, а просто текст подсказки для ввода
const PasswordRequest = "Введите пароль (Более 8 символов; должен содержать как символы, так и цифры)"

const DateRequest = "Введите дату и время (yyyy-mm-dd hh:mm)"
const JudgeRoleRequest = "Введите роль судьи (1 - главный судья, 2 - судья)"
const JudgePostRequest = "Введите должность судьи"

const WitnessesSailNumsRequest = "Введите номера парусов яхт-свидетелей"
const FinisherListRequest = "Введите номера яхт согласно очерёдности их финиша"
const FinishersWithSpecCircumstanceListRequest = "Для регистрации обстоятельств использовать" +
	"DNS = 1  - не стартовала (не подпадает под DNC и OCS),\n\t" +
	"DNF = 2  - не финишировала,\n\t" +
	"DNC = 3  - не стартовала; не прибыла в район старта,\n\t" +
	"OCS = 4  -  не стартовала; находилась на стороне дистанции от стартовой линии в момент сигнала \"Старт\" для нее и не стартовала или нарушила правило 30.1,\n\t" +
	"ZFP = 5  - 20% наказание по правилу 30.2,\n\t" +
	"UFD = 6  - дисквалификация по правилу 30.3,\n\t" +
	"BFD = 7  - дисквалификация по правилу 30.4,\n\t" +
	"SCP = 8  - применено \"Наказание штрафными очками\",\n\t" +
	"RET = 9  - вышла из гонки,\n\t" +
	"DSQ = 10 - дисквалификация,\n\t" +
	"DNE = 11 - не исключаемая дисквалификация,\n\t" +
	"RDG = 12 - исправлен результат,\n\t" +
	"DPI = 13 - наказание по усмотрению протестового комитета.\n" +
	"Введите пары номер яхты и номер обстоятельства"

const ClassRequest = "Laser       = 1\n\t" +
	"LaserRadial = 2\n\t" +
	"Optimist    = 3\n\t" +
	"Zoom8       = 4\n\t" +
	"Finn        = 5\n\t" +
	"SB20        = 6\n\t" +
	"J70         = 7\n\t" +
	"Nacra17     = 8\n\t" +
	"C49er       = 9\n\t" +
	"RS_X        = 10\n\t" +
	"Cadet       = 11\n\t" +
	"Введите номер соответствующего класса "

const BirthDateRequest = "Введите дату рождения"
const CoachNameRequest = "Введите ФИО тренера"
const GenderRequest = "Введите пол (0 - мужской, 1 - женский)"

const ParticipantCategoryRequest = "MasterInternational = 1\n\t" + // Master of Sports of Russia of international class
	"MasterRussia        = 2\n\t" + // Master of Sports of Russia
	"Candidate           = 3\n\t" + // Candidate for Master of Sports
	"Sport1category      = 4\n\t" + // 1 sports category
	"Sport2category      = 5\n\t" + // 2 sports category
	"Sport3category      = 6\n\t" + // 3 sports category
	"Junior1category     = 7\n\t" + // 1 junior category
	"Junior2category     = 8\n\t" + // 2 junior category
	"Введите номер соответствующего разряда "
