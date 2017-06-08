package handler

import (
	"net/http"
	"fmt"
	"log"
	"encoding/json"
	"github.com/rogersole/BasicBattleshipAPIGo/models"
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
	var msg ResponseMessage
	if r.Body == nil {
		msg = ResponseMessage {Message: fmt.Sprintf("%v", err.Error())}
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(msg); err != nil {
			panic(err)
		}
		return
	}

	var ships [][]int
	err = json.NewDecoder(r.Body).Decode(&ships)
	if err != nil {
		msg = ResponseMessage {Message: fmt.Sprintf("%v", err.Error())}
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(msg); err != nil {
			panic(err)
		}
		return
	}

	log.Printf("Ships: %v", ships)

	b, err = models.NewBoard(ships)
	if err != nil {
		msg = ResponseMessage{Message: fmt.Sprintf("%v", err.Error())}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(msg); err != nil {
			panic(err)
		}
		return
	}

	b.ShowBattlefield()

	msg = ResponseMessage {Message: "Board initialized successfully!"}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		panic(err)
	}
}

func UpdateTurn(w http.ResponseWriter, r *http.Request) {

	var msg ResponseMessage
	xPos, err := strconv.Atoi(r.URL.Query().Get("x"))
	if err != nil {
		msg = ResponseMessage{Message: fmt.Sprintf("%v", err.Error())}
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(msg); err != nil {
			panic(err)
		}
		return
	}

	yPos, err := strconv.Atoi(r.URL.Query().Get("y"))
	if err != nil {
		msg = ResponseMessage{Message: fmt.Sprintf("%v", err.Error())}
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(msg); err != nil {
			panic(err)
		}
		return
	}

	result, err := b.Attack(xPos, yPos)
	if err != nil {
		msg = ResponseMessage{Message: fmt.Sprintf("%v", err.Error())}
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(msg); err != nil {
			panic(err)
		}
		return
	}

	log.Printf("Response: %v", result)

	b.ShowBattlefield()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	msg = ResponseMessage{Message: result}
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		panic(err)
	}
}

