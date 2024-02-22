package calc

import (
	"testing"

	"github.com/anuarkaliyev23/goker/pkg/cards"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func generateHands(playersCount int) [][]cards.Card {
	deck := cards.NewFullDeck()
	hands := [][]cards.Card{}

	for i := 0; i < playersCount; i++ {
		hand := []cards.Card{}
		for j := 0; j < 2; j++ {
			card, err := deck.Draw()
			if err != nil {
				//This should never happen
				panic(err)
			}

			hand = append(hand, *card)
		}
		hands = append(hands, hand)
	}

	return hands
}

func generateHandOdds(playersCount int, iterations int) HandOddsResult {
	hands := generateHands(playersCount)
	config := HandOddsConfig {
		Hands: hands,
		IterationsCount: iterations,
	}

	odds, err := HandOdds(config)
	if err != nil {
		panic(err)
	}

	return *odds
}

func TestHandsResult_PlayerCombinations(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		t.Run("2 odds, no board", func(t *testing.T) {
			odds := generateHandOdds(2, 1000)
			playerHand, err := odds.PlayerHand(0)
			require.NoError(t, err)
			playerCombinations, err := odds.PlayerCombinations(0)
			require.NoError(t, err)

			contains := lo.Map(playerCombinations, func(c cards.Combination, _ int) bool {
				return lo.Some(c.AllCards(), playerHand)
			})
	
			present := lo.Contains(contains, true)
			require.True(t, present)
		})
	})

	t.Run("negative", func(t *testing.T) {
		hands := generateHands(2)
		iterationCount := 1000

		config := HandOddsConfig{
			Hands: hands,
			IterationsCount: iterationCount,
		}

		odds, err := HandOdds(config)
		require.NoError(t, err)
		require.NotNil(t, odds)

		
	})
}

func TestHandsOdds(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		hands := generateHands(2)
		iterationCount := 1000

		config := HandOddsConfig {
			Hands: hands, 
			IterationsCount: iterationCount,
		}

		odds, err := HandOdds(config)
		require.NoError(t, err)
		require.Equal(t, len(odds.Iterations), iterationCount)
	})

	t.Run("negative", func(t *testing.T) {
		t.Run("negative IterationCount", func(t *testing.T) {
			hands := generateHands(2)
			iterationCount := -1

			config := HandOddsConfig{
				Hands: hands,
				IterationsCount: iterationCount,
			}
			
			result, err := HandOdds(config)
			require.Nil(t, result)
			require.Error(t, err)
		})

		t.Run("zero IterationCount", func(t *testing.T) {
			hands := generateHands(2)
			iterationCount := 0

			config := HandOddsConfig{
				Hands: hands,
				IterationsCount: iterationCount,
			}
			
			result, err := HandOdds(config)
			require.Nil(t, result)
			require.Error(t, err)
		})

		t.Run("inconsistent hand cards count", func(t *testing.T) {
			hands := generateHands(1)
			secondsHand := []cards.Card{}

			secondsHandSingleCard, err := cards.NewCard(cards.Seven, cards.Spades)
			require.NoError(t, err)
			secondsHand = append(secondsHand, *secondsHandSingleCard)

			hands = append(hands, secondsHand)

			config := HandOddsConfig{
				Hands: hands,
				IterationsCount: 1,
			}

			result, err := HandOdds(config)
			require.Error(t, err)
			require.Nil(t, result)
		})

		t.Run("too many players", func(t *testing.T) {
			hands := generateHands(23)
			iterationCount := 1000

			config := HandOddsConfig {
				Hands: hands, 
				IterationsCount: iterationCount,
			}

			odds, err := HandOdds(config)
			require.Error(t, err)
			require.Nil(t, odds)
		})
	})
}
