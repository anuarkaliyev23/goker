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
	Board []cards.Card
}

func (r HandOddsIteration) StrongestCombination() (*cards.Combination, error) {
	if len(r.Combinations) == 0 {
		return nil, errors.New("Cannot determine winner in empty list of combinations")
	}
	maxCombination := r.Combinations[0]
	
	for i := 1; i < len(r.Combinations); i++ {
		if r.Combinations[i].More(maxCombination) {
			maxCombination = r.Combinations[i]
		}
	}
	return &maxCombination, nil
}

func (r HandOddsIteration) playersWithStrongestCombinations() ([]int, error) {
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

func (r HandOddsIteration) Winner() (int, error) {
	winners, err := r.playersWithStrongestCombinations()
	if err != nil {
		return -1, err
	}

	if len(winners) > 1 {
		return -1, nil 
	}

	return winners[0], nil
}

func (r HandOddsIteration) IsTie() (bool, error) {
	winner, err := r.Winner()
	if err != nil {
		return false, err
	}

	if winner != -1 {
		return false, nil
	}

	winners, err := r.playersWithStrongestCombinations()
	if err != nil {
		return false, err
	}

	if len(winners) == len(r.Combinations) {
		return true, nil
	}

	return false,nil
}

type HandOddsResult struct {
	Config HandOddsConfig
	Iterations []HandOddsIteration
}

func (r HandOddsResult) PlayerWins(index int) (int, error) {
	winners := lo.Map(r.Iterations, func(iteration HandOddsIteration, _ int) int {
		w, _ := iteration.Winner()
		return w
	})

	wins := lo.CountBy(winners, func(player int) bool {
		return index == player
	})
	
	return wins, nil
}

func (r HandOddsResult) AllPlayerWins() ([]int, error) {
	wins := []int{}
	for i := 0; i < r.NumberOfPlayers(); i++ {
		w, err := r.PlayerWins(i)
		if err != nil {
			return nil, err 
		}
		wins = append(wins, w)
	}

	return wins, nil
}

func (r HandOddsResult) WinRates() ([]float32, error) {
	wins, err := r.AllPlayerWins()
	if err != nil {
		return nil, err
	}

	winRates := lo.Map(wins, func(w int, _ int) float32 {
		return float32(w) / float32(r.Config.IterationsCount)
	})

	return winRates, nil
}

func (r HandOddsResult) TiePercentage() (float32, error) {
	ties, err := r.Ties()
	if err != nil {
		return 0.0, err
	}

	return float32(ties) / float32(r.Config.IterationsCount), nil
}

func (r HandOddsResult) WinningPlayer() (int, error) {
	allPlayerWins, err := r.AllPlayerWins()
	if err != nil {
		return 0, nil
	}
	maxWins := lo.Max(allPlayerWins)
	maxIndex := lo.IndexOf(allPlayerWins, maxWins)
	return maxIndex, nil
}


func (r HandOddsResult) Ties() (int, error) {
	toTies := lo.Map(r.Iterations, func(iteration HandOddsIteration, _ int) bool {
		tie, _ := iteration.IsTie()
		return tie
	})

	ties := lo.Count(toTies, true)
	return ties, nil
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

func collectExcludedCards(board []cards.Card, hands [][]cards.Card) []cards.Card {
	excludedCards := board

	lo.ForEach(hands, func(hand []cards.Card , _ int) {
		excludedCards = append(excludedCards, hand...)
	})

	return excludedCards
}

func excludeCards(deck cards.Deck, excludedCards []cards.Card) cards.Deck {
	lo.ForEach(excludedCards, func(card cards.Card, _ int) {
		err := deck.MoveToDrawn(card)
		if err != nil {
			// This should never happen
			panic(err)
		}
	})

	return deck
}

func validateIteration(hands [][]cards.Card, board []cards.Card) error {
	if len(hands) > ((cards.ValidDeckSize - CommunityCardsCount - BurnCardsCount) / HoleCardsCount) {
		return fmt.Errorf("Too many players {%d}", len(hands))
	}

	containedInvaludHands := lo.SomeBy(hands, func(hand []cards.Card) bool {
		return len(hand) != HoleCardsCount
	})

	if containedInvaludHands {
		return fmt.Errorf("Cannot construct iteration for hand of invalid size, should be {%d}", HoleCardsCount)
	}

	return nil
}

func drawCommunityCards(deck cards.Deck, board []cards.Card) (cards.Deck, []cards.Card) {
	cardsToDrawCount := CardsUsedInTexasHoldem - HoleCardsCount - len(board)
	drawnCards := []cards.Card{}
	for i := 0; i < cardsToDrawCount; i++ {
		drawnCard, err := deck.Draw()
		if err != nil {
			//This should never happen 
			panic(err)
		}
		drawnCards = append(drawnCards, *drawnCard)
	}
	return deck, drawnCards
}

func strongestHandCombination(hand []cards.Card, board []cards.Card, extraCommunityCards []cards.Card) cards.Combination {
	usedCards := []cards.Card{}
	usedCards = append(usedCards, hand...)
	usedCards = append(usedCards, board...)
	usedCards = append(usedCards, extraCommunityCards...)

	combination, err := cards.StrongestCombinationOf(usedCards)
	if err != nil {
		//This should never happen
		panic(err)
	}
	return *combination
}

func iterate(hands [][]cards.Card, board []cards.Card) (*HandOddsIteration, error) {
	deck := cards.NewFullDeck()
	deck.Shuffle()

	err := validateIteration(hands, board)
	if err != nil {
		return nil, err
	}
	
	excludedCards := collectExcludedCards(board, hands)
	deck = excludeCards(deck, excludedCards)
	deck, extraCommunityCards := drawCommunityCards(deck, board)
	
	combinations := lo.Map(hands, func(cs []cards.Card, _ int) cards.Combination {
		return strongestHandCombination(cs, board, extraCommunityCards)
	})

	iteration := HandOddsIteration {
		Combinations: combinations,
		Board: append(board, extraCommunityCards...),
	}

	return &iteration, nil
}
