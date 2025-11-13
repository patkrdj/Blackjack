package Model

import (
	"Blackjack/src/main/Model"
	"testing"
)

func Test_DealInit(t *testing.T) {
	cards := make([]Model.Card, 0)
	cards = append(cards, Model.NewCard("♠", "A"))
	cards = append(cards, Model.NewCard("♠", "2"))
	cards = append(cards, Model.NewCard("♠", "3"))
	cards = append(cards, Model.NewCard("♠", "4"))

	round := Model.NewRound(Model.NewDeck(cards), 1000)
	err := round.DealInit()
	if err != nil {
		t.Error(err)
	}

	dealer := round.GetDealerHand()
	player := round.GetPlayerHand()
	combine := append(dealer.GetCards(), player.GetCards()...)

	if len(combine) != len(cards) {
		t.Errorf("카드 개수가 다릅니다.")
	}

	for _, card := range cards {
		if contains(combine, card) == false {
			t.Errorf("전체 카드와 뽑은 카드가 다릅니다.")
		}
	}
}

func Test_DealInit_Input_Count_Exception(t *testing.T) {
	cards := make([]Model.Card, 0)
	cards = append(cards, Model.NewCard("♠", "A"))
	cards = append(cards, Model.NewCard("♠", "2"))
	cards = append(cards, Model.NewCard("♠", "3"))

	round := Model.NewRound(Model.NewDeck(cards), 1000)
	err := round.DealInit()
	if err == nil {
		t.Errorf("예상치 않은 에러 발생: %v", err)
	}
	if err != nil {
		t.Logf("에러가 정상적으로 동작")
	}
}

func contains(cards []Model.Card, target Model.Card) bool {
	for _, card := range cards {
		if card == target {
			return true
		}
	}
	return false
}
