package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const fileName = "todo.json"

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

func (todo *TodoModel) Dispay() {
	fmt.Println("todo text is :", todo.Text)
}

func (todo *TodoModel) PrintSuccessMessage() {
	fmt.Println("todo is saved in ", fileName)
}

func (todo *TodoModel) PrintFailedMessage() {
	fmt.Println("todo save is failed")
}

func (data *TodoModel) Save() error {
	file, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, file, 0644)
}
