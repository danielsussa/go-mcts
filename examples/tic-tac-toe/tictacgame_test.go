package tictactoe

import (
	"github.com/danielsussa/mcts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExample1(t *testing.T) {
	game := ticTacGame{
		playerTurn: X,
		board: []player{
			O, E, E,
			E, E, E,
			E, E, E,
		},
	}
	tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
	nodes := tree.Start(game)
	assert.Equal(t, nodes.NodeScore[0].State.ID(), "[O E E E X E E E E]")
}

func TestExample2(t *testing.T) {
	game := ticTacGame{
		playerTurn: X,
		board: []player{
			O, E, O,
			E, X, E,
			E, E, E,
		},
	}
	tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
	nodes := tree.Start(game)
	assert.Equal(t, nodes.NodeScore[0].State.ID(), "[O X O E X E E E E]")
}

func TestExample3(t *testing.T) {
	game := ticTacGame{
		playerTurn: X,
		board: []player{
			O, X, O,
			E, X, E,
			E, O, E,
		},
	}
	tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
	nodes := tree.Start(game)
	assert.Equal(t, nodes.NodeScore[0].State.ID(), "[O X O E X E E O E]")
}