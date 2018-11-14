package main

import (
	"fmt"
	"net/http"
)

func handleMain(w http.ResponseWriter, r *http.Request) {

}

func handleDog(w http.ResponseWriter, r *http.Request) {

}

func handleMe(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "My name is Hongyi")

}

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/dog/", handleDog)
	http.HandleFunc("/me/", handleMe)

	http.ListenAndServe(":8080", nil)
}
