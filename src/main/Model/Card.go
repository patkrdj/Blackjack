package Model

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
