package handler

import (
	"net/http"
	"fmt"
	"log"
	"encoding/json"
	"github.com/rogersole/simple_api/models"
	"strconv"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

var b *models.Board

func InitializeGame(w http.ResponseWriter, r *http.Request) {
	log.Printf("Initializing game with params")

	var err error
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	var ships [][]int
	err = json.NewDecoder(r.Body).Decode(&ships)
	if err != nil {
		http.Error(w, err.Error(), 422) // unprocessable entity
		return
	}

	log.Printf("Ships: %v", ships)

	b, err = models.NewBoard(ships)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	b.ShowBattlefield()

	msg := ResponseMessage {
		Message: "Board initialized successfully!",
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		panic(err)
	}
}

func UpdateTurn(w http.ResponseWriter, r *http.Request) {

	xPos, err := strconv.Atoi(r.URL.Query().Get("x"))
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	yPos, err := strconv.Atoi(r.URL.Query().Get("y"))
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	result, err := b.Attack(xPos, yPos)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	log.Printf("Response: %v", result)

	b.ShowBattlefield()

	msg := ResponseMessage{
		Message: result,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		panic(err)
	}
}

