package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/Kunal-deve1oper/todo-list-cli/pkg/handleerror"
	"github.com/Kunal-deve1oper/todo-list-cli/pkg/todo"
)

const filename string = "todo-list.json"

func main() {
	text := flag.String("a", "", "name of the task you want to add")
	comp := flag.Int("c", 0, "index of the task you want to set complete")
	del := flag.Int("d", 0, "index of the task you want to delete")
	view := flag.Bool("v", true, "view the tasks")

	flag.Parse()

	todos := &todo.Todos{}

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(filename)
		if err != nil {
			handleerror.HandleError(err)
		}
		_, err = file.WriteString("[]")
		if err != nil {
			handleerror.HandleError(err)
		}

		err = file.Close()
		if err != nil {
			handleerror.HandleError(err)
		}
	}

	err := todos.List(filename)
	if err != nil {
		handleerror.HandleError(err)
	}

	switch {
	case *text != "":
		todos.Add(*text)
		err := todos.Save(filename)
		if err != nil {
			handleerror.HandleError(err)
		}
	case *comp > 0:
		err := todos.CompletedTask(*comp)
		if err != nil {
			handleerror.HandleError(err)
		}
		err = todos.Save(filename)
		if err != nil {
			handleerror.HandleError(err)
		}
	case *del > 0:
		err := todos.Delete(*del)
		if err != nil {
			handleerror.HandleError(err)
		}
		err = todos.Save(filename)
		if err != nil {
			handleerror.HandleError(err)
		}
	case *view:
		todos.Display()
	default:
		fmt.Println("wrong flag")
		os.Exit(1)
	}
}
