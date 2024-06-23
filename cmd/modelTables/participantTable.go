package modelTables

import (
	"PPO_BMSTU/internal/models"
	"fmt"
	"os"
	"text/tabwriter"
)

func Participants(participants []models.Participant) error {
	var err error

	t := new(tabwriter.Writer)

	t.Init(os.Stdout, 2, 4, 4, ' ', 0)

	_, err = fmt.Fprintf(t, "\n %s\t%s\t%s\t%s\t%s\t%s",
		"№", "ФИО", "Разряд", "Пол", "Дата рождения", "Тренер")
	if err != nil {
		fmt.Println(err)
	}

	for i, participant := range participants {
		category, err := ParticipantCategoryToString(participant.Category)
		if err != nil {
			return err
		}
		gender, err := GenderToString(participant.Gender)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(t, "\n %d\t%s\t%s\t%s\t%s\t%s",
			i+1, participant.FIO, category, gender, participant.Birthday, participant.Coach)
		if err != nil {
			return err
		}
	}

	err = t.Flush()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(t, "\n")
	return nil
}
