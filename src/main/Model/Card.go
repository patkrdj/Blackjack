package Model

import "strconv"

type Card struct {
	suit  string
	value string
}

func NewCard(suit string, value string) Card {
	return Card{suit: suit, value: value}
}

func (card *Card) GetSuit() string {
	return card.suit
}

func (card *Card) GetValue() string {
	return card.value
}

func (card *Card) GetCardValue() int {
	if card.value == "J" || card.value == "Q" || card.value == "K" {
		return 10
	}
	if card.value == "A" {
		return 11
	}
	value, _ := strconv.Atoi(card.value)
	return value
}
