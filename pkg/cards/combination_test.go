package cards

import (
	"testing"

	"github.com/stretchr/testify/require"
)



func TestCombination_Kickers(t *testing.T) {
	t.Run("High Card Kickers", func(t *testing.T) {
		combination := Combination {
			cards: []Card{
				{face: Two, suit: Spades},
				{face: Seven, suit: Diamonds},
				{face: Three, suit: Clubs},
				{face: King, suit: Clubs},
				{face: Ten, suit: Spades},
			},
		}

		kickers := combination.Kickers()

		require.Equal(t, kickers, []Card{
			{face: Ten, suit: Spades},
			{face: Seven, suit: Diamonds},
			{face: Three, suit: Clubs},
			{face: Two, suit: Spades},
		})
	})
	
	t.Run("Pair Kickers", func(t *testing.T) {
		combination := Combination {
			cards: []Card{
				{face: Two, suit: Spades},
				{face: Two, suit: Diamonds},
				{face: Three, suit: Clubs},
				{face: King, suit: Clubs},
				{face: Ten, suit: Spades},
			},
		}

		kickers := combination.Kickers()

		require.Equal(t, kickers, []Card{
			{face: King, suit: Clubs}, 
			{face: Ten, suit: Spades},
			{face: Three, suit: Clubs},
		})
	})

	t.Run("Two Pair Kickers", func(t *testing.T) {
		combination := Combination {
			cards: []Card{
				{face: Two, suit: Spades},
				{face: Two, suit: Diamonds},
				{face: Three, suit: Clubs},
				{face: King, suit: Clubs},
				{face: King, suit: Spades},
			},
		}

		kickers := combination.Kickers()

		require.Equal(t, kickers, []Card{
			{face: Three, suit: Clubs},
		})
	})
	
	t.Run("Three Of A Kind Kickers", func(t *testing.T) {
		combination := Combination {
			cards: []Card{
				{face: Two, suit: Spades},
				{face: Two, suit: Diamonds},
				{face: Two, suit: Clubs},
				{face: Queen, suit: Clubs},
				{face: King, suit: Spades},
			},
		}

		kickers := combination.Kickers()
		require.Equal(t, kickers, []Card{
			{face: King, suit: Spades},
			{face: Queen, suit: Clubs},
		})
	})

	t.Run("Four Of A Kind Kickers", func(t *testing.T) {
		combination := Combination {
			cards: []Card{
				{face: Two, suit: Spades},
				{face: Two, suit: Diamonds},
				{face: Two, suit: Clubs},
				{face: Two, suit: Hearts},
				{face: King, suit: Spades},
			},
		}

		kickers := combination.Kickers()
		require.Equal(t, kickers, []Card{
			{face: King, suit: Spades},
		})
	})

	t.Run("Flush Kickers", func(t *testing.T) {
		combination := Combination {
			cards: []Card{
				{face: Two, suit: Spades},
				{face: Three, suit: Spades},
				{face: Queen, suit: Spades},
				{face: Seven, suit: Spades},
				{face: King, suit: Spades},
			},
		}

		kickers := combination.Kickers()
		require.Equal(t, kickers, []Card{
			{face: Queen, suit: Spades},
			{face: Seven, suit: Spades},
			{face: Three, suit: Spades},
			{face: Two, suit: Spades},
		})
	})
}

