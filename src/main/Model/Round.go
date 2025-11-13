package Model

type Round struct {
	deck       *Deck
	dealerHand *Hand
	playerHand *Hand
	bet        int
}

func NewRound(deck *Deck, bet int) *Round {
	deck.ShuffleDeck()
	dealerHand := NewHand("dealer")
	playerHand := NewHand("player")
	return &Round{deck, dealerHand, playerHand, bet}
}

func (round *Round) DealInit() error {
	hands := []*Hand{round.playerHand, round.playerHand, round.dealerHand, round.dealerHand}
	for _, hand := range hands {
		err := round.dealCard(hand)
		if err != nil {
			return err
		}
	}
	return nil
}

func (round *Round) PlayerHit() error {
	err := round.dealCard(round.playerHand)
	if err != nil {
		return err
	}
	return nil
}

func (round *Round) dealCard(hand *Hand) error {
	card, err := round.deck.DrawCard()
	if err != nil {
		return err
	}
	hand.AddCard(card)
	return nil
}

func (round *Round) GetDealerHand() *Hand {
	return round.dealerHand
}

func (round *Round) GetPlayerHand() *Hand {
	return round.playerHand
}
