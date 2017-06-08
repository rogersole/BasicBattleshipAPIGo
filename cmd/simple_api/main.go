package main

import (
	"log"
	"net/http"
	"github.com/rogersole/simple_api/cmd/simple_api/router"
)

// curl -H "Content-Type: application/json" -X POST -d '[[4, 8],[3, 2],[8, 4],[0,9]]' http://localhost:8080/game
// curl -X PUT http://localhost:8080/game?x=4&y=8
func main() {
	router := router.NewRouter()
	log.Println("Initializing Server at 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}


