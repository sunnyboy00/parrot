package main

import (
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
	"encoding/json"
)

// Creates a new game and returns a session-id.
func createGame() (string, error) {
	resp, err := http.Get(BASE_URL + "start/json")
	if err != nil {
		return "", err
	}
	url := resp.Request.URL.String()
	slice := strings.Split(url, "/")
	return slice[len(slice) - 2], nil
}

func getState(sessionId string) (GameState, error) {
	resp, err := http.Get(BASE_URL + "state/" + sessionId + "/json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var state GameState
	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, &state)
	return state, err
}

func handToString(h Hand) string {
	if h == Up {
		return "0"
	} else if h == Right {
		return "1"
	} else if h == Down {
		return "2"
	} else {
		return "3"
	}
}

func intToHand(i int) Hand {
	if i == 0 {
		return Up
	} else if i == 1 {
		return Right
	} else if i == 2 {
		return Down
	} else if i == 3 {
		return Left
	} else {
		return Quit
	}
}

func sendHand(sessionId string, hand Hand) (GameState, error) {
	url := BASE_URL + "state/" + sessionId + "/move/" + handToString(hand) + "/json"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var state GameState
	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, &state)
	return state, err
}

func (s *GameState) showState() {
	fmt.Println("---------------------")
	for _, row := range s.Grid {
		fmt.Print("|")
		for _, v := range row {
			fmt.Printf("%4d|", v)
		}
		fmt.Println("")
		fmt.Println("---------------------")
	}
	fmt.Printf("points = %d, score = %d\n", s.Points, s.Score)
}