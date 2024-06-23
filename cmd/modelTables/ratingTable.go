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
	return nil
}

func RatingTableLines(lines []models.RatingTableLine) error {
	var err error

	t := new(tabwriter.Writer)

	t.Init(os.Stdout, 2, 4, 8, ' ', 0)

	_, err = fmt.Fprintf(t, "\n %s\t%s\t%s",
		"№", "№ паруса", "Рулевой")
	if err != nil {
		fmt.Println(err)
	}
	for i := range len(lines[1].ResInRace) {
		_, err = fmt.Fprintf(t, "\t %s %d", "гон.", i+1)
		if err != nil {
			fmt.Println(err)
		}
	}
	_, err = fmt.Fprintf(t, "\t%s\t%s\n", "Очки", "Место")
	if err != nil {
		fmt.Println(err)
	}

	for i, line := range lines {
		if len(line.ParticipantNames) > 0 {
			_, err = fmt.Fprintf(t, " %d\t%d\t%s",
				i+1, line.SailNum, line.ParticipantNames[0])
		} else {
			_, err = fmt.Fprintf(t, " %d\t%d\t%s",
				i+1, line.SailNum, "")
		}
		if err != nil {
			fmt.Println(err)
		}
		//for i, _ := range line.ParticipantNames {
		//	category, err := ParticipantCategoryToString(line.ParticipantCategories[i])
		//	_, err = fmt.Fprintf(t, "\t%s\t%s\t%s",
		//		line.ParticipantNames[i], category, line.CoachNames[i])
		//	if err != nil {
		//		fmt.Println(err)
		//	}
		//}
		for i, _ := range line.ResInRace {
			_, err = fmt.Fprintf(t, "\t %d", line.ResInRace[i])
			if err != nil {
				fmt.Println(err)
			}
		}
		_, err = fmt.Fprintf(t, "\t%d\t%d\n", line.PointsSum, line.Rank)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = t.Flush()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(t, "\n")
	return nil
}
