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
	err = round.DealInit()
	if err != nil {
		return err
	}

	err = outputView.PrintDealerCard(*round)
	if err != nil {
		return err
	}

	err = outputView.PrintPlayerCards(*round)
	if err != nil {
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
			err = round.PlayerHit()
			if err != nil {
				return err
			}
			err = outputView.PrintDealerCard(*round)
			if err != nil {
				return err
			}
			err = outputView.PrintPlayerCards(*round)
			if err != nil {
				return err
			}
		} else {
			break
		}
	}

	if round.IsBlackjack() {
		err = round.DealerTurn()
		if err != nil {
			return err
		}
	}

	err = outputView.PrintResult(*round)
	if err != nil {
		return err
	}

	return nil
}
