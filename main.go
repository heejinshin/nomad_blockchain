package main

import (
	"net/http"
	"fmt"
	"log"
)

const port string = ":4000"

func main() {
	fmt.Println("listening")
	log.Fatal(http.ListenAndServe(port, nil))
}
