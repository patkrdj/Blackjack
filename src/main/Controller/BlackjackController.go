package Controller

import (
	"Blackjack/src/main/Model"
	"Blackjack/src/main/View"
)

func Run() error {
	inputView := View.NewInputView()
	outputView := View.NewOutputView()

	bet, err := inputView.ReadBet()
	if err != nil {
		return err
	}

	round := Model.NewRound(Model.NewDeck(nil), bet)

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

	if round.IsBlackjack() {
		if err = round.DealerTurn(); err != nil {
			return err
		}
	}

	if err = outputView.PrintResult(*round); err != nil {
		return err
	}

	return nil
}
