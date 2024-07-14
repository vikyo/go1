package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/model"
	"example.com/note/model/todo"
)

const fileName = "notes.json"
const todoFileName = "todo.json"

func main() {
	titleInput, contentInput := getNotesData()
	todoText := getTodoData()
	newTodo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	newNote, err := model.New(titleInput, contentInput)

	if err != nil {
		fmt.Println(err)
		return
	}

	newNote.DisplayNote()
	err = newNote.Save(fileName)

	if err != nil {
		fmt.Println("Saving note failed")
		return
	}

	fmt.Println("saved the note")

	newTodo.DisplayTodo()
	err = newTodo.Save(todoFileName)

	if err != nil {
		fmt.Println("Saving todo failed")
		return
	}

	fmt.Println("saved the todo")
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
