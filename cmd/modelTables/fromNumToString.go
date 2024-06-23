package modelTables

import (
	"PPO_BMSTU/internal/models"
	"fmt"
)

func JudgeRoleToString(role int) (string, error) {
	if role == models.MainJudge {
		return "Главный судья", nil
	} else if role == models.NotMainJudge {
		return "Судья", nil
	}
	return "", fmt.Errorf("Некорректная роль судьи", "role", role)
}

func ProtestStatusToString(status int) (string, error) {
	if status == models.PendingReview {
		return "Ожидает рассмотрения", nil
	} else if status == models.Reviewed {
		return "Рассмотрен", nil
	}
	return "", fmt.Errorf("Некорректный статус протеста", "status", status)
}

func GenderToString(status int) (string, error) {
	if status == models.Male {
		return "Male", nil
	} else if status == models.Female {
		return "Жен.", nil
	}
	return "", fmt.Errorf("Некорректный пол", "status", status)
}

func ParticipantCategoryToString(category int) (string, error) {
	if category == models.MasterInternational {
		return "Мастер спорта России медунородного класса", nil
	} else if category == models.MasterRussia {
		return "Мастер спорта России", nil
	} else if category == models.Candidate {
		return "Кандидат в мастера спорта ", nil
	} else if category == models.Sport1category {
		return "1-ый спортивный разряд", nil
	} else if category == models.Sport2category {
		return "2-ой спортивный разряд", nil
	} else if category == models.Sport3category {
		return "3-ий спортивный разряд", nil
	} else if category == models.Junior1category {
		return "1-ый юношеский разряд", nil
	} else if category == models.Junior2category {
		return "2-ой юношеский разряд", nil
	}
	return "", fmt.Errorf("Некорректный спортивный разряд", "category", category)
}

func SpecCircumstanceToString(specCircumstance int) (string, error) {
	if specCircumstance == models.DNS {
		return "DNS", nil
	} else if specCircumstance == models.DNF {
		return "DNF", nil
	} else if specCircumstance == models.DNC {
		return "DNC", nil
	} else if specCircumstance == models.OCS {
		return "OCS", nil
	} else if specCircumstance == models.ZFP {
		return "ZFP", nil
	} else if specCircumstance == models.UFD {
		return "UFD", nil
	} else if specCircumstance == models.BFD {
		return "BFP", nil
	} else if specCircumstance == models.SCP {
		return "SCP", nil
	} else if specCircumstance == models.RET {
		return "RET", nil
	} else if specCircumstance == models.DSQ {
		return "DSQ", nil
	} else if specCircumstance == models.DNE {
		return "DNE", nil
	} else if specCircumstance == models.RDG {
		return "RDG", nil
	} else if specCircumstance == models.DPI {
		return "DPI", nil
	} else if specCircumstance == 0 {
		return "-", nil
	}
	return "", fmt.Errorf("Некорректный номер, обозначающий специальные обстоятельства", "specCircumstance", specCircumstance)
}

func ClassToString(class int) (string, error) {
	if class == models.Laser {
		return "Laser", nil
	} else if class == models.LaserRadial {
		return "LaserRadial", nil
	} else if class == models.Optimist {
		return "Optimist", nil
	} else if class == models.Zoom8 {
		return "Zoom8", nil
	} else if class == models.Finn {
		return "Finn", nil
	} else if class == models.SB20 {
		return "SB20", nil
	} else if class == models.J70 {
		return "J70", nil
	} else if class == models.Nacra17 {
		return "Nacra17", nil
	} else if class == models.C49er {
		return "C49er", nil
	} else if class == models.RS_X {
		return "RS_X", nil
	} else if class == models.Cadet {
		return "Cadet", nil
	}
	return "", fmt.Errorf("Некорректный класс", "class", class)
}

func ProtestParticipantRoleToString(role int) (string, error) {
	if role == models.Protestor {
		return "Протестующий", nil
	} else if role == models.Protestee {
		return "Опротестованный", nil
	} else if role == models.Witness {
		return "Свидетель", nil
	}
	return "", fmt.Errorf("Некорректная роль участника протеста", "role", role)
}