func TestCombination_MainCombinationCard(t *testing.T) {
	t.Run("High Card", func(t *testing.T) {
		combination := Combination {
			cards: []Card{
				{face: Two, suit: Spades},
				{face: Seven, suit: Diamonds},
				{face: Three, suit: Clubs},
				{face: King, suit: Clubs},
				{face: Ten, suit: Spades},
			},
		}

		require.Equal(t, King, combination.MainCard())
	})

	t.Run("Pair", func(t *testing.T) {
		combination := Combination {
			cards: []Card{
				{face: Two, suit: Spades},
				{face: Seven, suit: Diamonds},
				{face: Three, suit: Clubs},
				{face: King, suit: Clubs},
				{face: King, suit: Spades},
			},
		}

		require.Equal(t, King, combination.MainCard())
	})

	t.Run("Two Pair", func(t *testing.T) {
		combination := Combination {
			cards: []Card{
				{face: Two, suit: Spades},
				{face: Two, suit: Diamonds},
				{face: Three, suit: Clubs},
				{face: King, suit: Clubs},
				{face: King, suit: Spades},
			},
		}

		require.Equal(t, King, combination.MainCard())
	})

	t.Run("Three Of A Kind", func(t *testing.T) {
		combination := Combination {
			cards: []Card{
				{face: Two, suit: Spades},
				{face: Two, suit: Diamonds},
				{face: Two, suit: Clubs},
				{face: Queen, suit: Clubs},
				{face: King, suit: Spades},
			},
		}

		require.Equal(t, Two, combination.MainCard())
	})

	t.Run("Straight", func(t *testing.T) {
		combination := Combination {
			cards: []Card{
				{face: Three, suit: Spades},
				{face: Four, suit: Diamonds},
				{face: Five, suit: Clubs},
				{face: Six, suit: Clubs},
				{face: Seven, suit: Spades},
			},
		}

		require.Equal(t, Seven, combination.MainCard())
	})

	t.Run("Flush", func(t *testing.T) {
		combination := Combination {
			cards: []Card{
				{face: Two, suit: Spades},
				{face: Three, suit: Spades},
				{face: Queen, suit: Spades},
				{face: Seven, suit: Spades},
				{face: King, suit: Spades},
			},
		}

		require.Equal(t, King, combination.MainCard())
	})

	t.Run("Full House", func(t *testing.T) {
		combination := Combination {
			cards: []Card{
				{face: Two, suit: Spades},
				{face: Two, suit: Diamonds},
				{face: Two, suit: Clubs},
				{face: King, suit: Clubs},
				{face: King, suit: Spades},
			},
		}

		require.Equal(t, Two, combination.MainCard())
	})

	t.Run("Four Of A Kind", func(t *testing.T) {
		combination := Combination {
			cards: []Card{
				{face: Two, suit: Spades},
				{face: Two, suit: Diamonds},
				{face: Two, suit: Clubs},
				{face: Two, suit: Hearts},
				{face: King, suit: Spades},
			},
		}

		require.Equal(t, Two, combination.MainCard())
	})

	t.Run("Straight Flush", func(t *testing.T) {
		combination := Combination {
			cards: []Card{
				{face: Three, suit: Spades},
				{face: Four, suit: Diamonds},
				{face: Five, suit: Clubs},
				{face: Six, suit: Clubs},
				{face: Seven, suit: Spades},
			},
		}
		
		require.Equal(t, Seven, combination.MainCard())

	})
}

func TestCombination_SecondaryCard(t *testing.T) {

	var highCard = Combination { 
		cards: []Card {
			{face: Two, suit: Spades},
			{face: King, suit: Diamonds},
			{face: Ace, suit: Hearts},
			{face: Ten, suit: Clubs},
			{face: Queen, suit: Spades},
		},
	}
			
	var	pair = Combination {
		cards: []Card {
			{face: Two, suit: Spades},
			{face: King, suit: Diamonds},
			{face: King, suit: Hearts},
			{face: Ten, suit: Clubs},
			{face: Queen, suit: Spades},
		},
	}
			


	var threeOfAKind = Combination {
		cards: []Card {
			{face: Two, suit: Spades},
			{face: King, suit: Diamonds},
			{face: King, suit: Hearts},
			{face: King, suit: Clubs},
			{face: Ten, suit: Spades},
		},
	}

	var straight = Combination {
		cards: []Card {
			{face: Two, suit: Spades},
			{face: Three, suit: Diamonds},
			{face: Four, suit: Hearts},
			{face: Five, suit: Clubs},
			{face: Six, suit: Spades},
		},
	}

	var flush = Combination {
		cards: []Card {
			{face: Two, suit: Spades},
			{face: Seven, suit: Spades},
			{face: Four, suit: Spades},
			{face: Five, suit: Spades},
			{face: Six, suit: Spades},
		},
	}

	var fourOfAKind = Combination {
		cards: []Card {
			{face: Two, suit: Spades},
			{face: King, suit: Diamonds},
			{face: King, suit: Hearts},
			{face: King, suit: Clubs},
			{face: King, suit: Clubs},
		},
	}

	var straightFlush = Combination {
		cards: []Card {
			{face: Two, suit: Spades},
			{face: Three, suit: Spades},
			{face: Four, suit: Spades},
			{face: Five, suit: Spades},
			{face: Six, suit: Spades},
		},
	}

	t.Run("Empty Secondaries", func(t *testing.T) {
		require.Nil(t, highCard.SecondaryCard())
		require.Nil(t, pair.SecondaryCard())
		require.Nil(t, threeOfAKind.SecondaryCard())
		require.Nil(t, straight.SecondaryCard())
		require.Nil(t, flush.SecondaryCard())
		require.Nil(t, fourOfAKind.SecondaryCard())
		require.Nil(t, straightFlush.SecondaryCard())
	})

	t.Run("Full House", func(t *testing.T) {
		fullHouse := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: King, suit: Diamonds},
				{face: King, suit: Hearts},
				{face: King, suit: Clubs},
				{face: Two, suit: Clubs},
			},
		}

		require.Equal(t, Two, *fullHouse.SecondaryCard())
	
	})
	
	t.Run("Two Pair", func(t *testing.T) {
		twoPair := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: King, suit: Diamonds},
				{face: King, suit: Hearts},
				{face: Ten, suit: Clubs},
				{face: Ten, suit: Spades},
			},
		}

		require.Equal(t, Ten, *twoPair.SecondaryCard())
	})

}

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

