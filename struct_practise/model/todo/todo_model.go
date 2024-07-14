package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type TodoModel struct {
	Text string `json:"todoText"`
}

func New(text string) (*TodoModel, error) {
	if text == "" {
		return &TodoModel{}, errors.New("invalid input")
	}
	return &TodoModel{
		Text: text,
	}, nil
}

func (todo *TodoModel) DisplayTodo() {
	fmt.Println("todo text is :", todo.Text)
}

func (data *TodoModel) Save(fileName string) error {
	file, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, file, 0644)
}
