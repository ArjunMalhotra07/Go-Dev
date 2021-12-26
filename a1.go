/*Creating a simple server which can handle HTTP request. 1)Home Page func: handles all requests to our root url
2) handle request func: will match url path 3)main func: act as entry point to our simple rest server
Creating simple rest api that returns a series of articles based on a get request made to a particular url*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	http.HandleFunc("/animals", allAnimals)
	log.Fatal(http.ListenAndServe(":8081", nil))
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
