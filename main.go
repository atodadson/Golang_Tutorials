package main

import "fmt"

func main() {
	// Declaration of variables
	var first_name String
	var last_name String
	var age := 34

	// Declaration and assignment of variables
	lady_fname := "Kwame"
	lady_lname := "Abena"
	
	// The Println function does not take formatters
	fmt.Println(lady_fname, lady_lname)

	// We use the Printf function which works with formatters
	// %s is for Strings and %d is for numbers
	fmt.Printf("%s %s is %d years old. Isn't it shocking")
}
