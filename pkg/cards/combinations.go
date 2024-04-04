package cards

import (
	"errors"
	"fmt"
	"sort"

	"gonum.org/v1/gonum/stat/combin"

	"github.com/samber/lo"
)

type CombinationType int

const validCardsLength = 5
const validStraightSum = 10

const validCombinatoricsLength = 7

const (
	HighCard CombinationType = iota
	Pair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

var ShortDeckCombinationStrength = []CombinationType{HighCard, Pair, TwoPair, ThreeOfAKind, Straight, FullHouse, Flush, FourOfAKind, StraightFlush}
var DefaultCombinationStrength = []CombinationType{HighCard, Pair, TwoPair, ThreeOfAKind, Straight, Flush, FullHouse, FourOfAKind, StraightFlush}

func comparedByHighestCard(ctype CombinationType) bool {
	return ctype == StraightFlush || ctype == Straight || ctype == Flush || ctype == HighCard
}


func comparedBySecondary(ctype CombinationType) bool {
	return ctype == TwoPair || ctype == FullHouse
}

func comparedByKickers(ctype CombinationType) bool {
	return ctype == HighCard || ctype == Pair || ctype == TwoPair || ctype == ThreeOfAKind || ctype == FourOfAKind || ctype == Flush
}


type Combination struct {
	cards []Card
	combinationStrengths []CombinationType
	shortDeck bool
}

func (r Combination) AllCards() []Card {
	return r.cards
}

//TODO Known issue: This function will perform ok with short-deck vs classic since the only difference is flush/full house strength
//since it's impossible to have both at the same time. However, if other custom rules for strength will be passed (as design suggests) 
//there can be issues with the order of combination type resolving.
//For some reason I couldn't fix it right away, so I will postpone resolving this issue if it will come up some time after
func (r Combination) Type() CombinationType {

	if r.isStraightFlush() {
		return StraightFlush
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
		if r.isStraight() {
			if lo.Contains(r.toFaces(), Ace) && lo.Contains(r.toFaces(), Two) {
				return Five
			}
		}
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
	
	if r.Type() != FullHouse {
		cardsDuplicatesByFace := lo.FindDuplicatesBy(r.cards, func(card Card) Face {
			return card.Face()
		})

		sort.Sort(ByFace(cardsDuplicatesByFace))
		result := cardsDuplicatesByFace[0].Face()
		return &result
	} else {
		uniques := lo.FindDuplicatesBy(r.cards, func(card Card) Face {
			return card.Face()
		})

		firstCount := lo.CountBy(r.cards, func(card Card) bool {
			return card.Face() == uniques[0].Face()
		})

		if firstCount == 2 {
			result := uniques[0].Face()
			return &result
		} else {
			result := uniques[1].Face()
			return &result
		}
	}
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
		return isStraightWithAce(r.cards, r.shortDeck)
	} else {
		return isStraightNoAce(r.cards)
	}
}

func (r Combination) isStraightFlush() bool {
	return r.isFlush() && r.isStraight()
}


// TODO: there should be some elegant and generalized way to do it.
func isStraightWithAce(cards []Card, shortDeck bool) bool {
	sort.Sort(ByFace(cards))
	faces := lo.Map(cards, func(card Card, index int) Face {
		return card.Face()
	})

	containsKing := lo.Contains(faces, King)
	containsTwo := lo.Contains(faces, Two)
	containsSix := lo.Contains(faces, Six)

	if containsKing {
		containsQueen := lo.Contains(faces, Queen)
		containsJack := lo.Contains(faces, Jack)
		containsTen := lo.Contains(faces, Ten)

		return containsQueen && containsJack && containsTen
	} else if !shortDeck {
		if containsTwo {
			containsThree := lo.Contains(faces, Three)
			containsFour := lo.Contains(faces, Four)
			containsFive := lo.Contains(faces, Five)
			return containsThree && containsFour && containsFive
		} else {
			if containsSix {
				containsSeven := lo.Contains(faces, Seven)
				containsEight := lo.Contains(faces, Eight)
				containsNine := lo.Contains(faces, Nine)
				return containsSeven && containsEight && containsNine
			}
		}
	}
	return false
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

func (r Combination) Tie(other Combination) bool {
	return !r.Less(other) && !other.Less(r)
}

func (r Combination) More(other Combination) bool {
	return other.Less(r)
}

func (r Combination) CombinationStrengths() []CombinationType {
	if r.combinationStrengths == nil {
		return DefaultCombinationStrength
	} else {
		return r.combinationStrengths
	}
}

func (r Combination) lessByType(first CombinationType, second CombinationType) bool {
	firstIndex := lo.IndexOf(r.CombinationStrengths(), first)
	secondIndex := lo.IndexOf(r.CombinationStrengths(), second)
	return firstIndex < secondIndex
}

func (r Combination) Less(other Combination) bool {
	combinations := []Combination{r, other}

	if r.lessByType(r.Type(), other.Type()) {
		return true
	} else if r.Type() == other.Type() {
		if r.MainCard() < other.MainCard() {
			return true
		} else if r.MainCard() == other.MainCard() {
			if comparedBySecondary(r.Type()) {
				if int(*r.SecondaryCard()) < int(*other.SecondaryCard()) {
					return lessBySecondary(combinations)
				} else if int(*r.SecondaryCard()) == int(*other.SecondaryCard()) {
					if comparedByKickers(r.Type()) {
						return lessByKickers(combinations)
					} else {
						return false
					}
				}
			} else if comparedByKickers(r.Type()) {
				return lessByKickers(combinations)
			}
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
	if int(*combinations[0].SecondaryCard()) < int(*combinations[1].SecondaryCard()) {
		return true
	} else {
		return false
	}
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
		} else if int(face) > int(opposite) {
			return false
		}
	}
	return false
}


func NewDefaultCombination(cards []Card) (*Combination, error) {
	return NewCombination(cards, DefaultCombinationStrength, false)
}

func NewShortDeckCombination(cards []Card) (*Combination, error) {
	return NewCombination(cards, ShortDeckCombinationStrength, true)
}

func NewCombination(cards []Card, combinationStrengths []CombinationType, shortDeck bool) (*Combination, error) {
	sort.Sort(ByFace(cards))
	if len(cards) == validCardsLength {
		uniques := lo.Uniq(cards)
		if len(uniques) != validCardsLength {
			return nil, errors.New("cannot construct combination with not unique cards")
		}
		return &Combination{cards: cards, combinationStrengths: combinationStrengths, shortDeck: shortDeck}, nil
	} else {
		return nil, fmt.Errorf("cannot construct a combination you must pass slice of size: %d", validCardsLength)
	}
}


func CombinationsOf(cards[] Card, combinationStrengths []CombinationType, shortDeck bool) ([]Combination, error) {
	if len(cards) < validCardsLength {
		return nil, fmt.Errorf("Cannot construct combinations from {%d} cards, must be more than 5", len(cards))
	}

	combinatoricsCombinations := combin.Combinations(len(cards), validCardsLength)
	for i := range combinatoricsCombinations {
		sort.Ints(combinatoricsCombinations[i])
	}
	
	combinations := []Combination{}

	for _, v := range combinatoricsCombinations {
		toCards := lo.Map(v, func(value int, index int) Card {
			return cards[value]
		})
		
		combination, err := NewCombination(toCards, combinationStrengths, shortDeck)

		if err != nil {
			return nil, err
		}

		combinations = append(combinations, *combination)
	}

	sort.Sort(ByCombinationReversed(combinations))
	return combinations, nil
}

func StrongestCombinationOf(cards []Card, combinationStrengths []CombinationType, shortDeck bool) (*Combination, error) {
	combinations, err := CombinationsOf(cards, combinationStrengths, shortDeck)
	if err != nil {
		return nil, err
	}

	return &combinations[0], nil
}

