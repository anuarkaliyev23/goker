package cmd

import (
	utils "github.com/anuarkaliyev23/goker/pkg/cmd/utils"
	"github.com/spf13/cobra"
)

var boardFlag string
var handsFlag []string

var handOddsCmd = &cobra.Command{
	Use: "hand-odds",
	Short: "compare hand odds with optional board",
	RunE: func(c *cobra.Command, args []string) error {
		boardCards, err := utils.ParseCards(boardFlag)
		if err != nil {
			return err
		}

		handsFlag
	},
}

func init() {
	//TODO figure out flags
	handOddsCmd.Flags().StringVar(&boardFlag, "board", "", "used to pass community/board cards")
	handOddsCmd.Flags().StringSlice(&handsFlag, "hands", "used to pass hole/hand cards")

	texasCmd.AddCommand(handOddsCmd)
}
