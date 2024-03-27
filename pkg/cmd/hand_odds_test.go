package cmd

import (
	"testing"

	"github.com/anuarkaliyev23/goker/pkg/cards"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func ignoreError[T any](fn func() (T, error)) T {
	result, err := fn()
	if err != nil {
		panic(err)
	}
	return result
}

func Test_handOdds(t *testing.T) {
	t.Run("KK vs 77, i=10", func(t *testing.T) {
		board := ""
		hands := []string{"KsKd", "7s7d"}
		iterations := 10

		odds, err := handOdds(board, hands, iterations)
		require.NoError(t, err)
		require.Equal(t, 10, odds.Config.IterationsCount)
		require.Equal(t, 2, odds.NumberOfPlayers())

		firstHand, err := odds.PlayerHand(0)
		require.NoError(t, err)

		secondHand, err := odds.PlayerHand(1)
		require.NoError(t, err)

		require.Equal(t, 2, len(firstHand))
		require.Equal(t, 2, len(secondHand))
		
		require.True(t, lo.Contains(firstHand, ignoreError(func() (cards.Card, error) {
			card, err := cards.NewCard(cards.King, cards.Spades)
			return *card, err
		})))

		require.True(t, lo.Contains(firstHand, ignoreError(func() (cards.Card, error) {
			card, err := cards.NewCard(cards.King, cards.Diamonds)
			return *card, err
		})))

		require.True(t, lo.Contains(secondHand, ignoreError(func() (cards.Card, error) {
			card, err := cards.NewCard(cards.Seven, cards.Diamonds)
			return *card, err
		})))

		require.True(t, lo.Contains(secondHand, ignoreError(func() (cards.Card, error) {
			card, err := cards.NewCard(cards.Seven, cards.Spades)
			return *card, err
		})))

	})
}