func TestCombination_isFourOAKind(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		cards := []Card {
			{face: Two, suit: Spades},
			{face: Two, suit: Diamonds},
			{face: Two, suit: Hearts},
			{face: Two, suit: Clubs},
			{face: Queen, suit: Spades},
		}

		combination, err := NewCombination(cards)
		require.NoError(t, err)
		require.True(t, combination.isFourOfAKind())
	})

	
	t.Run("Negative", func(t *testing.T) {
		cards := []Card {
			{face: Two, suit: Spades},
			{face: Two, suit: Diamonds},
			{face: Two, suit: Hearts},
			{face: Three, suit: Clubs},
			{face: Queen, suit: Spades},
		}

		combination, err := NewCombination(cards)
		require.NoError(t, err)
		require.False(t, combination.isFourOfAKind())
	})
}

func TestCombination_isThreeOAKind(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		cards := []Card {
			{face: Two, suit: Spades},
			{face: Two, suit: Diamonds},
			{face: Two, suit: Hearts},
			{face: Five, suit: Clubs},
			{face: Queen, suit: Spades},
		}

		combination, err := NewCombination(cards)
		require.NoError(t, err)
		require.True(t, combination.isThreeOfAKind())
	})

	
	t.Run("Negative", func(t *testing.T) {
		cards := []Card {
			{face: Two, suit: Spades},
			{face: Two, suit: Diamonds},
			{face: Five, suit: Hearts},
			{face: Three, suit: Clubs},
			{face: Queen, suit: Spades},
		}

		combination, err := NewCombination(cards)
		require.NoError(t, err)
		require.False(t, combination.isThreeOfAKind())
	})
}

func TestCombination_isTwoPair(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		cards := []Card {
			{face: Two, suit: Spades},
			{face: Two, suit: Diamonds},
			{face: Five, suit: Hearts},
			{face: Five, suit: Clubs},
			{face: Queen, suit: Spades},
		}

		combination, err := NewCombination(cards)
		require.NoError(t, err)
		require.True(t, combination.isTwoPair())
	})

	
	t.Run("Negative", func(t *testing.T) {
		cards := []Card {
			{face: Two, suit: Spades},
			{face: Two, suit: Diamonds},
			{face: Five, suit: Hearts},
			{face: Three, suit: Clubs},
			{face: Queen, suit: Spades},
		}

		combination, err := NewCombination(cards)
		require.NoError(t, err)
		require.False(t, combination.isTwoPair())
	})
}

func TestCombination_isPair(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		cards := []Card {
			{face: Two, suit: Spades},
			{face: Two, suit: Diamonds},
			{face: Five, suit: Hearts},
			{face: Seven, suit: Clubs},
			{face: Queen, suit: Spades},
		}

		combination, err := NewCombination(cards)
		require.NoError(t, err)
		require.True(t, combination.isPair())
	})

	
	t.Run("Negative", func(t *testing.T) {
		cards := []Card {
			{face: Two, suit: Spades},
			{face: King, suit: Diamonds},
			{face: Five, suit: Hearts},
			{face: Three, suit: Clubs},
			{face: Queen, suit: Spades},
		}

		combination, err := NewCombination(cards)
		require.NoError(t, err)
		require.False(t, combination.isPair())
	})
}

