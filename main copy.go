package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"GOLangAPIIIFL/class"

	"github.com/gorilla/mux"
)

type article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var articles []article

type configuration struct {
	ServiceURL string `json:"ServiceURL"`
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	//	vars := mux.Vars(r)
	//	key := vars["id"]
	//	name := vars["name"]

	//	fmt.Fprintf(w, "Key: "+key)
	//	fmt.Fprintf(w, " name: "+name)

	e := class.Employee{
		FirstName:   "Sam",
		LastName:    "Adolf",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	json.NewEncoder(w).Encode(e)
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	// replace http.HandleFunc with myRouter.HandleFunc
	//myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article/{id}/{name}", returnSingleArticle)
	myRouter.HandleFunc("/checkConfig", checkConfig)

	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func checkConfig(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	Configur := configuration{}
	err := decoder.Decode(&Configur)
	if err != nil {
		fmt.Println("error:", err)
	}
	json.NewEncoder(w).Encode(Configur)
	fmt.Println(Configur.ServiceURL) // output: [UserA, UserB]
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	//fmt.Println("Programmer = ", json.NewEncoder(w).Encode(e))
	handleRequests()
}
