package cards

import (
	"errors"
	"fmt"
	"sort"

	"github.com/samber/lo"
)

type CombinationType int

const validCardsLength = 5
const validStraightSum = 10

const (
	HighCard CombinationType = iota
	Pair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraghtFlush
)

func comparedByHighestCard(ctype CombinationType) bool {
	return ctype == StraghtFlush || ctype == Straight || ctype == Flush || ctype == HighCard
}


func comparedBySecondary(ctype CombinationType) bool {
	return ctype == TwoPair || ctype == FullHouse
}

func comparedByKickers(ctype CombinationType) bool {
	return ctype == HighCard || ctype == TwoPair || ctype == ThreeOfAKind || ctype == FourOfAKind || ctype == Flush
}


type Combination struct {
	cards []Card
}

func (r Combination) Type() CombinationType {
	if r.isStraightFlush() {
		return StraghtFlush
	} else if r.isFourOfAKind() {
		return FourOfAKind
	} else if r.isFullHouse() {
		return FullHouse
	} else if r.isFlush() {
		return Flush
	} else if r.isStraight() {
		return Straight
	} else if r.isThreeOfAKind() {
		return ThreeOfAKind
	} else if r.isTwoPair() {
		return TwoPair
	} else if r.isPair() {
		return Pair
	} else {
		return HighCard
	}
}

func (r Combination) HighestCardFace() Face {
	faces := r.toFaces()
	return faces[len(faces) - 1]
}

func (r Combination) Kickers() []Card {
	if r.Type() == HighCard || r.Type() == Flush {
		highestCardFace := r.HighestCardFace()
		filtered := lo.Filter(r.cards, func(card Card, index int) bool {
			return card.Face() != highestCardFace
		})

		sort.Sort(ByFaceReversed(filtered))
		return filtered
	}


	uniques := lo.FindUniquesBy(r.cards, func(card Card) Face  {
		return card.Face()
	})

	sort.Sort(ByFaceReversed(uniques))
	return uniques
}

func (r Combination) MainCard() Face {
	if (comparedByHighestCard(r.Type())) {
		return r.HighestCardFace()
	}

	cardsDuplicatesByFace := lo.FindDuplicatesBy(r.cards, func(card Card) Face {
		return card.Face()
	})

	if len(cardsDuplicatesByFace) == 1 {
		return cardsDuplicatesByFace[0].Face()	
	} else if r.Type() != FullHouse {
		sort.Sort(ByFaceReversed(cardsDuplicatesByFace))
		return cardsDuplicatesByFace[0].Face()
	} else {
		uniques := lo.FindDuplicatesBy(r.cards, func(card Card) Face {
			return card.Face()
		})

		firstCount := lo.CountBy(r.cards, func(card Card) bool {
			return card.Face() == uniques[0].Face()
		})

		if firstCount == 3 {
			return uniques[0].Face()
		} else {
			return uniques[1].Face()
		}
	}
}

func (r Combination) SecondaryCard() *Face {
	if !comparedBySecondary(r.Type()) {
		return nil
	}

	cardsDuplicatesByFace := lo.FindDuplicatesBy(r.cards, func(card Card) Face {
		return card.Face()
	})

	sort.Sort(ByFace(cardsDuplicatesByFace))
	result := cardsDuplicatesByFace[0].Face()
	return &result
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

func (r Combination) Less(other Combination) bool {
	combinations := []Combination{r, other}

	if r.Type() < other.Type() {
		return true
	} else if r.Type() == other.Type() {
		if comparedByHighestCard(r.Type()) {
			return lessByHighestCard(combinations)
		} else if comparedBySecondary(r.Type()) {
			return lessBySecondary(combinations)
		} else if comparedByKickers(r.Type()) {
			return lessByKickers(combinations)
		}
	}

	return false
}

func lessByHighestCard(combinations []Combination) bool {
	highestCards := lo.Map(combinations, func(combination Combination, index int) Face {
		return combination.HighestCardFace()
	})

	toInts := lo.Map(highestCards, func(face Face, index int) int {
		return int(face)
	})

	return toInts[0] < toInts[1]
}

func lessBySecondary(combinations []Combination) bool {
	toSecondary := lo.Map(combinations, func(combination Combination, index int) *Face {
		return combination.SecondaryCard()
	})
	
	if int(*toSecondary[0]) < int(*toSecondary[1]) {
		return true
	}
	
	return false
}

func lessByKickers(combinations []Combination) bool {
	toKickers := lo.Map(combinations, func(combination Combination, index int) []Card {
		return combination.Kickers()
	})

	first := lo.Map(toKickers[0], func(card Card, index int) Face {
		return card.Face()
	})

	second := lo.Map(toKickers[1], func(card Card, index int) Face {
		return card.Face()
	})

	for index, face := range first {
		opposite := second[index]
		if int(face) < int(opposite) {
			return true
		}
	}
	return false

}





func NewCombination(cards []Card) (*Combination, error) {
	sort.Sort(ByFace(cards))
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

