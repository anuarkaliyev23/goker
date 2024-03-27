package calc

import (
	"testing"

	"github.com/anuarkaliyev23/goker/pkg/cards"
	cmd "github.com/anuarkaliyev23/goker/pkg/cmd/utils"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func generateHands(playersCount int) [][]cards.Card {
	deck := cards.NewFullDeck()
	deck.Shuffle()
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

func combinationOf(representation string) cards.Combination {
	cs, err := cmd.ParseCards(representation)
	if err != nil {
		panic(err)
	}
	
	combination, err := cards.NewCombination(cs)
	if err != nil {
		panic(err)
	}

	return *combination
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

func card(face cards.Face, suit cards.Suit) cards.Card {
	card, err := cards.NewCard(face, suit)
	if err != nil {
		panic(err)
	}
	return *card
}

func Test_strongestHandCombination(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		t.Run("hand: AsAd, board: AhThKs, extra: Ts8h", func(t *testing.T) {
			hand := []cards.Card{card(cards.Ace, cards.Spades), card(cards.Ace, cards.Diamonds)}
			board := []cards.Card{card(cards.Ace, cards.Hearts), card(cards.Ten, cards.Hearts), card(cards.King, cards.Spades)}
			extra := []cards.Card{card(cards.Ten, cards.Spades), card(cards.Eight, cards.Hearts)}
	
			combination := strongestHandCombination(hand, board, extra)
			require.Equal(t, cards.FullHouse, combination.Type())
			require.Equal(t, cards.Ace, combination.MainCard())
			require.Equal(t, cards.Ten, *combination.SecondaryCard())
		})

		t.Run("hand: AsAd, board: 8s8h8d8c, extra: Qs", func(t *testing.T) {
			hand := []cards.Card{card(cards.Ace, cards.Spades), card(cards.Ace, cards.Diamonds)}
			board := []cards.Card{
				card(cards.Eight, cards.Spades), 
				card(cards.Eight, cards.Hearts), 
				card(cards.Eight, cards.Diamonds),
				card(cards.Eight, cards.Clubs),
			}
			extra := []cards.Card{card(cards.Queen, cards.Spades)}

			combination := strongestHandCombination(hand, board, extra)
			require.Equal(t, cards.FourOfAKind, combination.Type())
			require.Equal(t, cards.Ace, combination.Kickers()[0].Face())
		})

		t.Run("hand: AsAd, board: 8s8h8d8cQs, extra: (empty)", func(t *testing.T) {
			hand := []cards.Card{card(cards.Ace, cards.Spades), card(cards.Ace, cards.Diamonds)}
			board := []cards.Card{
				card(cards.Eight, cards.Spades), 
				card(cards.Eight, cards.Hearts), 
				card(cards.Eight, cards.Diamonds),
				card(cards.Eight, cards.Clubs),
				card(cards.Queen, cards.Spades),
			}
			extra := []cards.Card{}

			combination := strongestHandCombination(hand, board, extra)
			require.Equal(t, cards.FourOfAKind, combination.Type())
			require.Equal(t, cards.Ace, combination.Kickers()[0].Face())
		})
	})
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

func TestHandOddsIteration_StrongestCombination(t *testing.T) {
	t.Run("negative", func(t *testing.T) {
		t.Run("Empty combinations list", func(t *testing.T) {
			iteration := HandOddsIteration{
				Combinations: []cards.Combination{},
			}
			
			iterationStrongest, err := iteration.StrongestCombination()
			require.Error(t, err)
			require.Nil(t, iterationStrongest)
		})
	})

	t.Run("positive", func(t *testing.T) {
		t.Run("2 combinations", func(t *testing.T) {
			t.Run("AAAAK vs AAAKK", func(t *testing.T) {

				strongest, err := cards.NewCombination([]cards.Card{
						card(cards.Ace, cards.Clubs), 
						card(cards.Ace, cards.Spades),
						card(cards.Ace, cards.Diamonds), 
						card(cards.Ace, cards.Hearts),
						card(cards.King, cards.Clubs),
					},
				)

				require.NoError(t, err)

				weakest, err := cards.NewCombination([]cards.Card{
						card(cards.Ace, cards.Clubs), 
						card(cards.Ace, cards.Spades),
						card(cards.Ace, cards.Diamonds), 
						card(cards.King, cards.Hearts),
						card(cards.King, cards.Clubs),
					},
				)
				require.NoError(t, err)
				
				iteration := HandOddsIteration{
					Combinations: []cards.Combination{*strongest, *weakest},
				}
				
				iterationStrongest, err := iteration.StrongestCombination()
				require.NoError(t, err)

				require.Equal(t, strongest, iterationStrongest)
				require.NotEqual(t, weakest, iterationStrongest)
			})
		})
	})
}

func TestHandOddsResult_PlayerWins(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		t.Run("wins <= odds", func(t *testing.T) {
			handOdds := generateHandOdds(2, 1000)
			wins, err := handOdds.PlayerWins(0)
			require.NoError(t, err)
			require.True(t, wins <= handOdds.Config.IterationsCount)
		})

		t.Run("KT vs AK", func(t *testing.T) {
			iteration_1 := HandOddsIteration{
				Combinations: []cards.Combination{
					combinationOf("KsTs7h8hKd"), //lose
					combinationOf("AhKh7h8hKd"), //win
				},
			}

			iteration_2 := HandOddsIteration{
				Combinations: []cards.Combination{
					combinationOf("KsTs2s6s7s"), //win
					combinationOf("AhKh7s6sAd"), //lose
				},
			}

			iteration_3 := HandOddsIteration{
				Combinations: []cards.Combination{
					combinationOf("KsTsThTdAd"), //win
					combinationOf("AhKhThTdAd"), //lose
				},
			}

			handOddsResult := HandOddsResult{
				Config: HandOddsConfig{
					IterationsCount: 1000,
					Hands: [][]cards.Card{
						{card(cards.King, cards.Spades), card(cards.Ten, cards.Spades)},
						{card(cards.Ace, cards.Hearts), card(cards.King, cards.Hearts)},
					},
					Board: []cards.Card{},
				},
				Iterations: []HandOddsIteration{iteration_1, iteration_2, iteration_3},
			}
	
			firstPlayerWins, err := handOddsResult.PlayerWins(0)
			require.NoError(t, err)

			secondPlayerWins, err := handOddsResult.PlayerWins(1)
			require.NoError(t, err)

			require.Equal(t, 2, firstPlayerWins)
			require.Equal(t, 1, secondPlayerWins)

		})
	})

}

