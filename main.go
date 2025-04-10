package main

import "fmt"

func main() {
	var name string
	var age int

	// Ask for user input
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Enter your age:")
	fmt.Scanln(&age)

	// Use switch with conditions for age checks
	switch {
	case age < 18:
		fmt.Printf("%s, you are too young\n", name)
	case age >= 18:
		fmt.Printf("Welcome, %s! You are old enough!\n", name)
	}
}

