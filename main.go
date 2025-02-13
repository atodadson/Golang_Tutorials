package main

import "fmt"

func main() {
	// Declaration of variables
	// First Declare variable name and type
	var first_name string
	var last_name string
	first_name = "Kofi"
	last_name = "Adjo"
	

	// Declaration and assignment of variables
	// variable_name := value (Useful in functions)
	age := 34
	lady_fname := "Kwame"
	lady_lname := "Abena"
	
	// The Println function does not take formatters
	fmt.Println(lady_fname, lady_lname)

	// We use the Printf function which works with formatters
	// %s is for Strings and %d is for numbers
	fmt.Printf("%s %s is %d years old. Isn't it shocking", first_name, last_name, age)
}
