package main

import (
	"fmt"
	"os"
)

var users = map[string]int{}

func main() {
	functionality()
}

func add() {
	var count int
	var name string
	var age int
	fmt.Println("Enter the number of people:")
	fmt.Fscan(os.Stdin, &count)
	for i := 0; i < count; i++ {
		fmt.Println("Enter name: ")
		fmt.Fscan(os.Stdin, &name)
		fmt.Println("Enter age: ")
		fmt.Fscan(os.Stdin, &age)
		users[name] = age
	}
	fmt.Println(users)
	functionality()
}

func del() {
	fmt.Println(users)
	var delName string
	fmt.Println("Enter the name of the person to delete")
	fmt.Fscan(os.Stdin, &delName)
	delete(users, delName)
	fmt.Println(users)
	functionality()
}

func functionality() {
	var input uint8
	fmt.Println("Select an action:")
	fmt.Println("1: View list of people")
	fmt.Println("2: Add people")
	fmt.Println("3: Cut people in the list")
	fmt.Fscan(os.Stdin, &input)

	switch input {
	case 1:
		fmt.Println(users)
		functionality()
	case 2:
		add()
	case 3:
		del()
	default:
		fmt.Println("Invalid number!")
	}
}
