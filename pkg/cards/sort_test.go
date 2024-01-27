package cards

import (
	"sort"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestSort(t *testing.T) {
	t.Run("AsAdQdQsTh2h", func(t *testing.T){
		cards := []Card{
			{face: Ace, suit: Spades},
			{face: Ace, suit: Diamonds},
			{face: Queen, suit: Diamonds},
			{face: Queen, suit: Spades},
			{face: Ten, suit: Hearts},
			{face: Two, suit: Hearts},
		}

		shuffled := lo.Shuffle(cards)
		sort.Sort(ByFace(shuffled))
		faces := lo.Map(shuffled, func(card Card, index int) Face {
			return card.Face()
		})

		require.Equal(t, faces, []Face{Two, Ten, Queen, Queen, Ace, Ace})
	})
}
