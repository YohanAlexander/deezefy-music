package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func hello() string {
	message := "Hello World!"
	return message
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, hello())
}

func main() {
	http.HandleFunc("/", greet)
	log.Printf("Listening on %s...\n", os.Getenv("PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		panic(err)
	}
}