func TestCombination_isFullHouse(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		cards := []Card {
			{face: Two, suit: Spades},
			{face: Two, suit: Diamonds},
			{face: Five, suit: Hearts},
			{face: Five, suit: Clubs},
			{face: Five, suit: Spades},
		}

		combination, err := NewCombination(cards)
		require.NoError(t, err)
		require.True(t, combination.isFullHouse())
	})

	
	t.Run("Negative", func(t *testing.T) {
		cards := []Card {
			{face: Two, suit: Spades},
			{face: Three, suit: Diamonds},
			{face: Three, suit: Hearts},
			{face: Three, suit: Clubs},
			{face: Queen, suit: Spades},
		}

		combination, err := NewCombination(cards)
		require.NoError(t, err)
		require.False(t, combination.isFullHouse())
	})
}

func TestCombination_isHighCard(t *testing.T) {
	t.Run("Negative", func(t *testing.T) {
		cards := []Card {
			{face: Two, suit: Spades},
			{face: Two, suit: Diamonds},
			{face: Five, suit: Hearts},
			{face: Five, suit: Clubs},
			{face: Five, suit: Spades},
		}

		combination, err := NewCombination(cards)
		require.NoError(t, err)
		require.False(t, combination.isHighCard())
	})

	
	t.Run("Positive", func(t *testing.T) {
		cards := []Card {
			{face: Two, suit: Spades},
			{face: King, suit: Diamonds},
			{face: Ace, suit: Hearts},
			{face: Ten, suit: Clubs},
			{face: Queen, suit: Spades},
		}

		combination, err := NewCombination(cards)
		require.NoError(t, err)
		require.True(t, combination.isHighCard())
	})
}

func TestCombination_Type(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		highCard := Combination { 
			cards: []Card {
				{face: Two, suit: Spades},
				{face: King, suit: Diamonds},
				{face: Ace, suit: Hearts},
				{face: Ten, suit: Clubs},
				{face: Queen, suit: Spades},
			},
		}
		
		pair := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: King, suit: Diamonds},
				{face: King, suit: Hearts},
				{face: Ten, suit: Clubs},
				{face: Queen, suit: Spades},
			},
		}
		
		twoPair := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: King, suit: Diamonds},
				{face: King, suit: Hearts},
				{face: Ten, suit: Clubs},
				{face: Ten, suit: Spades},
			},
		}

		threeOfAKind := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: King, suit: Diamonds},
				{face: King, suit: Hearts},
				{face: King, suit: Clubs},
				{face: Ten, suit: Spades},
			},
		}

		straight := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: Three, suit: Diamonds},
				{face: Four, suit: Hearts},
				{face: Five, suit: Clubs},
				{face: Six, suit: Spades},
			},
		}

		flush := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: Seven, suit: Spades},
				{face: Four, suit: Spades},
				{face: Five, suit: Spades},
				{face: Six, suit: Spades},
			},
		}

		fullHouse := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: King, suit: Diamonds},
				{face: King, suit: Hearts},
				{face: King, suit: Clubs},
				{face: Two, suit: Clubs},
			},
		}

		fourOfAKind := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: King, suit: Diamonds},
				{face: King, suit: Hearts},
				{face: King, suit: Clubs},
				{face: King, suit: Clubs},
			},
		}

		straightFlush := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: Three, suit: Spades},
				{face: Four, suit: Spades},
				{face: Five, suit: Spades},
				{face: Six, suit: Spades},
			},
		}

		require.Equal(t, highCard.Type(), HighCard)
		require.Equal(t, pair.Type(), Pair)
		require.Equal(t, twoPair.Type(), TwoPair)
		require.Equal(t, threeOfAKind.Type(), ThreeOfAKind)
		require.Equal(t, straight.Type(), Straight)
		require.Equal(t, flush.Type(), Flush)
		require.Equal(t, fullHouse.Type(), FullHouse)
		require.Equal(t, fourOfAKind.Type(), FourOfAKind)
		require.Equal(t, straightFlush.Type(), StraghtFlush)
	})
}

