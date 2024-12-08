package modelTables

import (
	"PPO_BMSTU/internal/models"
	"fmt"
	"os"
	"text/tabwriter"
)

func Races(races []models.Race) error {
	var err error

	t := new(tabwriter.Writer)

	t.Init(os.Stdout, 2, 4, 4, ' ', 0)

	_, err = fmt.Fprintf(t, "\n %s\t%s\t%s\t%s",
		"№", "Дата гонки", "Номер гонки", "Класс яхт")
	if err != nil {
		fmt.Println(err)
	}

	for i, race := range races {
		class, err := ClassToString(race.Class)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(t, "\n %d\t%s\t%d\t%s",
			i+1, race.Date, race.Number, class)
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
