package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to files in golang")

	content := "this needs to to save into the file - libertyebikade06@gmail.com"

	file, err := os.Create("./mygithubfiles.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	//if err != nil {
	//	panic(err)
	//}
	//create err in separate functions and use it in an err if
	checkNilErr(err)
	fmt.Println("the length", length)

	defer file.Close()
	readFile("./mygithubfiles.txt")
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)

	//if err != nil {
	//	panic(err)
	//}
	checkNilErr(err)
	fmt.Println("read the file:", string(data))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
