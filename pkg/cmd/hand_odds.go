package cmd

import (
	"fmt"

	"github.com/anuarkaliyev23/goker/pkg/calc"
	"github.com/anuarkaliyev23/goker/pkg/cards"
	utils "github.com/anuarkaliyev23/goker/pkg/cmd/utils"
	"github.com/anuarkaliyev23/goker/pkg/game"
	"github.com/fatih/color"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

const (
	TexasFlagName = "texas"
	ShortDeckFlagName = "short-deck"
	OmahaFlagName = "omaha"
)

var boardFlag string
var handsFlag []string
var iterationsFlag int

var texasFlag bool
var shortDeckFlag bool
var omahaFlag bool

var handOddsCmd = &cobra.Command{
	Use: "hand-odds",
	Short: "compare hand odds with optional board",
	RunE: func(c *cobra.Command, args []string) error {
		err, executionDuration := utils.MeasureTime(func() error {
			var gameConfig game.Config
			
			if texasFlag {
				gameConfig = game.NewTexasConfig()
			} else if shortDeckFlag {
				gameConfig = game.NewShortDeckConfig()
			} else if omahaFlag {
				gameConfig = game.NewOmahaConfig()
			}

			handOdds, err := handOdds(boardFlag, handsFlag, iterationsFlag, gameConfig)
			if err != nil {
				return err
			}
			
			playersWins, err := handOdds.WinRates()
			if err != nil {
				return err
			}
			
			wonPlayer := lo.Max(playersWins)
			wonPlayerIndex := lo.IndexOf(playersWins, wonPlayer)

			for player := 0; player < len(playersWins); player++ {
				s := fmt.Sprintf("[%v]: %.1f%%", handsFlag[player], playersWins[player] * 100)
				if player == wonPlayerIndex {
					color.Green(s)
				} else {
					color.Red(s)
				}
			}

			ties, err := handOdds.TiePercentage()
			if err != nil {
				return err
			}

			color.Yellow(fmt.Sprintf("Ties: %.1f%%", ties * 100))
			return nil
		})
		if err != nil {
			return err
		}
		color.White(fmt.Sprintf("%d ms\n", executionDuration))
		return nil
	},
}

func handOdds(boardRepresentation string, handsRepresentation []string, iterations int, gameConfig game.Config) (*calc.HandOddsResult, error) {
	boardCards, err := utils.ParseCards(boardRepresentation)
	if err != nil {
		return nil, err
	}

	hands := lo.Map(handsRepresentation, func(representation string, _ int) []cards.Card {
		cards, err := utils.ParseCards(representation)
		if err != nil {
			return nil
		}

		return cards
	})

	handOddsConfig := calc.HandOddsConfig{
		Board: boardCards,
		Hands: hands,
		IterationsCount: iterations,
		GameConfig: gameConfig,
	}

	return calc.HandOdds(handOddsConfig)
}

func countPlayerWins(handOdds *calc.HandOddsResult, handsRepresentation []string) ([]int, error) {
	result := []int{}

	for player := 0; player < len(handsRepresentation); player++ {
		wins, err := handOdds.PlayerWins(player)
		if err != nil {
			return nil, err
		}
		
		result = append(result, wins)
	}
	return result, nil
}

func init() {
	handOddsCmd.Flags().StringVar(&boardFlag, "board", "", "used to pass community/board cards")
	handOddsCmd.Flags().StringSliceVarP(&handsFlag, "hands", "", nil, "used to pass hole/hand cards")
	handOddsCmd.Flags().IntVarP(&iterationsFlag, "iterations", "i", 1000, "how much iterations should simulation have")

	handOddsCmd.Flags().BoolVar(&texasFlag, TexasFlagName, false, "flag to indicate Texas Hold'em")
	handOddsCmd.Flags().BoolVar(&shortDeckFlag, ShortDeckFlagName, false, "flag to indicate Short-Deck")
	handOddsCmd.Flags().BoolVar(&omahaFlag, OmahaFlagName, false, "flag to indicate Omaha")

	handOddsCmd.MarkFlagsOneRequired(TexasFlagName, ShortDeckFlagName, OmahaFlagName)
	handOddsCmd.MarkFlagsMutuallyExclusive(TexasFlagName, ShortDeckFlagName, OmahaFlagName)

	rootCmd.AddCommand(handOddsCmd)
}
