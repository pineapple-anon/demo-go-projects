package main

import (
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	log.Printf("File content: %s", data)
	log.Println("File read successfully")
}