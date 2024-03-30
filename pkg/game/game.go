package game

import (
	"fmt"

	"github.com/anuarkaliyev23/goker/pkg/cards"
)

type Game int

const (
	Texas Game = iota
	ShortDeck
	Omaha
	Custom
)

type Config struct {
	Game Game

	DeckGenerator func() cards.Deck
	HoleCardsCount int
	CommunityCardsCount int

	HoleCardsAllowedToUseCount int
	CommunityCardsAllowedToUseCount int

	MaxPlayers int
}

func NewCustomConfig(
		deckGenerator func() cards.Deck,
		holeCardsCount, 
		communityCardsCount, 
		holeCardsAllowedToUseCount, 
		communityCardsAllowedToUseCount, 
		maxPlayers int,
	) Config {
	return Config {
		Game: Custom,
		DeckGenerator: deckGenerator,
		HoleCardsCount: holeCardsCount,
		CommunityCardsCount: communityCardsCount,
		HoleCardsAllowedToUseCount: holeCardsAllowedToUseCount,
		CommunityCardsAllowedToUseCount: communityCardsAllowedToUseCount,
		MaxPlayers: maxPlayers,
	}
}

func NewConfig(game Game) (Config, error) {
	if game == Texas {
		return NewTexasConfig(), nil
	} else if game == Omaha {
		return NewOmahaConfig(), nil
	} else if game == ShortDeck {
		return NewShortDeckConfig(), nil
	} else {
		return Config{}, fmt.Errorf("could not construct config for game=[%v]", game)
	}
}

func NewTexasConfig() Config {
	return Config{
		Game: Texas,

		DeckGenerator: cards.NewFullDeck,
		HoleCardsCount: 2,
		CommunityCardsCount: 5,

		HoleCardsAllowedToUseCount: 2,
		CommunityCardsAllowedToUseCount: 5,

		MaxPlayers: 10,
	}
}

func NewShortDeckConfig() Config {
	return Config{
		Game: ShortDeck,

		DeckGenerator: cards.NewShortDeck,
		HoleCardsCount: 2,
		CommunityCardsCount: 5,

		HoleCardsAllowedToUseCount: 2,
		CommunityCardsAllowedToUseCount: 5,

		MaxPlayers: 10,
	}
}

func NewOmahaConfig() Config {
	return Config {
		Game: Omaha,
		
		DeckGenerator: cards.NewFullDeck,
		HoleCardsCount: 4,
		CommunityCardsCount: 5,

		HoleCardsAllowedToUseCount: 2,
		CommunityCardsAllowedToUseCount: 3,

		MaxPlayers: 10,
	}
}
