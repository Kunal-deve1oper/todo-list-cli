package todo

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

type todoItems struct {
	Task        string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []todoItems

func (t *Todos) Add(str string) {
	todo := todoItems{
		Task:        str,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*t = append(*t, todo)
}

func (t *Todos) CompletedTask(index int) error {
	if index <= 0 || index > len(*t) {
		return errors.New("invalid index")
	}
	(*t)[index-1].Completed = true
	(*t)[index-1].CompletedAt = time.Now()
	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}
	if len(ls) == 0 {
		return errors.New("empty to-do-list")
	}
	*t = append(ls[:index-1], ls[index:]...)
	return nil
}

func (t *Todos) Save(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return errors.New("invalid data")
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return errors.New("cannot store the data to file")
	}
	return nil
}

func (t *Todos) List(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return errors.New("cannot read the file")
	}
	if len(data) == 0 {
		return errors.New("file is empty")
	}
	err = json.Unmarshal(data, t)
	if err != nil {
		return errors.New("cannot parse the json data")
	}
	return nil
}

func (t *Todos) Display() {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Id", "Task", "Completed", "Created_At", "Completed_At"})

	table.SetHeaderColor(tablewriter.Colors{tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.FgWhiteColor})
	table.SetColumnColor(tablewriter.Colors{tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.FgWhiteColor})

	for i, data := range *t {
		var completedColor tablewriter.Colors
		if data.Completed {
			completedColor = tablewriter.Colors{tablewriter.FgGreenColor}
			table.Rich([]string{strconv.Itoa(i + 1), data.Task, strconv.FormatBool(data.Completed), data.CreatedAt.Format("2006-01-02"), data.CompletedAt.Format("2006-01-02")}, []tablewriter.Colors{
				{},
				completedColor,
				completedColor,
				{},
				{},
			})
		} else {
			table.Append([]string{strconv.Itoa(i + 1), data.Task, strconv.FormatBool(data.Completed), data.CreatedAt.Format("2006-01-01"), data.CompletedAt.String()})
		}
	}
	table.Render()
}
