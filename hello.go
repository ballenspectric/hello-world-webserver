package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	who := os.Getenv("WHO")

	if len(who) == 0 {
		who = "World"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, "+who+"!")
	})

	fmt.Printf("Starting hello world server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
