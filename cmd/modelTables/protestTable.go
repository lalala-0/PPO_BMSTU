package modelTables

import (
	"PPO_BMSTU/internal/models"
	"fmt"
	"os"
	"text/tabwriter"
)

func Protests(protests []models.Protest) error {
	var err error

	t := new(tabwriter.Writer)

	t.Init(os.Stdout, 2, 4, 4, ' ', 0)

	_, err = fmt.Fprintf(t, "\n %s\t%s\t%s\t%s\t%s",
		"№", "Номер правила", "Дата рассмотрения", "Статус", "Комментарий")
	if err != nil {
		fmt.Println(err)
	}

	for i, protest := range protests {
		status, err := ProtestStatusToString(protest.Status)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(t, "\n %d\t%d\t%s\t%s\t%s",
			i+1, protest.RuleNum, protest.ReviewDate, status, protest.Comment)
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
