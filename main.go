// Features:
//  POST /card : will add a flashcard
//  GET /quiz : will randmonly get a flashcard's question

package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Flashcard struct {
	Question string
	Answer   string
}

var (
	Flashcards = make(map[int][]Flashcard)
	counter    = 0
	mu         sync.Mutex
)

func main() {
	http.HandleFunc("/card", addCard)
	http.HandleFunc("/quiz", randomCardGenerator)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
func addCard(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var card Flashcard
		err := json.NewDecoder(r.Body).Decode(&card)
		if err != nil || card.Question == "" || card.Answer == "" {
			http.Error(w, "Invalid JSON input", http.StatusBadRequest)
			return
		}
		mu.Lock()
		Flashcards[counter] = append(Flashcards[counter], card)
		counter++
		mu.Unlock()
		json.NewEncoder(w).Encode(card)
		return
	} else {
		http.Error(w, "Wrong method used", http.StatusMethodNotAllowed)
		return
	}
}

func randomCardGenerator(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var rng = rand.New(rand.NewSource(time.Now().UnixNano()))
		randomNumber := rng.Intn(len(Flashcards))
		if len(Flashcards) == 0 {
			fmt.Fprintf(w, "Please add any flashcard before trying to access them.")
			return
		}
		randomcard := Flashcards[randomNumber]
		json.NewEncoder(w).Encode(randomcard)
		return

	} else {
		http.Error(w, "Wrong method used", http.StatusMethodNotAllowed)
		return
	}
}
