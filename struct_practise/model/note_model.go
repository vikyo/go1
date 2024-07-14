package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type NoteModel struct {
	Title     string    `json:"noteTitle"`
	Content   string    `json:"noteContent"`
	CreatedAt time.Time `json:"noteCreatedAt"`
}

func New(title, content string) (*NoteModel, error) {
	if title == "" || content == "" {
		return &NoteModel{}, errors.New("invalid input")
	}
	return &NoteModel{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note *NoteModel) DisplayNote() {
	fmt.Println("title is :", note.Title, "\n", "contnet is: ", note.Content)
}

func (data *NoteModel) Save(fileName string) error {
	file, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, file, 0644)
}
