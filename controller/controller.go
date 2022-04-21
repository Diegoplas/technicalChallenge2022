package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Diegoplas/technicalChallenge2022/configuration"
	"github.com/Diegoplas/technicalChallenge2022/model"
)

func HelloWize(w http.ResponseWriter, r *http.Request) {
	helloString := "Hello, Wizeline!"
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helloString)
}

func PlayCards(w http.ResponseWriter, r *http.Request) {
	computerValue, computerSuite, err := GetCard()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error playing card game")
	}
	userValue, userSuite, err := GetCard()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error playing card game")
	}
	winnerPhrase := whoWinsTheGame(computerValue, userValue, computerSuite, userSuite)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(winnerPhrase)
}

func GetCard() (value int, suit string, err error) {
	var card model.DrawCard
	responsCard, err := http.Get(configuration.DrawACard)
	if err != nil {
		return 0, "", errors.New("something happened, no cards for now :(")
	}
	jsonCard, err := ioutil.ReadAll(responsCard.Body)
	if err != nil {
		return 0, "", errors.New("error reading card info")
	}
	err = json.Unmarshal([]byte(jsonCard), &card)
	if err != nil {
		panic(err)
	}
	cardValue := card.Card[0].Value
	cardSuit := card.Card[0].Suit
	fmt.Println("looking val", cardValue)
	value = ConvertCardValueToInt(cardValue)
	return value, cardSuit, nil
}

func ConvertCardValueToInt(strValue string) int {
	switch strValue {
	case "ACE":
		return 14
	case "KING":
		return 13
	case "QUEEN":
		return 12
	case "JACK":
		return 11
	default:
		value, _ := strconv.Atoi(strValue)
		return value
	}
}

func whoWinsTheGame(computerValue, userValue int, computerSuit, userSuit string) string {
	resultString := ""
	if computerValue > userValue {
		resultString = fmt.Sprintf("Computer Wins!! %v of %s against %v of %s",
			computerValue, computerSuit, userValue, userSuit)
	} else if computerValue < userValue {
		resultString = fmt.Sprintf("User Wins!! %v of %s against %v of %s",
			userValue, userSuit, computerValue, computerSuit)
	} else {
		resultString = fmt.Sprintf("Is a Tie!! both got %v's",
			userValue)
	}
	return resultString
}
