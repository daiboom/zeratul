package main

import (
	"log"
	"net/http"

	"github.com/daiboom/zeratul/backend/server"
)

func main() {
	// dir, _ := os.Getwd()

	// fmt.Println("current path: " + dir)
	// fs := http.FileServer(http.Dir("./static"))
	// fmt.Println(fs)
	// http.Handle("/", fs)

	log.Print("Listening on:8888...")
	err := http.ListenAndServe(":8888", server.ServerHanlder())
	if err != nil {
		log.Fatal(err)
	}
}
