package Strategy

type PlayerInfo struct {
	Sum   int
	Cards []string
}

type DealerInfo struct {
	UpCard string
}

type GameContext struct {
	Player PlayerInfo
	Dealer DealerInfo
}
