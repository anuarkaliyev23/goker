package cmd

import (
	"fmt"
	"strings"

	"github.com/anuarkaliyev23/goker/pkg/cards"
)

const Ace string = "A"
const King string = "K"
const Queen string = "Q"
const Jack string = "J"
const Ten string = "T"
const Nine string = "9"
const Eight string = "8"
const Seven string = "7"
const Six string = "6"
const Five string = "5"
const Four string = "4"
const Three string = "3"
const Two string = "2"
	
const Clubs string = "C"
const Spades string = "S"
const Diamonds string = "D"
const Hearts string = "H"

const validCardStringSize = 2

func parseFace(face string) (cards.Face, error) {
	switch face {
	case Ace: return cards.Ace, nil
	case King: return cards.King, nil
	case Queen: return cards.Queen, nil
	case Jack: return cards.Jack, nil
	case Ten: return cards.Ten, nil
	case Nine: return cards.Nine, nil
	case Eight: return cards.Eight, nil
	case Seven: return cards.Seven, nil
	case Six: return cards.Six, nil
	case Five: return cards.Five, nil
	case Four: return cards.Four, nil
	case Three: return cards.Three, nil
	case Two: return cards.Two, nil
	default: return 0, fmt.Errorf("Cannot parse face from {%s}", face)
	}
}

func parseSuit(suit string) (cards.Suit, error) {
	switch suit {
	case Clubs: return cards.Clubs, nil
	case Spades: return cards.Spades, nil
	case Diamonds: return cards.Diamonds, nil
	case Hearts: return cards.Hearts, nil
	default: return 0, fmt.Errorf("Cannot parse suit from {%s}", suit)
	}
}

func ParseCard(card string) (*cards.Card, error) {
	if len(card) != validCardStringSize {
		return nil, fmt.Errorf("Cannot parse card from rune {%s}, length must be {%d}", card, len(card))
	}
	card = strings.ToUpper(card)

	face, err := parseFace(string(card[0]))
	if err != nil {
		return nil, err
	}

	suit, err := parseSuit(string(card[1]))
	if err != nil {
		return nil, err
	}

	return cards.NewCard(face, suit)
}

func ParseCards(representation string) ([]cards.Card, error) {
	if len(representation) % 2 != 0 {
		return nil, fmt.Errorf("Cannot parse cards from {%s}, length is not even", representation)
	}
	
	parsed := []cards.Card{}

	for i := 0; i < len(representation); i += 2 {
		sub := representation[i:i + 2]
		card, err := ParseCard(sub)
		if err != nil {
			return nil, err
		}
		
		parsed = append(parsed, *card)
	}
	
	return parsed, nil
}
