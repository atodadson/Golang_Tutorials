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
    fmt.Printf("Account Details: %v \n", hisAccount)

    type Address struct {
        city string
        street string
        postal string
    }

    // Embedding a struct in another struct
    type Person struct {
        name string
        address Address // You can use the type as the name by ommitting the name
    }

    person1 := Person{"Kwame BIOS", Address{"New York", "Street 34", "SWE23"}}
    fmt.Printf("person1: %v \n", person1)

    // Adding implementation to structs
    func strr(a Address) string {
        return fmt.Sprintf("City: %s, Street: %s, Postal: %s", a.city, a.street, a.postal)
    }

    fmt.Println(strr(person1.address))

}