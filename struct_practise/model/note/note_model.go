package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

const fileName = "notes.json"

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

func (note *NoteModel) Dispay() {
	fmt.Println("title is :", note.Title)
	fmt.Println("contnet is: ", note.Content)
}

func (note *NoteModel) PrintSuccessMessage() {
	fmt.Println("note is saved in ", fileName)
}

func (note *NoteModel) PrintFailedMessage() {
	fmt.Println("note save is failed")
}

func (data *NoteModel) Save() error {
	file, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, file, 0644)
}
