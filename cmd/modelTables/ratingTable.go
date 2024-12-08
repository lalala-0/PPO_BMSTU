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

	_, err = fmt.Fprintf(t, "\n %s\t%s\t%s",
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

	_, err = fmt.Fprintf(t, "\n")
	if err != nil {
		return err
	}
	return nil
}

func RatingTableLines(lines []models.RatingTableLine) error {
	t := new(tabwriter.Writer)
	t.Init(os.Stdout, 2, 4, 8, ' ', 0)

	if err := printHeader(t, lines); err != nil {
		return err
	}

	for i, line := range lines {
		if err := printLine(t, i+1, line); err != nil {
			return err
		}
	}

	return t.Flush()
}

func printHeader(t *tabwriter.Writer, lines []models.RatingTableLine) error {
	// Проверка на наличие элементов в lines
	if len(lines) == 0 {
		return fmt.Errorf("no rating table lines provided")
	}

	_, err := fmt.Fprintf(t, "\n %s\t%s\t%s", "№", "№ паруса", "Рулевой")
	if err != nil {
		return err
	}

	// Цикл для вывода гонок, если строки не пустые
	for i := 0; i < len(lines[0].ResInRace); i++ {
		_, err = fmt.Fprintf(t, "\t %s %d", "гон.", i+1)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprintf(t, "\t%s\t%s\n", "Очки", "Место")
	return err
}

func printLine(t *tabwriter.Writer, i int, line models.RatingTableLine) error {
	// Печать информации о гонке
	_, err := fmt.Fprintf(t, " %d\t%d\t%s", i, line.SailNum, getParticipantName(line))
	if err != nil {
		return err
	}

	// Печать результатов в гонке
	for _, res := range line.ResInRace {
		if err := printResult(t, res); err != nil {
			return err
		}
	}

	// Печать итоговых очков и места
	_, err = fmt.Fprintf(t, "\t%d\t%d\n", line.PointsSum, line.Rank)
	return err
}

func getParticipantName(line models.RatingTableLine) string {
	if len(line.ParticipantNames) > 0 {
		return line.ParticipantNames[0]
	}
	return ""
}

func printResult(t *tabwriter.Writer, result int) error {
	_, err := fmt.Fprintf(t, "\t %d", result)
	return err
}
