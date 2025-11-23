package Model

import "strconv"

type Hand struct {
	role  string
	cards []Card
}

func NewHand(role string) *Hand {
	return &Hand{role, make([]Card, 0)}
}

func (hand *Hand) AddCard(card Card) {
	hand.cards = append(hand.cards, card)
}

func (hand *Hand) GetCards() []Card {
	return hand.cards
}

func (hand *Hand) calc() (int, int) {
	cardSum := 0
	aceCount := 0

	for _, card := range hand.cards {
		if card.value == "A" {
			cardSum += 11
			aceCount++
		} else if card.value == "J" || card.value == "Q" || card.value == "K" {
			cardSum += 10
		} else {
			number, _ := strconv.Atoi(card.value)
			cardSum += number
		}
	}

	for cardSum > 21 && aceCount > 0 {
		cardSum -= 10
		aceCount--
	}

	return cardSum, aceCount
}

func (hand *Hand) GetSum() int {
	sum, _ := hand.calc()
	return sum
}

func (hand *Hand) IsSoftHand() bool {
	_, aceCount := hand.calc()
	return aceCount > 0
}

func (hand *Hand) IsBurst() bool {
	sum := hand.GetSum()
	if sum > 21 {
		return true
	}
	return false
}
