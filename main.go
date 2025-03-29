package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Printf("Welcome to your To-Do App! Please select an option by typing the corresponding number and pressing <Enter> to continue: \n")
		fmt.Printf("1 - View Tasks for the Day\n")
		fmt.Printf("2 - View Tasks for the Week\n")
		fmt.Printf("3 - Set New Task\n")
		fmt.Printf("4 - Help\n")
		fmt.Printf("5 - Quit\n")

		fmt.Print("> ")
		reader.Scan()
		text := reader.Text()

		switch text {
		case "1":
			fmt.Println("Daily Tasks")
			// daily tasks
		case "2":
			fmt.Println("Weekly Tasks")
			// weekly tasks
		case "3":
			fmt.Println("Set New Task")
			// create task
		case "4":
			fmt.Println("Help")
			// help menu
		case "5":
			fmt.Println("Goodbye")
			os.Exit(0)
		default:
			fmt.Printf("Invalid entry, please try again\n\n\n--------------------------------------------------------------------------------\n\n\n")
		}
	}
}