func TestCombination_Less(t *testing.T) {
	t.Run("combination order", func (t *testing.T) {
		highCard := Combination { 
			cards: []Card {
				{face: Two, suit: Spades},
				{face: King, suit: Diamonds},
				{face: Ace, suit: Hearts},
				{face: Ten, suit: Clubs},
				{face: Queen, suit: Spades},
			},
		}
		
		pair := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: King, suit: Diamonds},
				{face: King, suit: Hearts},
				{face: Ten, suit: Clubs},
				{face: Queen, suit: Spades},
			},
		}
		
		twoPair := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: King, suit: Diamonds},
				{face: King, suit: Hearts},
				{face: Ten, suit: Clubs},
				{face: Ten, suit: Spades},
			},
		}

		threeOfAKind := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: King, suit: Diamonds},
				{face: King, suit: Hearts},
				{face: King, suit: Clubs},
				{face: Ten, suit: Spades},
			},
		}

		straight := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: Three, suit: Diamonds},
				{face: Four, suit: Hearts},
				{face: Five, suit: Clubs},
				{face: Six, suit: Spades},
			},
		}

		flush := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: Seven, suit: Spades},
				{face: Four, suit: Spades},
				{face: Five, suit: Spades},
				{face: Six, suit: Spades},
			},
		}

		fullHouse := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: King, suit: Diamonds},
				{face: King, suit: Hearts},
				{face: King, suit: Clubs},
				{face: Two, suit: Clubs},
			},
		}

		fourOfAKind := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: King, suit: Diamonds},
				{face: King, suit: Hearts},
				{face: King, suit: Clubs},
				{face: King, suit: Clubs},
			},
		}

		straightFlush := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: Three, suit: Spades},
				{face: Four, suit: Spades},
				{face: Five, suit: Spades},
				{face: Six, suit: Spades},
			},
		}

		require.False(t, highCard.Less(highCard))
		require.False(t, pair.Less(pair))
		require.False(t, twoPair.Less(twoPair))
		require.False(t, threeOfAKind.Less(threeOfAKind))
		require.False(t, straight.Less(straight))
		require.False(t, flush.Less(flush))
		require.False(t, fullHouse.Less(fullHouse))
		require.False(t, fourOfAKind.Less(fourOfAKind))
		require.False(t, straightFlush.Less(straightFlush))

		require.True(t, highCard.Less(pair))
		require.True(t, highCard.Less(twoPair))
		require.True(t, highCard.Less(threeOfAKind))
		require.True(t, highCard.Less(straight))
		require.True(t, highCard.Less(flush))
		require.True(t, highCard.Less(fullHouse))
		require.True(t, highCard.Less(fourOfAKind))
		require.True(t, highCard.Less(straightFlush))

		require.False(t, pair.Less(highCard))
		require.True(t, pair.Less(threeOfAKind))
		require.True(t, pair.Less(straight))
		require.True(t, pair.Less(flush))
		require.True(t, pair.Less(fullHouse))
		require.True(t, pair.Less(fourOfAKind))
		require.True(t, pair.Less(straightFlush))

		require.False(t, twoPair.Less(highCard))
		require.False(t, twoPair.Less(pair))
		require.True(t, twoPair.Less(threeOfAKind))
		require.True(t, twoPair.Less(straight))
		require.True(t, twoPair.Less(flush))
		require.True(t, twoPair.Less(fullHouse))
		require.True(t, twoPair.Less(fourOfAKind))
		require.True(t, twoPair.Less(straightFlush))


		require.False(t, threeOfAKind.Less(highCard))
		require.False(t, threeOfAKind.Less(pair))
		require.False(t, threeOfAKind.Less(twoPair))
		require.True(t, threeOfAKind.Less(straight))
		require.True(t, threeOfAKind.Less(flush))
		require.True(t, threeOfAKind.Less(fullHouse))
		require.True(t, threeOfAKind.Less(fourOfAKind))
		require.True(t, threeOfAKind.Less(straightFlush))

		require.False(t, straight.Less(highCard))
		require.False(t, straight.Less(pair))
		require.False(t, straight.Less(twoPair))
		require.False(t, straight.Less(threeOfAKind))
		require.True(t, straight.Less(flush))
		require.True(t, straight.Less(fullHouse))
		require.True(t, straight.Less(fourOfAKind))
		require.True(t, straight.Less(straightFlush))

		require.False(t, flush.Less(highCard))
		require.False(t, flush.Less(pair))
		require.False(t, flush.Less(twoPair))
		require.False(t, flush.Less(threeOfAKind))
		require.False(t, flush.Less(straight))
		require.True(t, flush.Less(fullHouse))
		require.True(t, flush.Less(fourOfAKind))
		require.True(t, flush.Less(straightFlush))

		require.False(t, fullHouse.Less(highCard))
		require.False(t, fullHouse.Less(pair))
		require.False(t, fullHouse.Less(twoPair))
		require.False(t, fullHouse.Less(threeOfAKind))
		require.False(t, fullHouse.Less(straight))
		require.False(t, fullHouse.Less(flush))
		require.True(t, fullHouse.Less(fourOfAKind))
		require.True(t, fullHouse.Less(straightFlush))

		require.False(t, fourOfAKind.Less(highCard))
		require.False(t, fourOfAKind.Less(pair))
		require.False(t, fourOfAKind.Less(twoPair))
		require.False(t, fourOfAKind.Less(threeOfAKind))
		require.False(t, fourOfAKind.Less(straight))
		require.False(t, fourOfAKind.Less(flush))
		require.False(t, fourOfAKind.Less(fullHouse))
		require.True(t, fourOfAKind.Less(straightFlush))
		
		require.False(t, straightFlush.Less(highCard))
		require.False(t, straightFlush.Less(pair))
		require.False(t, straightFlush.Less(twoPair))
		require.False(t, straightFlush.Less(threeOfAKind))
		require.False(t, straightFlush.Less(straight))
		require.False(t, straightFlush.Less(flush))
		require.False(t, straightFlush.Less(fullHouse))
		require.False(t, straightFlush.Less(fourOfAKind))
	})

	t.Run("High Card", func(t *testing.T) {
		highCardSeven:= Combination { 
			cards: []Card {
				{face: Two, suit: Spades},
				{face: Three, suit: Diamonds},
				{face: Seven, suit: Hearts},
				{face: Six, suit: Clubs},
				{face: Five, suit: Spades},
			},
		}

		highCardEight := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: Three, suit: Diamonds},
				{face: Eight, suit: Hearts},
				{face: Six, suit: Clubs},
				{face: Five, suit: Spades},
			},
		}


		highCardNine := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: Three, suit: Diamonds},
				{face: Nine, suit: Hearts},
				{face: Six, suit: Clubs},
				{face: Five, suit: Spades},
			},
		}

		highCardTen := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: Three, suit: Diamonds},
				{face: Ten, suit: Hearts},
				{face: Six, suit: Clubs},
				{face: Five, suit: Spades},
			},
		}

		highCardJack:= Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: Three, suit: Diamonds},
				{face: Jack, suit: Hearts},
				{face: Six, suit: Clubs},
				{face: Five, suit: Spades},
			},
		}

		highCardQueen := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: Three, suit: Diamonds},
				{face: Queen, suit: Hearts},
				{face: Six, suit: Clubs},
				{face: Five, suit: Spades},
			},
		}

		highCardKing := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: Three, suit: Diamonds},
				{face: King, suit: Hearts},
				{face: Six, suit: Clubs},
				{face: Five, suit: Spades},
			},
		}

		highCardAce := Combination {
			cards: []Card {
				{face: Two, suit: Spades},
				{face: Three, suit: Diamonds},
				{face: Ace, suit: Hearts},
				{face: Six, suit: Clubs},
				{face: Five, suit: Spades},
			},
		}

		require.True(t, highCardSeven.Less(highCardEight))
		require.True(t, highCardSeven.Less(highCardNine))
		require.True(t, highCardSeven.Less(highCardTen))
		require.True(t, highCardSeven.Less(highCardJack))
		require.True(t, highCardSeven.Less(highCardQueen))
		require.True(t, highCardSeven.Less(highCardKing))
		require.True(t, highCardSeven.Less(highCardAce))
		
		require.False(t, highCardEight.Less(highCardSeven))
		require.True(t, highCardEight.Less(highCardNine))
		require.True(t, highCardEight.Less(highCardTen))
		require.True(t, highCardEight.Less(highCardJack))
		require.True(t, highCardEight.Less(highCardQueen))
		require.True(t, highCardEight.Less(highCardKing))
		require.True(t, highCardEight.Less(highCardAce))
	})
}

func TestLessByKickers(t *testing.T) {
	t.Run("High Card", func(t *testing.T) {
		more := Combination { 
			cards: []Card {
				{face: Two, suit: Spades},
				{face: King, suit: Diamonds},
				{face: Ace, suit: Hearts},
				{face: Ten, suit: Clubs},
				{face: Queen, suit: Spades},
			},
		}

		less := Combination { 
			cards: []Card {
				{face: Two, suit: Spades},
				{face: Seven, suit: Diamonds},
				{face: Eight, suit: Hearts},
				{face: Ten, suit: Clubs},
				{face: Jack, suit: Spades},
			},
		}

		moreLess := lessByKickers([]Combination {more, less})
		lessMore := lessByKickers([]Combination {less, more})
		require.Equal(t, false , moreLess)
		require.Equal(t, true, lessMore)
	})
}
