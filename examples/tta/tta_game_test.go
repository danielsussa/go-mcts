package tta

import (
	"fmt"
	"github.com/danielsussa/mcts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTTAGame(t *testing.T) {
	p1 := player{
		ID:           "p1",
		science:      0,
		rocks:        0,
		food:         0,
		culture:      0,
		totalActions: 4,
		remainAction: 4,
		civilCards:   nil,
	}
	game := ttaGame{
		civilCards: []string{res1, res1},
		players: []player{p1},
		currentPlayer: p1,
	}

	tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
	nodes, err := tree.Start(game)
	assert.NoError(t, err)
	fmt.Println(nodes)

}