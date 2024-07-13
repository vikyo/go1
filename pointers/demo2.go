package main

import "fmt"

func TestAgeUsingPointers(age int) {
	agePointer := &age
	fmt.Println("Address: ", agePointer, "Age: ", *agePointer)
	getAdultYears(agePointer)
	fmt.Println("Adult years: ", age)
}

func getAdultYears(val *int) {
	*val -= 18
}

func TestAge(age int) {
	fmt.Println("Address: ", &age, "Age: ", age)
	var years = getAdultYears2(age)
	fmt.Println("Adult years: ", years)
}

func getAdultYears2(val int) int {
	return val - 18
}
