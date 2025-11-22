package Strategy

import "Blackjack/src/main/Model"

type DealerStrategy struct {
	defaultBet int
}

func newDealerStrategy(defaultBet int) *DealerStrategy {
	return &DealerStrategy{defaultBet}
}

func (s *DealerStrategy) DecideBetAmount() int {
	return s.defaultBet
}

func (s *DealerStrategy) DecideAction(playerHand Model.Hand, dealerCard Model.Card) Action {
	if playerHand.GetSum() <= 16 {
		return Hit
	}
	return Stand
}

func (s *DealerStrategy) OnRoundEnd(playerHand Model.Hand, dealerHand Model.Hand) {
	return
}
