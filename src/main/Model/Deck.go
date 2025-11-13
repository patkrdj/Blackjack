package Model

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Deck struct {
	cards []Card
}

func NewDeck(cards []Card) *Deck {
	if cards != nil {
		return &Deck{cards: cards}
	}
	suits := []string{"♠", "♣", "♥", "♦"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	deck := make([]Card, 0)
	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, NewCard(suit, value))
		}
	}
	return &Deck{cards: deck}
}

func (deck *Deck) ShuffleDeck() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.cards), func(i, j int) {
		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	})
}

func (deck *Deck) PrintDeck() {
	for _, card := range deck.cards {
		fmt.Println(card.suit, card.value)
	}
}

func (deck *Deck) DrawCard() (Card, error) {
	if len(deck.cards) == 0 {
		return Card{}, errors.New("덱에 카드가 없습니다.")
	}
	card := deck.cards[0]
	deck.cards = deck.cards[1:]
	return card, nil
}
