package Strategy

type DealerStrategy struct {
	defaultBet int
}

func newDealerStrategy(defaultBet int) *DealerStrategy {
	return &DealerStrategy{defaultBet}
}

func (s *DealerStrategy) DecideBetAmount() int {
	return s.defaultBet
}

func (s *DealerStrategy) DecideAction(cxt GameContext) Action {
	if cxt.Player.Sum <= 16 {
		return Hit
	}
	return Stand
}
