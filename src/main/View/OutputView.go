package View

import (
	"Blackjack/src/main/Model"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type OutputView struct {
	Writer *bufio.Writer
}

func NewOutputView() *OutputView {
	return &OutputView{bufio.NewWriter(os.Stdout)}
}

func (outputView *OutputView) PrintPlayerCards(round Model.Round) error {
	playerHand := round.GetPlayerHand()
	cards := playerHand.GetCards()
	_, err := outputView.Writer.WriteString("현재 플레이어의 카드\n")
	if err != nil {
		return err
	}

	var cardsTemp []string
	for _, card := range cards {
		cardsTemp = append(cardsTemp, card.GetSuit()+card.GetValue())
	}
	cardLine := strings.Join(cardsTemp, ", ")
	_, err = outputView.Writer.WriteString(cardLine + "\n")
	if err != nil {
		return err
	}

	_, err = outputView.Writer.WriteString("현재 카드의 합은 " + strconv.Itoa(round.GetPlayerHand().GetSum()) + "입니다.\n")
	if err != nil {
		return err
	}

	err = outputView.Writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

func (outputView *OutputView) PrintDealerCard(round Model.Round) error {
	playerHand := round.GetPlayerHand()
	cards := playerHand.GetCards()
	_, err := outputView.Writer.WriteString("딜러의 카드: " + cards[0].GetSuit() + cards[0].GetValue() + "\n")
	if err != nil {
		return err
	}
	err = outputView.Writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

func (outputView *OutputView) PrintResult(round Model.Round) error {
	dealerHand := round.GetDealerHand()
	cards := dealerHand.GetCards()
	_, err := outputView.Writer.WriteString("딜러의 카드\n")
	if err != nil {
		return err
	}

	var cardsTemp []string
	for _, card := range cards {
		cardsTemp = append(cardsTemp, card.GetSuit()+card.GetValue())
	}
	cardLine := strings.Join(cardsTemp, ", ")
	_, err = outputView.Writer.WriteString(cardLine + "\n")
	if err != nil {
		return err
	}

	_, err = outputView.Writer.WriteString("딜러의 카드 합은 " + strconv.Itoa(round.GetDealerHand().GetSum()) + "입니다.\n")
	if err != nil {
		return err
	}

	resultAmount := float64(round.GetBet()) + float64(round.GetBet())*round.CalculatePayout()
	_, err = outputView.Writer.WriteString("결과 금액은 " + strconv.Itoa(int(resultAmount)) + "입니다.\n")
	if err != nil {
		return err
	}

	err = outputView.Writer.Flush()
	if err != nil {
		return err
	}
	return nil
}
