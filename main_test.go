package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var jsonString = []byte(`{
	"game": {
	  "id": "game-00fe20da-94ad-11ea-bb37",
	  "ruleset": {
		"name": "standard",
		"version": "v.1.2.3"
	  },
	  "timeout": 500
	},
	"turn": 14,
	"board": {
	  "height": 11,
	  "width": 11,
	  "food": [
		{"x": 5, "y": 5}, 
		{"x": 9, "y": 0}, 
		{"x": 2, "y": 6}
	  ],
	  "hazards": [
		{"x": 3, "y": 2}
	  ],
	  "snakes": [
		{
		  "id": "snake-508e96ac-94ad-11ea-bb37",
		  "name": "My Snake",
		  "health": 54,
		  "body": [
			{"x": 0, "y": 0}, 
			{"x": 1, "y": 0}, 
			{"x": 2, "y": 0}
		  ],
		  "latency": "111",
		  "head": {"x": 0, "y": 0},
		  "length": 3,
		  "shout": "why are we shouting??",
		  "squad": ""
		}, 
		{
		  "id": "snake-b67f4906-94ae-11ea-bb37",
		  "name": "Another Snake",
		  "health": 16,
		  "body": [
			{"x": 5, "y": 4}, 
			{"x": 5, "y": 3}, 
			{"x": 6, "y": 3},
			{"x": 6, "y": 2}
		  ],
		  "latency": "222",
		  "head": {"x": 5, "y": 4},
		  "length": 4,
		  "shout": "I'm not really sure...",
		  "squad": ""
		}
	  ]
	},
	"you": {
	  "id": "snake-508e96ac-94ad-11ea-bb37",
	  "name": "My Snake",
	  "health": 54,
	  "body": [
		{"x": 0, "y": 0}, 
		{"x": 1, "y": 0}, 
		{"x": 2, "y": 0}
	  ],
	  "latency": "111",
	  "head": {"x": 0, "y": 0},
	  "length": 3,
	  "shout": "why are we shouting??",
	  "squad": ""
	}
  }`)
var bodyReader = strings.NewReader(`{
	"game": {
	  "id": "game-00fe20da-94ad-11ea-bb37",
	  "ruleset": {
		"name": "standard",
		"version": "v.1.2.3"
	  },
	  "timeout": 500
	},
	"turn": 14,
	"board": {
	  "height": 11,
	  "width": 11,
	  "food": [
		{"x": 5, "y": 5}, 
		{"x": 9, "y": 0}, 
		{"x": 2, "y": 6}
	  ],
	  "hazards": [
		{"x": 3, "y": 2}
	  ],
	  "snakes": [
		{
		  "id": "snake-508e96ac-94ad-11ea-bb37",
		  "name": "My Snake",
		  "health": 54,
		  "body": [
			{"x": 0, "y": 0}, 
			{"x": 1, "y": 0}, 
			{"x": 2, "y": 0}
		  ],
		  "latency": "111",
		  "head": {"x": 0, "y": 0},
		  "length": 3,
		  "shout": "why are we shouting??",
		  "squad": ""
		}, 
		{
		  "id": "snake-b67f4906-94ae-11ea-bb37",
		  "name": "Another Snake",
		  "health": 16,
		  "body": [
			{"x": 5, "y": 4}, 
			{"x": 5, "y": 3}, 
			{"x": 6, "y": 3},
			{"x": 6, "y": 2}
		  ],
		  "latency": "222",
		  "head": {"x": 5, "y": 4},
		  "length": 4,
		  "shout": "I'm not really sure...",
		  "squad": ""
		}
	  ]
	},
	"you": {
	  "id": "snake-508e96ac-94ad-11ea-bb37",
	  "name": "My Snake",
	  "health": 54,
	  "body": [
		{"x": 0, "y": 0}, 
		{"x": 1, "y": 0}, 
		{"x": 2, "y": 0}
	  ],
	  "latency": "111",
	  "head": {"x": 0, "y": 0},
	  "length": 3,
	  "shout": "why are we shouting??",
	  "squad": ""
	}
  }`)

func TestGetSnakeInfo(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	HandleIndex(response, request)

	t.Run("Index Returns 200Status", func(t *testing.T) {
		got := response.Code
		want := http.StatusOK

		assertStatus(t, got, want)

	})

	t.Run("Index Returns Snake Info", func(t *testing.T) {
		got := response.Body.String()
		want := "{\"apiversion\":\"1\",\"author\":\"ddeweerd\",\"color\":\"cc00ff\",\"head\":\"caffeine\",\"tail\":\"round-bum\",\"version\":\"0.0.1\"}\n"

		assertString(t, got, want)
	})
}

func TestHandleStartGame(t *testing.T) {

	request, _ := http.NewRequest(http.MethodPost, "/start", bodyReader)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	HandleStart(response, request)

	t.Run("Start returns Ok Status", func(t *testing.T) {

		got := response.Result().StatusCode
		want := http.StatusOK

		assertStatus(t, got, want)
	})
}

func TestHandleMove(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/move", bytes.NewBuffer(jsonString))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	HandleMove(response, request)

	t.Run("Move returns Ok status", func(t *testing.T) {
		got := response.Result().StatusCode
		want := http.StatusOK

		assertStatus(t, got, want)
	})
	t.Run("Response Returns a move Direction", func(t *testing.T) {
		got := response.Body.String()
		want := "{\"move\":\"down\"}\n"

		assertString(t, got, want)
	})
}

func TestHandleEnd(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/move", bytes.NewBuffer(jsonString))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	HandleEnd(response, request)

	t.Run("Verify Status Code of End", func(t *testing.T) {
		got := response.Result().StatusCode
		want := http.StatusOK

		assertStatus(t, got, want)
	})

}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Incorrect Status got %d want %d", got, want)
	}
}
