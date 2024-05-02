package modelTables

import (
	"PPO_BMSTU/internal/models"
	"fmt"
)

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
