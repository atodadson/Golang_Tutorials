package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Product struct {
	Id   int    `json: "id"`
	Name string `json: "name"`
}

type Products struct {
	Products []Product `json: products`
}

type Response struct {
	Products []Product `json: "products"`
}

func main() {
	str := `{"name": "my product", "id": 1}`
	product := Product{}
	json.Unmarshal([]byte(str), &product)
	fmt.Println(product)

	file, _ := ioutil.ReadFile("products.json")

	data := Products{}

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.Products); i++ {
		fmt.Println("Product Id: ", data.Products[i].Id)
		fmt.Println("Name: ", data.Products[i].Name)
	}
}
