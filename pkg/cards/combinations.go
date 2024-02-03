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
	uniques := lo.FindUniquesBy(r.cards, func(card Card) Face  {
		return card.Face()
	})

	sort.Sort(ByFace(uniques))
	return uniques
}

func (r Combination) MainCombinationCard() Face {
	if (comparedByHighestCard(r.Type())) {
		return r.HighestCardFace()
	}

	cardsDuplicatesByFace := lo.FindDuplicatesBy(r.cards, func(card Card) Face {
		return card.Face()
	})

	if len(cardsDuplicatesByFace) == 1 {
		return cardsDuplicatesByFace[0].Face()	
	} else {
		sort.Sort(ByFace(cardsDuplicatesByFace))
		return cardsDuplicatesByFace[len(cardsDuplicatesByFace) - 1].Face()
	}
}

func (r Combination) SecondaryCombinationCard() *Face {
	if (comparedBySecondary(r.Type())) {
		
		cardsDuplicatesByFace := lo.FindDuplicatesBy(r.cards, func(card Card) Face {
			return card.Face()
		})
		sort.Sort(ByFace(cardsDuplicatesByFace))
		result := cardsDuplicatesByFace[0].Face()
		return &result
	}
	return nil
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
	types := lo.Map(combinations, func(combination Combination, index int) CombinationType {
		return combination.Type()
	})

	if (types[0] == types[1]) {
		return false
	}

	if types[0] < types[1] {
		return true
	} else if types[0] == types[1] {
		if comparedByHighestCard(types[0]) {
			highestCards := lo.Map(combinations, func(combination Combination, index int) Face {
				return combination.HighestCardFace()
			})

			toInts := lo.Map(highestCards, func(face Face, index int) int {
				return int(face)
			})

			if toInts[0] < toInts[1] {
				return true
			} else if toInts[0] == toInts[1] {
				if comparedBySecondary(types[0]) {
					toSecondary := lo.Map(combinations, func(combination Combination, index int) *Face {
						return combination.SecondaryCombinationCard()
					})

					if int(*toSecondary[0]) < int(*toSecondary[1]) {
						return true
					}
				}

				if comparedByKickers(types[0]) {
					toKickers := lo.Map(combinations, func(combination Combination, index int) []Card {
						return combination.Kickers()
					})

					kickerFaces := lo.Map(toKickers, func(cards []Card, index int) []Face {
						faces := lo.Map(cards, func(card Card, index int) Face {
							return card.Face()
						})
						return faces
					})
					
					for face, index := range kickerFaces[0] {
						opposite := kickerFaces[1][index]
						if int(face) < int(opposite) {
							return true
						}
					}

				}
			}
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

