package Model

type Hand struct {
	role  string
	cards []Card
}

func NewHand(role string) *Hand {
	return &Hand{role, make([]Card, 0)}
}

func (hand *Hand) addCard(card Card) {
	hand.cards = append(hand.cards, card)
}

func (hand *Hand) GetCards() []Card {
	return hand.cards
}
