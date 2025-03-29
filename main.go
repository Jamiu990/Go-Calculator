package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("**********NUMBER GUESSING GAME**********")

	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(50) + 1

	attempts := 0
	maxAttempts := 10

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Printf("I'm thinking of a number between 1 and 50. What is the number? You have %d attempts.\n", maxAttempts)

	for attempts < maxAttempts {
		attempts++
		fmt.Printf("Attempt %d: Enter your guess: ", attempts)

		var guess int
		_, err := fmt.Scan(&guess)

		if err != nil {
			fmt.Println("Invalid input. Please enter a whole number.")

			var dummy string
			fmt.Scanln(&dummy)
			continue
		}

		if guess > 50 {
			fmt.Println("Please enter a number between 1 and 50.")
			continue
		}else if guess < secretNumber {
			fmt.Println("Too low!")
		}else if guess > secretNumber {
			fmt.Println("Too high!")
		}else {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts!\n", attempts)
			break
		}
		
		if attempts == maxAttempts {
			fmt.Printf("Opps! You ran out of attempts! The secret number is %d.\n", secretNumber)
		}
	}

	fmt.Println("**********GAME OVER**********")
}
