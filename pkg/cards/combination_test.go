package cards

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCombinations(t *testing.T) {
	t.Run("Non-Unique slice of cards", func(t *testing.T) {
		cards := []Card {
			{face: Two, suit: Spades},
			{face: Two, suit: Diamonds},
			{face: Two, suit: Clubs},
			{face: Two, suit: Hearts},
			{face: Two, suit: Spades},
		}

		combination, err := NewCombination(cards)
		require.Nil(t, combination)
		require.NotNil(t, err)
	})

	t.Run("Less than 5 cards", func(t *testing.T) {
		cards := []Card {
			{face: Two, suit: Spades},
			{face: Two, suit: Diamonds},
			{face: Two, suit: Clubs},
			{face: Two, suit: Hearts},
		}

		combination, err := NewCombination(cards)
		require.Nil(t, combination)
		require.NotNil(t, err)
	})

	t.Run("Positive", func(t *testing.T) {
		cards := []Card {
			{face: Two, suit: Spades},
			{face: Three, suit: Diamonds},
			{face: Queen, suit: Clubs},
			{face: King, suit: Hearts},
			{face: Ten, suit: Spades},
		}

		combination, err := NewCombination(cards)
		require.NotNil(t, combination)
		require.NoError(t, err)
	})
}

func TestCombination_isStraiht(t *testing.T) {
	t.Run("A 2 3 4 5", func(t *testing.T) {
		cards := []Card {
			{face: Two, suit: Spades},
			{face: Ace, suit: Diamonds},
			{face: Three, suit: Clubs},
			{face: Five, suit: Hearts},
			{face: Four, suit: Spades},
		}

		combination, err := NewCombination(cards)
		require.NoError(t, err)
		require.True(t, combination.isStraight())
	})
	
	t.Run("A K Q J T", func(t *testing.T) {
		cards := []Card {
			{face: Jack, suit: Spades},
			{face: King, suit: Diamonds},
			{face: Ace, suit: Clubs},
			{face: Queen, suit: Hearts},
			{face: Ten, suit: Spades},
		}

		combination, err := NewCombination(cards)
		require.NoError(t, err)
		require.True(t, combination.isStraight())
	})

	t.Run("5 6 7 8 9", func(t *testing.T) {
		cards := []Card {
			{face: Nine, suit: Spades},
			{face: Six, suit: Diamonds},
			{face: Seven, suit: Clubs},
			{face: Eight, suit: Hearts},
			{face: Five, suit: Spades},
		}

		combination, err := NewCombination(cards)
		require.NoError(t, err)
		require.True(t, combination.isStraight())
	})

	t.Run("2 4 5 6 7", func(t *testing.T) {
		cards := []Card {
			{face: Two, suit: Spades},
			{face: Four, suit: Diamonds},
			{face: Five, suit: Clubs},
			{face: Six, suit: Hearts},
			{face: Seven, suit: Spades},
		}
		combination, err := NewCombination(cards)
		require.NoError(t, err)
		require.False(t, combination.isStraight())
	})
}

func TestCombination_isFlush(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		cards := []Card {
			{face: Two, suit: Spades},
			{face: Four, suit: Spades},
			{face: Five, suit: Spades},
			{face: Six, suit: Spades},
			{face: Seven, suit: Spades},
		}

		combination, err := NewCombination(cards)
		require.NoError(t, err)
		require.True(t, combination.isFlush())
	})

	
	t.Run("Negative", func(t *testing.T) {
		cards := []Card {
			{face: Two, suit: Diamonds},
			{face: Four, suit: Spades},
			{face: Five, suit: Spades},
			{face: Six, suit: Spades},
			{face: Seven, suit: Spades},
		}

		combination, err := NewCombination(cards)
		require.NoError(t, err)
		require.False(t, combination.isFlush())
	})
}
