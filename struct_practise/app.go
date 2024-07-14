package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/model/note"
	"example.com/note/model/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Dispay()
}

func main() {
	titleInput, contentInput := getNotesData()
	todoText := getTodoData()
	newTodo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	newNote, err := note.New(titleInput, contentInput)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(newNote)

	if err != nil {
		return
	}

	outputData(newTodo)
}

func outputData(data outputtable) error {
	data.Dispay()
	return saveDate(data)
}

func saveDate(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving todo failed")
		return err
	}

	fmt.Println("saved the todo")
	return nil
}

func getTodoData() string {
	text := getUserInput("Enter text: ")
	return text
}

func getNotesData() (string, string) {
	titleInput := getUserInput("Enter title: ")
	contentInput := getUserInput("Enter content: ")
	return titleInput, contentInput
}

func getUserInput(prmopt string) string {
	var value string
	fmt.Println(prmopt)

	// Handle long user input
	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")

	return value
}
