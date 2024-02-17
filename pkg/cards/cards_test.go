package cards

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


func allCards() []Card {
	cc := []Card{}
	for _, face := range Faces {
		for _, suit := range Suits {
			c := Card{
				face: face,
				suit: suit,
			}
			cc = append(cc, c)
		}
	}
	return cc
}

func TestNewCard(t *testing.T) {
	t.Run("valid card creation", func(t *testing.T) {
		for _, face := range Faces{
			for _, suit := range Suits {
				c, err := NewCard(face, suit)
				assert.NoError(t, err)
				assert.Equal(t, face, c.face)
				assert.Equal(t, suit, c.suit)
			}
		}
	})

	t.Run("invalid suit", func(t *testing.T) {
		_, err := NewCard(Ace, 9)
		assert.Error(t, err)
	})

	t.Run("Invalid Face", func(t *testing.T) {
		_, err := NewCard(18, Diamonds)
		assert.Error(t, err)
	})
}

func TestNewDeck(t *testing.T) {
	t.Run("Full Deck", func(t *testing.T) {
		deck := NewFullDeck()
		assert.Equal(t, 52, len(deck.left))
		assert.Equal(t, 0, len(deck.drawn))
	})

	t.Run("Valid, without validation", func(t *testing.T) {
		deck := NewDeckWithoutValidation(allCards())
		assert.Equal(t, 52, len(deck.left))
		assert.Equal(t, 0, len(deck.drawn))
	})

	t.Run("Not valid, without validation", func(t *testing.T) {
		deck := NewDeckWithoutValidation(allCards()[15:])
		assert.Equal(t, 52 - 15, len(deck.left))
		assert.Equal(t, 0, len(deck.drawn))
	})
	
	t.Run("Valid, With Validation", func(t *testing.T) {
		deck, err := NewDeck(allCards())
		assert.NoError(t, err)
		assert.Equal(t, 52, len(deck.left))
		assert.Equal(t, 0, len(deck.drawn))
	})

	t.Run("Invalid, fails on validation", func(t *testing.T) {
		deck, err := NewDeck(allCards()[15:])
		assert.Nil(t, deck)
		assert.Error(t, err)
	})
}

func TestDeck_Draw(t *testing.T) {
	t.Run("Error on empty deck", func(t *testing.T) {
		deck := Deck {
			left: []Card{},
			drawn: allCards(),
		}

		c, err := deck.Draw()
		assert.Nil(t, c)
		assert.Error(t, err)
	})

	t.Run("Positive", func(t *testing.T) {
		deck := Deck {
			left: allCards()[:1],
			drawn: allCards()[1:],
		}

		c, err := deck.Draw()
		assert.NoError(t, err)
		assert.NotNil(t, c)
		assert.Equal(t, 0, len(deck.left))
		assert.Equal(t, 52, len(deck.drawn))
		assert.True(t, lo.Contains(deck.drawn, *c))
	})
}

func TestDeck_MoveToDraw(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		drawn := Card{face: Queen, suit: Spades}
		deck := NewFullDeck()

		err := deck.MoveToDrawn(drawn)
		require.Equal(t, drawn, deck.drawn[0])
		require.False(t, lo.Contains(deck.left, drawn))
		require.NoError(t, err)
	})

	t.Run("negative", func(t *testing.T) {
		t.Run("card already drawn", func(t *testing.T) {
			deck := NewFullDeck()
			drawnCard, err := deck.Draw()
			require.NoError(t, err)

			err = deck.MoveToDrawn(*drawnCard)
			require.Error(t, err)
		})

		t.Run("card is not in left", func(t *testing.T) {
			deck := NewFullDeck()
			drawnCard, err := deck.Draw()
			require.NoError(t, err)

			deck.drawn = []Card{}

			err = deck.MoveToDrawn(*drawnCard)
			require.Error(t, err)
		})
	})
}
