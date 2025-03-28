package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
)

func main() {
    fmt.Println("Hello, RSS Aggregator!")
	//load environment variable
	godotenv.Load()

	//get port from environment variable
	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("Port is not Found in the Environment")
	}
	fmt.Println("Port :" , portString)

	router := chi.NewRouter() //Creating a router object
	
	//initialized the server
	srv := &http.Server{
		Handler : router,
		Addr : ":" + portString,
	}
	log.Printf("Server starting on port %v" , portString)
	err := srv.ListenAndServe()
	if err!= nil{
		log.Fatal(err)
	}

}
