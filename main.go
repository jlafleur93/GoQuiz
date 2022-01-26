package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(file)
}
