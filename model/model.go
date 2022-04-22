package model

type DrawCard struct {
	Card []Card `json:"cards"`
}

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
}