func TestHandOddsIteration_Winner(t *testing.T) {
	t.Run("negative", func(t *testing.T) {
		t.Run("Empty combination list", func(t *testing.T) {
			iteration := HandOddsIteration{
				Combinations: []cards.Combination{},
			}
			winners, err := iteration.Winner()
			
			require.Error(t, err)
			require.Equal(t, -1, winners)

		})
	})

	t.Run("positive", func(t *testing.T) {
		t.Run("AAAAK vs AAAKK", func(t *testing.T) {
			strongest, err := cards.NewCombination([]cards.Card{
					card(cards.Ace, cards.Clubs), 
					card(cards.Ace, cards.Spades),
					card(cards.Ace, cards.Diamonds), 
					card(cards.Ace, cards.Hearts),
					card(cards.King, cards.Clubs),
				},
			)

			require.NoError(t, err)

			weakest, err := cards.NewCombination([]cards.Card{
					card(cards.Ace, cards.Clubs), 
					card(cards.Ace, cards.Spades),
					card(cards.Ace, cards.Diamonds), 
					card(cards.King, cards.Hearts),
					card(cards.King, cards.Clubs),
				},
			)
			require.NoError(t, err)
			
			iteration := HandOddsIteration{
				Combinations: []cards.Combination{*strongest, *weakest},
			}
			
			winner, err := iteration.Winner()
			require.NoError(t, err)
			require.Equal(t, 0, winner)
		})

		t.Run("AAKK8 vs AAAKK8 vs AAKK7", func(t *testing.T) {

			strongest, err := cards.NewCombination([]cards.Card{
					card(cards.Ace, cards.Clubs), 
					card(cards.Ace, cards.Spades),
					card(cards.King, cards.Diamonds), 
					card(cards.King, cards.Hearts),
					card(cards.Eight, cards.Clubs),
				},
			)

			require.NoError(t, err)

			weakest, err := cards.NewCombination([]cards.Card{
					card(cards.Ace, cards.Clubs), 
					card(cards.Ace, cards.Spades),
					card(cards.King, cards.Diamonds), 
					card(cards.King, cards.Hearts),
					card(cards.Eight, cards.Diamonds),
				},
			)
			require.NoError(t, err)

			tie, err := cards.NewCombination([]cards.Card{
					card(cards.Ace, cards.Clubs), 
					card(cards.Ace, cards.Spades),
					card(cards.King, cards.Diamonds), 
					card(cards.King, cards.Hearts),
					card(cards.Seven, cards.Diamonds),
				},
			)

			require.NoError(t, err)

			iteration := HandOddsIteration{
				Combinations: []cards.Combination{*strongest, *weakest, *tie},
			}
			
			winner, err := iteration.Winner()
			require.NoError(t, err)
			require.Equal(t, -1, winner)
		})
	})
}
