package main

import (
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	"html/template"
	"net/http"
	//	"encoding/json"
)

// template/basic/main.go

func main() {
	http.HandleFunc("/red-tea-pot", redTeaPotHandler)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		panic(err)
	}
}

// template/basic/main.go
func redTeaPotHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./views/product.html")
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, "test")

	// handle error
	type Product struct {
		Name        string
		Price       string
		Description string
	}

	teaPot := Product{Name: "Red Tea Pot 250ml", Description: "Test", Price: "19.99"}
	err = tmpl.Execute(w, teaPot)

	fmt.Printf("\n\n%+v", teaPot.Name)

}
