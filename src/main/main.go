package main

import (
	"Blackjack/src/main/Strategy"
	"fmt"
	"strconv"
)

const (
	simulationCount = 1_000_000
	roundPerGame    = 10
	defaultBet      = 1000
)

func main() {
	// 사용자가 직접 블랙잭을 플레이하려는 경우 아래 주석을 제거한다.
	//if err := Controller.Run(); err != nil {
	//	panic(err)
	//}

	simulate(simulationCount, roundPerGame, Strategy.NewMyStrategy(defaultBet), true)
}

func simulate(simulateCount int, roundPerGame int, strategy Strategy.Strategy, goroutine bool) {
	var result = Strategy.SimResult{}
	if goroutine {
		result = Strategy.Run(simulateCount, roundPerGame, strategy)
	} else {
		result = Strategy.RunWithoutGoroutine(simulateCount, roundPerGame, strategy)
	}
	fmt.Println("승리 횟수:\t" + strconv.Itoa(result.Wins))
	fmt.Println("무승부 횟수:\t" + strconv.Itoa(result.Draws))
	fmt.Println("패배 횟수:\t" + strconv.Itoa(result.Losses))
	fmt.Println("총 배팅 금액:\t" + strconv.Itoa(result.TotalBet))
	fmt.Println("총 수익 금액:\t" + strconv.Itoa(int(result.TotalEarn)))
	rateStr := fmt.Sprintf("%.2f", result.ReturnRate)
	fmt.Println("총 수익률:\t" + rateStr + "%")
}
