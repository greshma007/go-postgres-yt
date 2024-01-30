package main

import (
	"fmt"
	"go-postgres-yt/router"
	"log"
	"net/http"
)

func main() {

	// call router() from Router.go
	r := router.Router()
	fmt.Println("Starting server on port 8080....")
	log.Fatal(http.ListenAndServe(":8080", r))

}
