package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Diary struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Data    string `json:"created data"`
}

var diaries []Diary

func getDiaries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(diaries)
}

func deleteDiary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range diaries {
		if item.ID == params["id"] {
			diaries = append(diaries[:index], diaries[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(diaries)
}

func getDiary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range diaries {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createDiary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var diary Diary
	_ = json.NewDecoder(r.Body).Decode(&diary)
	diary.ID = strconv.Itoa(rand.Intn(10000000))
	diaries = append(diaries, diary)
	json.NewEncoder(w).Encode(diary)

}

func updateDiary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range diaries {
		if item.ID == params["id"] {
			diaries = append(diaries[:index], diaries[index+1:]...)
			var diary Diary
			_ = json.NewDecoder(r.Body).Decode(&diary)
			diary.ID = params["id"]
			diaries = append(diaries, diary)
			json.NewEncoder(w).Encode(diary)
			return
		}
	}
}

func main() {

	diaries = append(diaries, Diary{ID: "1", Content: "First note", Data: "15:28"})
	diaries = append(diaries, Diary{ID: "2", Content: "Second note", Data: "15:30"})
	diaries = append(diaries, Diary{ID: "3", Content: "Third note", Data: "15:32"})
	r := mux.NewRouter()
	r.HandleFunc("/diaries", getDiaries).Methods("GET")
	r.HandleFunc("/diaries/{id}", getDiary).Methods("GET")
	r.HandleFunc("/diaries", createDiary).Methods("POST")
	r.HandleFunc("/diaries/{id}", updateDiary).Methods("PUT")
	r.HandleFunc("/diaries/{id}", deleteDiary).Methods("DELETE")

	fmt.Println("Starting the server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
