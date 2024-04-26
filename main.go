package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func Setup() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}
}

func main() {
	Setup()

	fmt.Println("File uploaded successfully!!!")
}
