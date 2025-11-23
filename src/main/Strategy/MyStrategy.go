package Strategy

import "Blackjack/src/main/Model"

type MyStrategy struct {
	defaultBet   int
	runningCount int
}

func NewMyStrategy(defaultBet int) *MyStrategy {
	return &MyStrategy{defaultBet, 0}
}

func (s *MyStrategy) DecideBetting() int {
	if s.runningCount > 0 {
		return s.defaultBet * (1 + s.runningCount)
	}
	return s.defaultBet
}

func (s *MyStrategy) DecideAction(playerHand *Model.Hand, dealerCard Model.Card) Action {
	playerScore := playerHand.GetSum()
	dealerScore := dealerCard.GetCardValue()

	if !playerHand.IsSoftHand() {
		if playerScore >= 17 {
			return Stand
		}
		if playerScore >= 13 && playerScore <= 16 {
			if dealerScore >= 2 && dealerScore <= 6 {
				return Stand
			}
			return Hit
		}
		if playerScore == 12 {
			if dealerScore >= 4 && dealerScore <= 6 {
				return Stand
			}
		}
		return Hit
	}

	if playerHand.IsSoftHand() {
		if playerScore >= 19 {
			return Stand
		}
		if playerScore == 18 {
			if dealerScore >= 9 || dealerScore == 1 {
				return Hit
			}
			return Stand
		}
		return Hit
	}

	return Hit
}

func (s *MyStrategy) OnRoundEnd(playerHand *Model.Hand, dealerHand *Model.Hand) {
	for _, card := range dealerHand.GetCards() {
		s.updateCount(card)
	}
	for _, card := range playerHand.GetCards() {
		s.updateCount(card)
	}
	return
}

func (s *MyStrategy) updateCount(card Model.Card) {
	val := card.GetCardValue()
	if val >= 2 && val <= 6 {
		s.runningCount++
	} else if val == 10 || val == 11 || val == 1 {
		s.runningCount--
	}
}

func (s *MyStrategy) OnShuffleDeck() {
	s.runningCount = 0
}
