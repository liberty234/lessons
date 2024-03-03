package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type User1 struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	user1 := User{
		Name: "Cole Palmer",
		Age:  21,
	}

	jsondata, err := json.Marshal(user1)
	if err != nil {
		http.Error(w, "Json server error..", http.StatusInternalServerError)
	}
	if r.Method != http.MethodPost {
		http.Error(w, "only Post request are allow", http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")

	w.Write(jsondata)

}

func handlerUser1(w http.ResponseWriter, r *http.Request) {

	user1 := User1{
		FirstName: "Cole",
		LastName:  "Palmer",
		Age:       21,
	}

	jsondata, err := json.Marshal(user1)
	if err != nil {
		http.Error(w, "Json server error..", http.StatusInternalServerError)
	}
	if r.Method != http.MethodPost {
		http.Error(w, "only Post request are allow", http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")

	w.Write(jsondata)

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/user1", handlerUser1)
	fmt.Println("The server is running on port 3000...")
	http.ListenAndServe(":3000", nil)

}
