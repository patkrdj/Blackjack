package Model

import (
	"Blackjack/src/main/Model"
	"testing"
)

func Test_GetSum_IsBurst(t *testing.T) {
	hand := Model.NewHand("player")
	hand.AddCard(Model.NewCard("♠", "A"))
	hand.AddCard(Model.NewCard("♠", "Q"))

	if hand.GetSum() != 21 || hand.IsBurst() {
		t.Fail()
	}
}

func Test_Burst(t *testing.T) {
	hand := Model.NewHand("player")
	hand.AddCard(Model.NewCard("♠", "A"))
	hand.AddCard(Model.NewCard("♠", "7"))
	hand.AddCard(Model.NewCard("♠", "Q"))
	hand.AddCard(Model.NewCard("♠", "5"))

	if hand.GetSum() != 23 || !hand.IsBurst() {
		t.Fail()
	}
}
