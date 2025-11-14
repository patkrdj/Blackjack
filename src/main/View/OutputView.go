package View

import (
	"Blackjack/src/main/Model"
	"bufio"
	"os"
	"strconv"
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
	for _, card := range cards {
		_, err := outputView.Writer.WriteString(card.GetSuit() + card.GetValue() + ", ")
		if err != nil {
			return err
		}
	}
	_, err = outputView.Writer.WriteString("현재 카드의 합은 " + strconv.Itoa(round.GetPlayerHand().GetSum()) + "입니다.")
	if err != nil {
		return err
	}

	err = outputView.Writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
