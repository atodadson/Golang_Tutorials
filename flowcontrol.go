package main

import "fmt"

var class_score = 68

func main() {
	// Declare and assign a boolean
	printMessage := true

	// Take note of the structure of the structure of the if block

	if printMessage {
		fmt.Println("Hey, print message is true")
	}

	account_credit := 200
	account_balance := -300
	// If block with a condition to be evaluated
	if account_balance+account_credit > 0 {
		fmt.Println("I have money to buy stuffs.")
	} else {
		fmt.Println("Go get yourself a good job.")
	}

	// Using else if. Taking you to places if and else can't get you to
	class_score := -40
	if class_score < 0 {
		fmt.Println("Man, the test score should not be negative")
	} else if class_score >= 70 {
		fmt.Println("Yoo, that is Grade A")
	} else if class_score >= 60 {
		fmt.Println("Not bad, Grade B")
	} else if class_score >= 50 {
		fmt.Println("Wellll, it's an average performance. Grade C")
	} else if class_score >= 40 {
		fmt.Println("You passed, but nearly failed. Grade D")
	} else {
		fmt.Println("Yoo, very bad performance. Grade F")
	}

	// Testing multiple conditions
	has_food := true
	feels_hungry := true

	// && represents AND || represents OR
	if has_food && feels_hungry {
		fmt.Println("You have some food. Go and eat")
	}

	has_jollof := true
	has_spaghetti := true
	if has_jollof || has_spaghetti {
		fmt.Println("You have some food. Go and eat")
	}

	// Negating Expressions
	if !has_food {
		fmt.Println("Are you sure you don't have food")
	} else {
		fmt.Println("Go and eat")
	}

}
