package modelTables

import (
	"PPO_BMSTU/internal/models"
	"fmt"
	"os"
	"text/tabwriter"
)

func Judges(judges []models.Judge) error {
	var err error

	t := new(tabwriter.Writer)

	t.Init(os.Stdout, 2, 4, 5, ' ', 0)

	_, err = fmt.Fprintf(t, "\n %s\t%s\t%s\t%s\t%s",
		"№", "ФИО", "Роль", "Должность", "Логин")
	if err != nil {
		fmt.Println(err)
	}

	for i, judge := range judges {
		role, err := JudgeRoleToString(judge.Role)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(t, "\n %d\t%s\t%s\t%s\t%s",
			i+1, judge.FIO, role, judge.Post, judge.Login)
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
