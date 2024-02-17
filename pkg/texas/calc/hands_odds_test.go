package calc

import (
	"testing"

	"github.com/anuarkaliyev23/goker/pkg/cards"
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

func TestHandsOdds(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		hands := generateHands(2)
		iterationCount := 1000

		config := HandOddsConfig {
			Hands: hands, 
			IterationsCount: iterationCount,
		}

		odds, err := HandsOdds(config)
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
			
			result, err := HandsOdds(config)
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
			
			result, err := HandsOdds(config)
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

			result, err := HandsOdds(config)
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

			odds, err := HandsOdds(config)
			require.Error(t, err)
			require.Nil(t, odds)
		})
	})
}
