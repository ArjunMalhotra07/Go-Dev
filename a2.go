//Using gorilla mux router : to things like what verbs we can use to access what endpoints
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	myRouter.HandleFunc("/animals", allAnimals).Methods("GET")
	myRouter.HandleFunc("/animals", testPostAnimals).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Description"`
	Content string `json:"content"`
}
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Hello World"},
	}
	fmt.Println("Endpoint Hit: All articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test Post Articles EndPoint worked")
}

type Animal struct {
	Name  string `json:"Name"`
	Wings string `json:"Have wings or not"`
	Color string `json:"Color"`
}
type Animals []Animal

func allAnimals(w http.ResponseWriter, r *http.Request) {
	animals := Animals{
		Animal{Name: " Tiger", Wings: "No WIngs", Color: "Yellow"},
		Animal{Name: " Shark", Wings: "Have WIngs", Color: "Bluish"},
	}
	fmt.Println("Endpoint Hit: All Animals Endpoint")
	json.NewEncoder(w).Encode(animals)
}
func testPostAnimals(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test Post Animals EndPoint worked")
}
