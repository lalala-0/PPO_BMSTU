package modelTables

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/internal/registry"
	"fmt"
	"os"
	"text/tabwriter"
)

func Crews(service registry.Services, crews []models.Crew) error {
	var err error

	t := new(tabwriter.Writer)

	t.Init(os.Stdout, 2, 4, 5, ' ', 0)

	_, err = fmt.Fprintf(t, "\n %s\t%s\t%s\t%s",
		"№", "Номер на парусе", "Класс", "Участники")
	if err != nil {
		fmt.Println(err)
	}

	for i, crew := range crews {
		participants, err := service.ParticipantService.GetParticipantsDataByCrewID(crew.ID)
		if err != nil {
			return err
		}
		participantsList := ""
		for _, participant := range participants {
			participantsList += participant.FIO + " "
		}

		class, err := ClassToString(crew.Class)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(t, "\n %d\t%d\t%s\t%s",
			i+1, crew.SailNum, class, participantsList)
		if err != nil {
			return err
		}
	}

	err = t.Flush()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(t, "\n")
	if err != nil {
		return err
	}
	return nil
}
