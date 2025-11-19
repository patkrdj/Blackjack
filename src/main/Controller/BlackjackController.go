package Controller

import (
	"Blackjack/src/main/Model"
	"Blackjack/src/main/View"
)

func playRound(inputView *View.InputView, outputView *View.OutputView, deck *Model.Deck) error {
	bet, err := inputView.ReadBet()
	if err != nil {
		return err
	}

	round := Model.NewRound(deck, bet)

	if err = round.DealInit(); err != nil {
		return err
	}

	if err = outputView.PrintDealerCard(*round); err != nil {
		return err
	}

	if err = outputView.PrintPlayerCards(*round); err != nil {
		return err
	}

	for !round.GetPlayerHand().IsBurst() {
		if round.IsBlackjack() {
			break
		}

		opt, err := inputView.ReadOption()
		if err != nil {
			return err
		}

		if opt == 1 {
			if err = round.PlayerHit(); err != nil {
				return err
			}
			if err = outputView.PrintDealerCard(*round); err != nil {
				return err
			}
			if err = outputView.PrintPlayerCards(*round); err != nil {
				return err
			}
		} else {
			break
		}
	}

	if !round.GetPlayerHand().IsBurst() && !round.IsBlackjack() {
		if err = round.DealerTurn(); err != nil {
			return err
		}
	}

	if err = outputView.PrintResult(*round); err != nil {
		return err
	}

	return nil
}

func Run() error {
	inputView := View.NewInputView()
	outputView := View.NewOutputView()

	deck := Model.NewDeck(nil)
	for {
		if deck.IsNeedToShuffle() {
			deck = Model.NewDeck(nil)
		}
		if err := playRound(inputView, outputView, deck); err != nil {
			return err
		}
	}
}
