package main

import "fmt"

func main() {
    // Defining structs
    type Account struct {
        firstname string
        lastname string
        balance float32
        id int32
    }

    // Creating an instance of a struct - Defining individual attributes
    var myAccount Account
    myAccount.firstname = "Bernard"
    myAccount.lastname = "Darko"
    myAccount.balance = 35.45
    myAccount.id = 20

    // Creating an instance of a struct - Defining at once
    hisAccount := Account{"Kofi", "Manu", 1234.43, 12}

    fmt.Printf("Account Details: %v \n", myAccount)
    fmt.Printf("Account Details: %v", hisAccount)

}