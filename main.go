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
	s := r.URL.Path[len("/view/"):]
	i, _ := strconv.Atoi(s)
	json.NewEncoder(w).Encode(todos[i])
}

func deleteBydId(w http.ResponseWriter, r *http.Request) {
	if http.MethodDelete != r.Method {
		http.Error(w, "Bad Method", http.StatusBadRequest)
		return
	}

	s := r.URL.Path[len("/delete/"):]
	i, _ := strconv.Atoi(s)

	todos = removeAtIndex(todos, i)
	json.NewEncoder(w).Encode(todos)
}

func handleRequests() {
	http.HandleFunc("/", getAll)
	http.HandleFunc("/view/", getBydId)
	http.HandleFunc("/delete/", deleteBydId)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func removeAtIndex(s []TodoItem, i int) []TodoItem {
	// bring element to remove at the end if its not there yet
	if i != len(s)-1 {
		s[i] = s[len(s)-1]
	}

	// drop the last element
	return s[:len(s)-1]
}

func main() {
	todos = []TodoItem{
		TodoItem{Title: "Install Go"},
		TodoItem{Title: "Use Go"},
	}

	handleRequests()
}