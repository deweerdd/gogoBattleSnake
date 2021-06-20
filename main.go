package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

type MoveResponse struct {
	Move  string `json:"move"`
	Shout string `json:"shout,omitempty"`
}

type SnakeInfo struct {
	Apiversion string `json:"apiversion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
	Version    string `json:"version"`
}

type Game struct {
	Id      string `json:"id"`
	Timeout int    `json:"timeout"`
}

type GameRequest struct {
	Game        Game        `json:"game"`
	Turn        int         `json:"turn"`
	Board       Board       `json:"board"`
	Battlesnake Battlesnake `json:"you"`
}

type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Battlesnake struct {
	Id      string       `json:"id"`
	Name    string       `json:"name"`
	Health  int          `json:"health"`
	Body    []Coordinate `json:"body"`
	Latency string       `json:"latency"`
	Head    Coordinate   `json:"head"`
	Length  int          `json:"length"`
	Shout   string       `json:"shout"`
	Squad   string       `json:"squad"`
}

type Board struct {
	Height  int           `json:"height"`
	Width   int           `json:"width"`
	Food    []Coordinate  `json:"food"`
	Hazards []Coordinate  `json:"hazards"`
	Snakes  []Battlesnake `json:"snakes"`
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	response := SnakeInfo{
		Apiversion: "1",
		Author:     "ddeweerd",
		Color:      "cc00ff",
		Head:       "caffeine",
		Tail:       "round-bum",
		Version:    "0.0.1",
	}

	w.Header().Set("Content-Type", "application/Json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}

func HandleStart(w http.ResponseWriter, r *http.Request) {
	request := GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("START\n")
}

func HandleMove(w http.ResponseWriter, r *http.Request) {
	request := GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("getUnsafe")
	unsafe := getNeck(request)

	possibleMoves := []string{"up", "down", "left", "right"}
	move := possibleMoves[rand.Intn(len(possibleMoves))]

	for move == unsafe {
		move = possibleMoves[rand.Intn(len(possibleMoves))]
	}

	response := MoveResponse{
		Move: move,
	}

	fmt.Printf("MOVE: %s\n", response.Move)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}

func HandleEnd(w http.ResponseWriter, r *http.Request) {
	request := GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("END\n")
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/start", HandleStart)
	http.HandleFunc("/move", HandleMove)
	http.HandleFunc("/end", HandleEnd)
	fmt.Printf("Starting Battlesnake Server at http://localhost:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
