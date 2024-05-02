package modelTables

import (
	"PPO_BMSTU/cmd/cmdUtils"
	"PPO_BMSTU/internal/models"
	"fmt"
	"os"
	"text/tabwriter"
)

func Ratings(ratings []models.Rating) error {
	var err error

	maxNameLen := 0
	for _, rating := range ratings {
		if len(rating.Name) > maxNameLen {
			maxNameLen = len(rating.Name)
		}
	}

	t := new(tabwriter.Writer)

	t.Init(os.Stdout, 2, 4, 3, ' ', 0)

	_, err = fmt.Fprintf(t, "\n %s\t%s\t%s\t%s\t%s",
		"№", "Название регаты", "Класс яхт")
	if err != nil {
		fmt.Println(err)
	}

	for i, rating := range ratings {
		class, err := ClassToString(rating.Class)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(t, "\n %d\t%s\t%s",
			i+1, cmdUtils.TruncateString(rating.Name, 20), class)
		if err != nil {
			return err
		}
	}

	err = t.Flush()
	if err != nil {
		return err
	}

	return nil
}
