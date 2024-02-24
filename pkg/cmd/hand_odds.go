package cmd

import (
	"fmt"

	"github.com/anuarkaliyev23/goker/pkg/cards"
	utils "github.com/anuarkaliyev23/goker/pkg/cmd/utils"
	"github.com/anuarkaliyev23/goker/pkg/texas/calc"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var boardFlag string
var handsFlag []string
var iterationsFlag int

var handOddsCmd = &cobra.Command{
	Use: "hand-odds",
	Short: "compare hand odds with optional board",
	RunE: func(c *cobra.Command, args []string) error {
		boardCards, err := utils.ParseCards(boardFlag)
		if err != nil {
			return err
		}

		hands := lo.Map(handsFlag, func(representation string, _ int) []cards.Card {
			cards, err := utils.ParseCards(representation)
			if err != nil {
				return nil
			}

			return cards
		})

		handOddsConfig := calc.HandOddsConfig{
			Board: boardCards,
			Hands: hands,
			IterationsCount: iterationsFlag,
		}
	
		handOdds, err := calc.HandOdds(handOddsConfig)
		if err != nil {
			return err
		}

		for player := 0; player < len(handsFlag); player++ {
			wins, err := handOdds.PlayerWins(player)
			if err != nil {
				return err
			}

			fmt.Println(fmt.Sprintf("Player[%d]: %d", player, wins))
		}
		return nil
	},
}

func init() {
	//TODO figure out flags
	handOddsCmd.Flags().StringVar(&boardFlag, "board", "", "used to pass community/board cards")
	handOddsCmd.Flags().StringSliceVarP(&handsFlag, "hands", "", nil, "used to pass hole/hand cards")
	handOddsCmd.Flags().IntVarP(&iterationsFlag, "iterations", "i", 1000, "how much iterations should simulation have")

	texasCmd.AddCommand(handOddsCmd)
}
