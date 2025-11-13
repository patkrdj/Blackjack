package Model

type Hand struct {
	role  string
	cards []Card
}

func NewHand(role string, card []Card) *Hand {
	return &Hand{role, card}
}
