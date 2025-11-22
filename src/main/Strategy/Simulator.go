package Strategy

import (
	"Blackjack/src/main/Model"
	"fmt"
	"sync"
	"time"
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

func playRound(deck *Model.Deck, strategy Strategy) (float64, int, int) {
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

	payout := round.CalculatePayout()
	earn := float64(bet) * payout

	resultType := 0
	if payout > 0 {
		resultType = 1
	} else if payout < 0 {
		resultType = -1
	}

	return earn, bet, resultType
}

func playSingleGame(round int, strategy Strategy) SimResult {
	deck := Model.NewDeck(nil)
	var result SimResult

	for i := 0; i < round; i++ {
		if deck.IsNeedToShuffle() {
			deck = Model.NewDeck(nil)
		}
		earn, bet, outcome := playRound(deck, strategy)

		result.TotalBet += bet
		result.TotalEarn += earn
		if outcome == 1 {
			result.Wins++
		} else if outcome == -1 {
			result.Losses++
		} else {
			result.Draws++
		}
	}
	return result
}

func Run() SimResult {
	simulationCount := 1_000_000
	roundPerGame := 10
	defaultBetAmount := 1000

	start := time.Now()

	var wg sync.WaitGroup
	results := make(chan SimResult, simulationCount)

	for i := 0; i < simulationCount; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			simResult := playSingleGame(roundPerGame, newDealerStrategy(defaultBetAmount))
			results <- simResult
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var finalResult SimResult

	for res := range results {
		finalResult.Wins += res.Wins
		finalResult.Draws += res.Draws
		finalResult.Losses += res.Losses
		finalResult.TotalBet += res.TotalBet
		finalResult.TotalEarn += res.TotalEarn
	}

	totalGames := finalResult.Wins + finalResult.Draws + finalResult.Losses
	if totalGames > 0 {
		finalResult.WinRate = (float64(finalResult.Wins) / float64(totalGames)) * 100
	}

	if finalResult.TotalBet > 0 {
		finalResult.ReturnRate = (finalResult.TotalEarn / float64(finalResult.TotalBet)) * 100
	}

	fmt.Println("걸린 시간:", time.Since(start))

	return finalResult
}

func RunWithoutGoroutine() SimResult {
	simulationCount := 1_000_000
	roundPerGame := 10
	defaultBetAmount := 1000

	start := time.Now()

	results := make([]SimResult, simulationCount)

	for i := 0; i < simulationCount; i++ {
		results[i] = playSingleGame(roundPerGame, newDealerStrategy(defaultBetAmount))
	}

	var finalResult SimResult

	for _, res := range results {
		finalResult.Wins += res.Wins
		finalResult.Draws += res.Draws
		finalResult.Losses += res.Losses
		finalResult.TotalBet += res.TotalBet
		finalResult.TotalEarn += res.TotalEarn
	}

	totalGames := finalResult.Wins + finalResult.Draws + finalResult.Losses
	if totalGames > 0 {
		finalResult.WinRate = (float64(finalResult.Wins) / float64(totalGames)) * 100
	}

	if finalResult.TotalBet > 0 {
		finalResult.ReturnRate = (finalResult.TotalEarn / float64(finalResult.TotalBet)) * 100
	}

	fmt.Println("걸린 시간:", time.Since(start))

	return finalResult
}
