package Strategy

type BasicStrategy struct {
	defaultBet int
}

func NewBasicStrategy(defaultBet int) *BasicStrategy {
	return &BasicStrategy{defaultBet}
}

func (s *BasicStrategy) DecideBetAmount() int {
	//todo
	return s.defaultBet
}

func (s *BasicStrategy) DecideAction(cxt GameContext) Action {
	//todo
	return Hit
}
