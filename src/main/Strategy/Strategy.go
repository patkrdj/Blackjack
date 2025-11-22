package Strategy

import "Blackjack/src/main/Model"

type Action int

const (
	Hit Action = iota
	Stand
)

type Strategy interface {
	DecideAction(playerHand Model.Hand, dealerCard Model.Card) Action
	DecideBetting() int
	OnRoundEnd(playerHand Model.Hand, dealerHand Model.Hand)
}
