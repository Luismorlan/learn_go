package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2+2 is %d", sum))
}

func AddValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/home", Home)
	http.HandleFunc("/about", About)

	log.Println("Starting server on port:", port)
	http.ListenAndServe(port, nil)
}
