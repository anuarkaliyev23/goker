package calc

import (
	"github.com/anuarkaliyev23/goker/pkg/cards"
)

type HandOddsCalcReport struct {
	Hands [][]cards.Card
	Board []cards.Card
	Wins map[int]int
	Ties int
}

