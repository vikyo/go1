package main

import "fmt"

func main() {
	test := 32
	age := 19
	fmt.Println("1: ", &test, test)
	passByValue(test)
	fmt.Println("3: ", &test, test)
	passByReference(&test)
	fmt.Println("5: ", &test, test)

	TestAgeUsingPointers(age)
	TestAge(age)
}

func passByValue(value int) {
	value = 100
	fmt.Println("2: ", &value, value)
}

func passByReference(value *int) {
	*value = 200
	fmt.Println("4: ", value, *value)
}
