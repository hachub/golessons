package main

import (
	"fmt"
	"time"
)
import "math/rand"

func main() {
	const attempts = 10
	const max = 100

	rand.Seed(time.Now().UnixNano())
	var number = rand.Intn(max)

	var input int
	win := false

	fmt.Println("Hello guy! Let's play a game!")

	for i := 0; i < attempts; i++ {
		fmt.Println("ATTEMPT ", i)
		fmt.Printf("Enter a number between 0 and %d: ", max)
		_, err := fmt.Scanln(&input)
		if err != nil || input < 0 || input > 100 {
			fmt.Println("Error input!")
		} else if input < number {
			fmt.Println("Your number is lower than my number!")
		} else if input > number {
			fmt.Println("Your number is bigger than my number!")
		} else {
			fmt.Println("Congratulations!")
			win = true
			break
		}
	}

	if !win {
		fmt.Println("Sorry! You lose (")
	}
}
