package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleHomePage)
	http.HandleFunc("/welcome", handleWelcomePage)
	http.ListenAndServe(":8080", nil)

}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("welcome to home page")
	fmt.Fprintf(w, "welcome to home page")
}

func handleWelcomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("welcome to Main Welcome page")
	fmt.Fprintf(w, "welcome to Main Welcome page")
}
