package main

import (
	"fmt"
	"math"
	"strconv"
)

func main(){
	var op int
	var num1, num2, base, exponent float64

	for {
		fmt.Println("..........Calculator Menu.........: \n 1. Add \n 2. Subtract \n 3. Multiply \n 4. Divide \n 5. Power \n 6. Exit")
		fmt.Println("Choose an option: ")

		_, err := fmt.Scan(&op)
		if err != nil{
			fmt.Println("Invalid input for operation. Please enter a number between 1 and 6.")
			var dummy string
			fmt.Scanln(&dummy)
			continue
		}

		if op >= 1 && op <= 4{
			fmt.Printf("Enter the first number: ")
			var input1 string
			_, err := fmt.Scan(&input1)
			if err != nil {
				fmt.Println("Invlaid input. Please enter a number for the first operand.")
				var dummy string
				fmt.Scanln(&dummy)
				continue
			}
			num1, err = strconv.ParseFloat(input1, 64)
			if err != nil{
				fmt.Println("Invalid input. Please enter a valid number for the first operand.")
				var dummy string
				fmt.Scanln(&dummy)
				continue
			}
			
			fmt.Printf("Enter the second number: ")
			var input2 string
			_, err = fmt.Scan(&input2)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number for the second operand.")
				continue
			}
			num2, err = strconv.ParseFloat(input2, 64)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number for the second operand.")
				continue
			}
		}else if op == 5{
			fmt.Println("Enter the base number: ")
			var inputBase string
			_, err := fmt.Scan(&inputBase)
			if err != nil {
				fmt.Println("Invalid input. Please enter a number for the base.")
				var dummy string
				fmt.Scanln(&dummy)
				continue
			}
			base,  err = strconv.ParseFloat(inputBase, 64)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number for the base.")
				continue
			}

			fmt.Println("Enter the exponent number: ")
			var inputExponent string
			_, err = fmt.Scan(&inputExponent)
			if err != nil {
				fmt.Println("Invalid input. Pleae enter a number for the exponent.")
				var dummy string
				fmt.Scanln(&dummy)
				continue
			}
			exponent, err = strconv.ParseFloat(inputExponent, 64)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number for the exponent.")
				continue
			}
		}else if op == 6{
			fmt.Println("..........Exiting Calculator..........")
			break
		}else{
			fmt.Println("Choose an appropriate option from the menu")
			continue
		}

		switch int(op){
		case 1: 
			fmt.Printf("Result: %v\n", num1 + num2)
		case 2:
			fmt.Printf("Result: %v\n", num1 - num2)
		case 3:
			fmt.Printf("Result: %v\n", num1 * num2)
		case 4:
			if num2 == 0{
				fmt.Printf("Error!\n")
			}else{
				fmt.Printf("Result: %.3v\n", num1 / num2)
			}
		case 5:
			fmt.Printf("Result: %.2f\n", math.Pow(base, exponent))
		case 6:
			fmt.Println("..........Exiting Calculator..........")
		default:
			fmt.Println("Invalid Choice! Please try again!")

		}

	}
}
