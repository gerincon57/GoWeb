package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	leerJson()

}

func leerJson() {
	data, err := ioutil.ReadFile("../products.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("data:  ", string(data))
	var slice []string
	err = json.Unmarshal(data, &slice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("slice: %q\n", slice)
}
