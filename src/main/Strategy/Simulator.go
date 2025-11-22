package Strategy

import (
	"Blackjack/src/main/Model"
)

type SimResult struct {
	Wins       int     // 승리 횟수
	Draws      int     // 무승부 횟수
	Losses     int     // 패배 횟수
	TotalBet   int     // 총 베팅 금액
	TotalEarn  float64 // 총 수익 금액 (순수익)
	WinRate    float64 // 승률 (%)
	ReturnRate float64 // 수익률 (%)
}

func playRound(deck *Model.Deck, strategy Strategy) float64 {
	bet := strategy.DecideBetting()
	round := Model.NewRound(deck, bet)

	_ = round.DealInit()

	for !round.GetPlayerHand().IsBurst() {
		if round.IsBlackjack() {
			break
		}

		action := strategy.DecideAction(round.GetPlayerHand(), round.GetDealerHand().GetCards()[0])

		if action == Hit {
			_ = round.PlayerHit()
		} else {
			break
		}
	}

	if !round.GetPlayerHand().IsBurst() && !round.IsBlackjack() {
		_ = round.DealerTurn()
	}

	strategy.OnRoundEnd(round.GetPlayerHand(), round.GetDealerHand())

	return float64(bet) * round.CalculatePayout()
}

func playSingleGame(round int, strategy Strategy) float64 {
	deck := Model.NewDeck(nil)
	earn := 0.0
	for i := 0; i < round; i++ {
		if deck.IsNeedToShuffle() {
			deck = Model.NewDeck(nil)
		}
		earn += playRound(deck, strategy)
	}
	return earn
}

func Run() float64 {
	earn := playSingleGame(10, newDealerStrategy(1000))
	return earn
}
