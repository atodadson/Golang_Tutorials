package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	Id   int    `json: "id"`
	Name string `json: "name"`
}

type Products struct {
	Products []Product `json: products`
}

type Person struct {
	Id   int    `json: "id"`
	Name string `json: "name"`
	Age  int    `json: "age"`
}

type Persons struct {
	Persons []Person `json: persons`
}

type Response struct {
	Products []Product `json: "products"`
}

func main() {
	str := `{"name": "my product", "id": 1}`
	product := Product{}
	json.Unmarshal([]byte(str), &product)
	fmt.Println(product)

	file, _ := os.ReadFile("products.json")

	data := Products{}

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.Products); i++ {
		fmt.Println("Product Id: ", data.Products[i].Id)
		fmt.Println("Name: ", data.Products[i].Name)
	}

	Persons_list := Persons{
		Persons: []Person{
			{Id: 1, Name: "Kwame", Age: 34},
			{Id: 2, Name: "Kojo", Age: 23},
		},
	}

	jsonData, _ := json.MarshalIndent(Persons_list, "", " ")
	fmt.Println(string(jsonData))

	os.WriteFile("persons.json", jsonData, 0644)
}
