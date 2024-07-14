package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/model"
)

const fileName = "notes.json"

func main() {
	titleInput, contentInput := getNotesDate()
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

	fmt.Printf("saved the note")
}

func getNotesDate() (string, string) {
	titleInput := getUserInput("Enter title: ")
	contentInput := getUserInput("Enter content: ")
	return titleInput, contentInput
}

func getUserInput(prmopt string) string {
	var value string
	fmt.Printf(prmopt)

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
