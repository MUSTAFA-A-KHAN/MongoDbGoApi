package main

import (
	"fmt"
	"log"
	"net/http"

	"mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("SERVER IS GETTING STARted...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at Port 4000 ...")
}
