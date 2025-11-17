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

func (round *Round) DealerTurn() error {
	for round.dealerHand.GetSum() < 17 {
		err := round.dealCard(round.dealerHand)
		if err != nil {
			return err
		}
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

func (round *Round) CalculatePayout() float64 {
	playerSum := round.playerHand.GetSum()
	dealerSum := round.dealerHand.GetSum()

	if round.playerHand.IsBurst() {
		return -1
	}
	if round.dealerHand.IsBurst() {
		return 1
	}

	if playerSum == dealerSum {
		return 0
	} else if playerSum > dealerSum {
		if round.IsBlackjack() {
			return 1.5
		} else {
			return 1
		}
	} else {
		return -1
	}
}

func (round *Round) IsBlackjack() bool {
	if round.playerHand.GetSum() == 21 && len(round.playerHand.GetCards()) == 2 {
		return true
	}
	return false
}

func (round *Round) GetBet() int {
	return round.bet
}

func (round *Round) GetDealerHand() *Hand {
	return round.dealerHand
}

func (round *Round) GetPlayerHand() *Hand {
	return round.playerHand
}
