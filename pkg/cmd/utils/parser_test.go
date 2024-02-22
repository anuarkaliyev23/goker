package cmd

import (
	"testing"

	"github.com/anuarkaliyev23/goker/pkg/cards"
	"github.com/stretchr/testify/require"
)

func TestParseCard(t *testing.T) {
	t.Run("invalid cased", func(t *testing.T) {
		card, err := ParseCard("8J")
		require.Nil(t, card)
		require.Error(t, err)
	})

	t.Run("valid cases", func(t *testing.T) {
		t.Run("Ace", func(t *testing.T) {
			t.Run("Clubs", func(t *testing.T) {
				card, err := ParseCard("AC")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.Ace, cards.Clubs)
				require.Equal(t, expected, card)
			})

			t.Run("Diamonds", func(t *testing.T) {
				card, err := ParseCard("AD")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.Ace, cards.Diamonds)
				require.Equal(t, expected, card)
			})


			t.Run("Spades", func(t *testing.T) {
				card, err := ParseCard("AS")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.Ace, cards.Spades)
				require.Equal(t, expected, card)
			})

			t.Run("Hearts", func(t *testing.T) {
				card, err := ParseCard("AH")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.Ace, cards.Hearts)
				require.Equal(t, expected, card)
			})
		})

		t.Run("King", func(t *testing.T) {
			t.Run("Clubs", func(t *testing.T) {
				card, err := ParseCard("KC")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.King, cards.Clubs)
				require.Equal(t, expected, card)
			})

			t.Run("Diamonds", func(t *testing.T) {
				card, err := ParseCard("KD")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.King, cards.Diamonds)
				require.Equal(t, expected, card)
			})


			t.Run("Spades", func(t *testing.T) {
				card, err := ParseCard("KS")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.King, cards.Spades)
				require.Equal(t, expected, card)
			})

			t.Run("Hearts", func(t *testing.T) {
				card, err := ParseCard("KH")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.King, cards.Hearts)
				require.Equal(t, expected, card)
			})
		})

		t.Run("Queen", func(t *testing.T) {
			t.Run("Clubs", func(t *testing.T) {
				card, err := ParseCard("QC")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.Queen, cards.Clubs)
				require.Equal(t, expected, card)
			})

			t.Run("Diamonds", func(t *testing.T) {
				card, err := ParseCard("QD")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.Queen, cards.Diamonds)
				require.Equal(t, expected, card)
			})


			t.Run("Spades", func(t *testing.T) {
				card, err := ParseCard("QS")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.Queen, cards.Spades)
				require.Equal(t, expected, card)
			})

			t.Run("Hearts", func(t *testing.T) {
				card, err := ParseCard("QH")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.Queen, cards.Hearts)
				require.Equal(t, expected, card)
			})
		})
		
		t.Run("Jack", func(t *testing.T) {
			t.Run("Clubs", func(t *testing.T) {
				card, err := ParseCard("JC")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.Jack, cards.Clubs)
				require.Equal(t, expected, card)
			})

			t.Run("Diamonds", func(t *testing.T) {
				card, err := ParseCard("JD")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.Jack, cards.Diamonds)
				require.Equal(t, expected, card)
			})


			t.Run("Spades", func(t *testing.T) {
				card, err := ParseCard("JS")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.Jack, cards.Spades)
				require.Equal(t, expected, card)
			})

			t.Run("Hearts", func(t *testing.T) {
				card, err := ParseCard("JH")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.Jack, cards.Hearts)
				require.Equal(t, expected, card)
			})
		})


		t.Run("Ten", func(t *testing.T) {
			t.Run("Clubs", func(t *testing.T) {
				card, err := ParseCard("TC")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.Ten, cards.Clubs)
				require.Equal(t, expected, card)
			})

			t.Run("Diamonds", func(t *testing.T) {
				card, err := ParseCard("TD")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.Ten, cards.Diamonds)
				require.Equal(t, expected, card)
			})


			t.Run("Spades", func(t *testing.T) {
				card, err := ParseCard("TS")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.Ten, cards.Spades)
				require.Equal(t, expected, card)
			})

			t.Run("Hearts", func(t *testing.T) {
				card, err := ParseCard("TH")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.Ten, cards.Hearts)
				require.Equal(t, expected, card)
			})
		})

		t.Run("Nine", func(t *testing.T) {
			t.Run("Clubs", func(t *testing.T) {
				card, err := ParseCard("9C")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.Nine, cards.Clubs)
				require.Equal(t, expected, card)
			})

			t.Run("Diamonds", func(t *testing.T) {
				card, err := ParseCard("9D")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.Nine, cards.Diamonds)
				require.Equal(t, expected, card)
			})


			t.Run("Spades", func(t *testing.T) {
				card, err := ParseCard("9S")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.Nine, cards.Spades)
				require.Equal(t, expected, card)
			})

			t.Run("Hearts", func(t *testing.T) {
				card, err := ParseCard("9H")
				require.NoError(t, err)

				expected, _ := cards.NewCard(cards.Nine, cards.Hearts)
				require.Equal(t, expected, card)
			})
		})
	})
}


func TestParseCards(t *testing.T) {
	t.Run("invalid cases", func(t *testing.T) {
		t.Run("Ac7h8", func(t *testing.T) {
			representation := "Ac7h8"
			cards, err := ParseCards(representation)
			require.Error(t, err)
			require.Nil(t, cards)
		})
	})

	t.Run("valid cases", func(t *testing.T) {
		t.Run("AcKhQh", func(t *testing.T) {
			representation := "AcKhQh"
			parsed, err := ParseCards(representation)

			require.NotNil(t, parsed)
			require.NoError(t, err)
			require.Equal(t, 3, len(parsed))

			require.Equal(t, parsed[0].Face(), cards.Ace)
			require.Equal(t, parsed[0].Suit(), cards.Clubs)

			require.Equal(t, parsed[1].Face(), cards.King)
			require.Equal(t, parsed[1].Suit(), cards.Hearts)

			require.Equal(t, parsed[2].Face(), cards.Queen)
			require.Equal(t, parsed[2].Suit(), cards.Hearts)
		})
	})
}
