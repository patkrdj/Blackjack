package Model

import (
	"Blackjack/src/main/Model"
	"testing"
)

func Test_Draw(t *testing.T) {
	deck := Model.NewDeck()
	card, err := deck.DrawCard()
	if err != nil {
		t.Error(err)
	}
	if card != Model.NewCard("♠", "A") {
		t.Errorf("덱의 제일 첫 장이 ♠A가 아님")
	}
}

func Test_Empty_Draw(t *testing.T) {
	deck := Model.Deck{}
	_, err := deck.DrawCard()
	if err == nil {
		t.Errorf("예상치 않은 에러 발생: %v", err)
	}
	if err != nil {
		t.Logf("에러가 정상적으로 동작")
	}
}
