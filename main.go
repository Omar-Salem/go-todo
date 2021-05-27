package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
)

type TodoItem struct {
	Title string `json:"Title"`
}

var todos []TodoItem

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello go!")
	fmt.Println("logging...")
}

func getAll(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(todos)
}
func getBydId(w http.ResponseWriter, r *http.Request){
	i, _ := strconv.Atoi(r.URL.Path[len("/view/"):])
	json.NewEncoder(w).Encode(todos[i])
}

func handleRequests() {
	http.HandleFunc("/", getAll)
	http.HandleFunc("/view/", getBydId)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	todos = []TodoItem{
		TodoItem{Title: "Install Go"},
		TodoItem{Title: "Use Go"},
	}

	handleRequests()
}