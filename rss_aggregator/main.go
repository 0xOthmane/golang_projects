package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main()  {
	fmt.Println("Hello world")
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found")
	}
	fmt.Println("PORT:", portString)
}