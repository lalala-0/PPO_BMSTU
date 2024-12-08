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
	return "", fmt.Errorf("Некорректная роль судьи: role = %v", role)
}

func ProtestStatusToString(status int) (string, error) {
	if status == models.PendingReview {
		return "Ожидает рассмотрения", nil
	} else if status == models.Reviewed {
		return "Рассмотрен", nil
	}
	return "", fmt.Errorf("Некорректный статус протеста: status=%v", status)
}

func GenderToString(status int) (string, error) {
	if status == models.Male {
		return "Male", nil
	} else if status == models.Female {
		return "Female", nil
	}
	return "", fmt.Errorf("Некорректный пол: status=%v", status)
}

func ParticipantCategoryToString(category int) (string, error) {
	categories := map[int]string{
		models.MasterInternational: "Мастер спорта России международного класса",
		models.MasterRussia:        "Мастер спорта России",
		models.Candidate:           "Кандидат в мастера спорта",
		models.Sport1category:      "1-ый спортивный разряд",
		models.Sport2category:      "2-ой спортивный разряд",
		models.Sport3category:      "3-ий спортивный разряд",
		models.Junior1category:     "1-ый юношеский разряд",
		models.Junior2category:     "2-ой юношеский разряд",
	}

	if result, exists := categories[category]; exists {
		return result, nil
	}

	return "", fmt.Errorf("Некорректный спортивный разряд: category = %v", category)
}

func SpecCircumstanceToString(specCircumstance int) (string, error) {
	specCircumstances := map[int]string{
		models.DNS: "DNS",
		models.DNF: "DNF",
		models.DNC: "DNC",
		models.OCS: "OCS",
		models.ZFP: "ZFP",
		models.UFD: "UFD",
		models.BFD: "BFP", // Исправлено на "BFP", так как в предыдущем коде был "BFP"
		models.SCP: "SCP",
		models.RET: "RET",
		models.DSQ: "DSQ",
		models.DNE: "DNE",
		models.RDG: "RDG",
		models.DPI: "DPI",
		0:          "-",
	}

	if result, exists := specCircumstances[specCircumstance]; exists {
		return result, nil
	}

	return "", fmt.Errorf("Некорректный номер, обозначающий специальные обстоятельства: specCircumstance = %v", specCircumstance)
}

func ClassToString(class int) (string, error) {
	classes := map[int]string{
		models.Laser:       "Laser",
		models.LaserRadial: "LaserRadial",
		models.Optimist:    "Optimist",
		models.Zoom8:       "Zoom8",
		models.Finn:        "Finn",
		models.SB20:        "SB20",
		models.J70:         "J70",
		models.Nacra17:     "Nacra17",
		models.C49er:       "C49er",
		models.RS_X:        "RS_X",
		models.Cadet:       "Cadet",
	}

	if result, exists := classes[class]; exists {
		return result, nil
	}

	return "", fmt.Errorf("Некорректный класс: class = %v", class)
}

func ProtestParticipantRoleToString(role int) (string, error) {
	if role == models.Protestor {
		return "Протестующий", nil
	} else if role == models.Protestee {
		return "Опротестованный", nil
	} else if role == models.Witness {
		return "Свидетель", nil
	}
	return "", fmt.Errorf("Некорректная роль участника протеста: role = %v", role)
}
