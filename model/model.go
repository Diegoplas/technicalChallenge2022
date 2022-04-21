package model

type DrawCard struct {
	Success bool   `json:"success"`
	Card    []Card `json:"cards"`
}

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
}
