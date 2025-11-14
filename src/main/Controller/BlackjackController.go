package Controller

import (
	"Blackjack/src/main/Model"
	"Blackjack/src/main/View"
)

func run() error {
	inputView := View.NewInputView()
	bet, err := inputView.ReadBet()
	if err != nil {
		return err
	}

	round := Model.NewRound(Model.NewDeck(nil), bet)

	inputView.Read
}
