package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rohandas-max/apimongo/router"
)

func main() {

	fmt.Println("Noob Netflix")
	r := router.Router()
	fmt.Println("API is starting ....")
	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("API is up and running")

}
