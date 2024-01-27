package cards

import (
	"fmt"
	"github.com/samber/lo"
)

type Suit int
type Face int

const (
	Clubs Suit = iota
	Diamonds
	Spades
	Hearts
)

var Suits = []Suit{Clubs, Diamonds, Spades, Hearts}

const (
	Two Face = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

var Faces = []Face{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

type Card struct {
	suit Suit
	face Face
}

func (c Card) Suit() Suit {
	return c.suit
}

func (c Card) Face() Face {
	return c.face
}


func isValidSuit(s Suit) bool {
	return lo.Contains(Suits, s)
}

func isValidFace(f Face) bool {
	return lo.Contains(Faces, f)
}

func NewCard(face Face, suit Suit) (*Card, error) {
	if (!isValidFace(face) || !isValidSuit(suit)) {
		return nil, fmt.Errorf("Cannot construct Card with Face={%d}, Suit={%d}", face, suit)
	} 

	return &Card{
		face: face,
		suit: suit,
	}, nil
}

const ValidDeckSize = 52

type Deck struct {
	left []Card
	drawn []Card
}

func (d Deck) allCards() []Card {
	return lo.Union(d.left, d.drawn)
}

func (d Deck) size() int {
	return len(d.left) + len(d.drawn);
}


func (d Deck) isUnique() bool {
	uniques := lo.FindUniques(d.allCards())
	return len(uniques) == len(d.allCards())
}

func (d Deck) isDirty() bool {
	return isValidDeck(d) && len(d.left) != ValidDeckSize
}

func isValidDeck(d Deck) bool {
	return d.isUnique() && d.size() == ValidDeckSize
}

func (d Deck) IsEmpty() bool {
	return len(d.left) == 0
}

func (d *Deck) Collect() {
	d.left = lo.Union(d.left, d.drawn)
	d.drawn = []Card{}
}

func (d *Deck) Shuffle() {
	shuffled := lo.Shuffle(d.left)
	d.left = shuffled
}

func (d *Deck) CollectAndShuffle() {
	d.Collect()
	d.Shuffle()
}

func (d *Deck) Draw() (*Card, error) {
	if d.IsEmpty() {
		return nil, fmt.Errorf("Deck is Empty, {%d} drawn cards and {%d} left cards", len(d.drawn), len(d.left))
	}

	drawn := &d.left[0]
	d.drawn = append(d.drawn, *drawn)
	d.left = d.left[1:]
	return drawn, nil
}

func NewDeckWithoutValidation(cards []Card) Deck {
	d := Deck{
		left: cards,
		drawn: []Card{},
	}
	return d
}

func NewDeck(cards []Card) (*Deck, error) {
	d := NewDeckWithoutValidation(cards)
	if !isValidDeck(d) {
		return nil, fmt.Errorf("Deck is invalid. Cards = {%v}", cards)
	}
	return &d, nil
}

func NewFullDeck() Deck {
	cards := []Card{}
	for _, face := range Faces {
		for _, suit := range Suits {
			c, _ := NewCard(face, suit)
			cards = append(cards, *c)
		}
	}
	return NewDeckWithoutValidation(cards)
}
