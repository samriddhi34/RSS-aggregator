package main

import (
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
)

func main() {
    fmt.Println("Hello, RSS Aggregator!")

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("Port is not Found in the Environment")
	}
	fmt.Println("Port :" , portString)
}
