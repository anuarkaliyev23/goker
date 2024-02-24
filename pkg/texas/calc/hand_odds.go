package calc

import (
	"errors"
	"fmt"

	"github.com/anuarkaliyev23/goker/pkg/cards"
	"github.com/samber/lo"
)

const CardsUsedInTexasHoldem = 7
const BurnCardsCount = 3
const CommunityCardsCount = 5
const HoleCardsCount = 2

type HandOddsConfig struct {
	Hands [][]cards.Card
	Board []cards.Card
	IterationsCount int
}

type HandOddsIteration struct {
	Combinations []cards.Combination
}

func (r HandOddsIteration) StrongestCombination() (*cards.Combination, error) {
	if len(r.Combinations) == 0 {
		return nil, errors.New("Cannot determine winner in empty list of combinations")
	}
	maxCombination := r.Combinations[0]
	
	for i := 1; i < len(r.Combinations); i++ {
		if r.Combinations[i].More(maxCombination) {
			maxCombination = r.Combinations[0]
		}
	}
	return &maxCombination, nil
}

func (r HandOddsIteration) Winners() ([]int, error) {
	winner, err := r.StrongestCombination()
	if err != nil {
		return nil, err
	}

	winners := []int{}
	
	for i, v := range r.Combinations {
		if v.Tie(*winner) {
			winners = append(winners, i)
		}
	}

	return winners, nil
}

type HandOddsResult struct {
	Config HandOddsConfig
	Iterations []HandOddsIteration
}

func (r HandOddsResult) PlayerWins(index int) (*int, error) {
	winners := lo.Map(r.Iterations, func(iteration HandOddsIteration, _ int) []int {
		w, err := iteration.Winners()
		if err != nil {
			return nil
		}
		return w
	})

	for i, v := range winners {
		if v == nil {
			return nil, fmt.Errorf("Cannot decide winner for iteration {%d}, error in calculations", i)
		}
	}

	wins := lo.CountBy(winners, func(slice []int) bool {
		return lo.Contains(slice, index)
	})
	
	return &wins, nil
}

func (r HandOddsResult) NumberOfPlayers() int {
	return len(r.Config.Hands)
}

func (r HandOddsResult) PlayerHand(index int) ([]cards.Card, error) {
	if index > r.NumberOfPlayers() {
		return nil, fmt.Errorf("Player {%d} is out of range, number of players: {%d}", index, r.NumberOfPlayers())
	}

	return r.Config.Hands[index], nil
}

func (r HandOddsResult) PlayerCombinations(index int) ([]cards.Combination, error) {
	if index > r.NumberOfPlayers() {
		return nil, fmt.Errorf("Player {%d} is out of range, number of players: {%d}", index, r.NumberOfPlayers())
	}
	
	combinations := lo.Map(r.Iterations, func(iteration HandOddsIteration, _ int) cards.Combination {
		return iteration.Combinations[index]
	})
	
	return combinations, nil
}

func HandOdds(config HandOddsConfig) (*HandOddsResult, error) {
	if config.IterationsCount <= 0 {
		return nil, fmt.Errorf("Cannot simulate hands odds for non-positive or zero iterations, was given {%d}", config.IterationsCount)
	}
	
	iterations := []HandOddsIteration{}
	
	for i := 0; i < config.IterationsCount; i++ {
		iteration, err := iterate(config.Hands, config.Board)
		if err != nil {
			return nil, err 
		}
		
		iterations = append(iterations, *iteration)
	}
	
	return &HandOddsResult{
		Config: config,
		Iterations: iterations,
	}, nil
}

func iterate(hands [][]cards.Card, board []cards.Card) (*HandOddsIteration, error) {
	deck := cards.NewFullDeck()
	deck.Shuffle()

	if len(hands) > ((cards.ValidDeckSize - CommunityCardsCount - BurnCardsCount) / HoleCardsCount) {
		return nil, fmt.Errorf("Too many players {%d}", len(hands))
	}

	containedInvaludHands := lo.SomeBy(hands, func(hand []cards.Card) bool {
		return len(hand) != HoleCardsCount
	})

	if containedInvaludHands {
		return nil, fmt.Errorf("Cannot construct iteration for hand of invalid size, should be {%d}", HoleCardsCount)
	}

	excludedCards := board

	lo.ForEach(hands, func(hand []cards.Card , _ int) {
		lo.ForEach(hand, func(card cards.Card, _ int) {
			excludedCards = append(excludedCards, card)
		})
	})

	lo.ForEach(excludedCards, func(card cards.Card, _ int) {
		err := deck.MoveToDrawn(card)
		if err != nil {
			// This should never happen
			panic(err)
		}
	})

	combinations := lo.Map(hands, func(cs []cards.Card, _ int) cards.Combination  {
		usedCards := cs
		usedCards = append(usedCards, board...)
		
		leftCards := CardsUsedInTexasHoldem - len(hands)
		for i := 0; i < leftCards; i++ {
			drawnCard, err := deck.Draw()
			if err != nil {
				//This should never happen
				panic(err)
			}

			usedCards = append(usedCards, *drawnCard)
		}

		combination, err := cards.StrongestCombinationOf(usedCards)
		if err != nil {
			//This should never happen
			panic(err)
		}
		
		return *combination
	})

	iteration := HandOddsIteration {
		Combinations: combinations,
	}

	return &iteration, nil
}
