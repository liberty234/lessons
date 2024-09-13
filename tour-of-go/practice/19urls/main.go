package main

import (
	"fmt"
	"net/url"
)

const myurl string = "http://lco.dev:3000/learn?coursename=reactjs&paymentid=63hdsbbs"

func main() {
	fmt.Println("welcome to handling urls in golang")
	fmt.Println(myurl)

	// parsing

	result, _ := url.Parse(myurl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("the type of qparams: %T\n", qparams)
	fmt.Println(qparams["coursename"])
	for _, val := range qparams {
		fmt.Println("params is ", val)
	}

	partOfUrl := &url.URL{
		Scheme:  "http",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitech",
	}

	anotherUrl := partOfUrl.String()
	fmt.Println(anotherUrl)
}
