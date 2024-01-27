package cards

import (
	"errors"
	"fmt"
	"sort"

	"github.com/samber/lo"
)

type combinationType int

const validCardsLength = 5
const validStraightSum = 10

const (
	typeHighCard combinationType = iota
	typePair
	typeTwoPair
	typeThreeOfAKind
	typeStraight
	typeFlush
	typeFullHouse
	typeFourOfAKind
	typeStraghtFlush
)

type Combination struct {
	cards []Card
}

func (r Combination) isFlush() bool {
	suits := lo.Map(r.cards, func(card Card, index int) Suit {
		return card.Suit()
	})

	suitCount := lo.Count(suits, suits[0])
	return suitCount == validCardsLength
}

func (r Combination) isStraight() bool {
	faces := lo.Map(r.cards, func(card Card, index int) Face {
		return card.Face()
	})

	containsAce := lo.Contains(faces, Ace)
	if containsAce {
		return isStraightWithAce(r.cards)
	} else {
		return isStraightNoAce(r.cards)
	}
}

func (r Combination) isStraightFlush() bool {
	return r.isFlush() && r.isStraight()
}

func isStraightWithAce(cards []Card) bool {
	sort.Sort(ByFace(cards))
	faces := lo.Map(cards, func(card Card, index int) Face {
		return card.Face()
	})

	containsKing := lo.Contains(faces, King)
	containsTwo := lo.Contains(faces, Two)

	if containsKing {
		containsQueen := lo.Contains(faces, Queen)
		containsJack := lo.Contains(faces, Jack)
		containsTen := lo.Contains(faces, Ten)

		return containsQueen && containsJack && containsTen
	} else if containsTwo {
		containsThree := lo.Contains(faces, Three)
		containsFour := lo.Contains(faces, Four)
		containsFive := lo.Contains(faces, Five)

		return containsThree && containsFour && containsFive
	} else {
		return false
	}
}

func isStraightNoAce(cards []Card) bool {	
	sort.Sort(ByFace(cards))
	faces := lo.Map(cards, func(card Card, index int) Face {
		return card.Face()
	})
	
	toInts := lo.Map(faces, func(face Face, index int) int {
		return int(face)
	})

	first := toInts[0]
	reduced := lo.Map(toInts, func(faceInt int, index int) int {
		return faceInt - first
	})

	sum := lo.Sum(reduced)

	uniques := lo.Uniq(reduced)
	return len(uniques) == len(cards) && sum == validStraightSum
}

func (r Combination) isFourOfAKind() bool {
	toFaceCounts := r.toFaceCounts()
	return lo.Contains(toFaceCounts, 4)
}

func (r Combination) isThreeOfAKind() bool {
	toFaceCounts := r.toFaceCounts()	
	return lo.Contains(toFaceCounts, 3)
}

func (r Combination) isTwoPair() bool {
	toFaceCounts := r.toFaceCounts()
	return lo.Count(toFaceCounts, 2) == 2 * 2
}

func (r Combination) isPair() bool {
	toFaceCounts := r.toFaceCounts()
	return lo.Count(toFaceCounts, 2) == 2
}

func (r Combination) isFullHouse() bool {
	toFaceCounts := r.toFaceCounts()
	return lo.Count(toFaceCounts, 3) == 3 && lo.Count(toFaceCounts, 2) == 2
}

func (r Combination) isHighCard() bool {
	return !r.isStraight() && !r.isFlush() && !r.isFourOfAKind() && !r.isThreeOfAKind() && !r.isFullHouse() && !r.isTwoPair() && !r.isPair()
}

func (r Combination) toFaces() []Face {
	return lo.Map(r.cards, func(card Card, index int) Face {
		return card.Face()
	})
}

func (r Combination) toFaceCounts() []int {
	faces := r.toFaces()
	toCount := lo.Map(faces, func(face Face, index int) int {
		return lo.Count(faces, face)
	})
	return toCount
}

func NewCombination(cards []Card) (*Combination, error) {
	if len(cards) == validCardsLength {
		uniques := lo.Uniq(cards)
		if len(uniques) != validCardsLength {
			return nil, errors.New("cannot construct combination with not unique cards")
		}
		return &Combination{cards: cards}, nil
	} else {
		return nil, fmt.Errorf("cannot construct a combination you must pass slice of size: %d", validCardsLength)

	}
}
