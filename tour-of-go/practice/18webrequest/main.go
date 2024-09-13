package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://lco.dev"

func main() {
	fmt.Println("LCO web request")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Respose is fo type of %T\n", response)

	defer response.Body.Close() //caller resposibility to close the connection

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(data)
	fmt.Println(content)

}
