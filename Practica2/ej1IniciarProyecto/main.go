package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	Id          int
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  string
	Price       float64
}

var productsList = []Product{}

func main() {
	leerJson(&productsList)

}

func leerJson(slice *[]Product) {
	data, err := os.ReadFile("../products.json")
	if err != nil {
		panic(err)

	}
	//fmt.Print("data:  ", string(data))

	err = json.Unmarshal([]byte(data), &slice)
	if err != nil {
		fmt.Println("cocurrio error:", err)
		panic(err)
	}
	fmt.Print("data:  ", string(data))

}
