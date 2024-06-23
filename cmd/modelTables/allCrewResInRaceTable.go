package modelTables

import (
	"PPO_BMSTU/internal/models"
	"fmt"
	"github.com/google/uuid"
	"os"
	"text/tabwriter"
)

func getSailNumByCrewID(id uuid.UUID, crews []models.Crew) int {
	for _, crew := range crews {
		if crew.ID == id {
			return crew.SailNum
		}
	}
	return -1
}

func AllCrewResInRace(allCrewResInRace []models.CrewResInRace, crews []models.Crew) error {
	var err error

	t := new(tabwriter.Writer)

	t.Init(os.Stdout, 2, 4, 4, ' ', 0)

	_, err = fmt.Fprintf(t, "\n %s\t%s\t%s\t%s\t%s",
		"№", "№ паруса", "Очки", "Спец. обстоятельства")
	if err != nil {
		fmt.Println(err)
	}

	for i, res := range allCrewResInRace {
		sailNum := getSailNumByCrewID(res.CrewID, crews)
		if sailNum == -1 {
			return fmt.Errorf("Такой команды в данной гонке не зарегистрировано")
		}

		specCircumstance, err := SpecCircumstanceToString(res.SpecCircumstance)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(t, "\n %d\t%d\t%d\t%s",
			i+1, sailNum, res.Points, specCircumstance)
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
