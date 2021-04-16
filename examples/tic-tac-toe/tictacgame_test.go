package tictactoe

import (
	"fmt"
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

	assert.Equal(t, nodes.NodeScore[0].State.ID(), fmt.Sprintf("%s",[]player{
			O, E, E,
			E, X, E,
			E, E, E,
	}))
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

	assert.Equal(t, nodes.NodeScore[0].State.ID(), fmt.Sprintf("%s",[]player{
			O, X, O,
			E, X, E,
			E, E, E,
	}))
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
	assert.Equal(t, nodes.NodeScore[0].State.ID(), fmt.Sprintf("%s",[]player{
			O, X, O,
			X, X, E,
			E, O, E,
	}))
}

func TestExample4(t *testing.T) {
	game := ticTacGame{
		playerTurn: X,
		board: []player{
			O, X, O,
			X, X, O,
			E, O, E,
		},
	}
	tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
	nodes := tree.Start(game)
	assert.Equal(t, nodes.NodeScore[0].State.ID(), fmt.Sprintf("%s",[]player{
			O, X, O,
			X, X, O,
			E, O, X,
	}))
}

func TestExample5(t *testing.T) {
	game := ticTacGame{
		playerTurn: X,
		board: []player{
			X, E, E,
			E, O, E,
			O, E, X,
		},
	}
	tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
	nodes := tree.Start(game)
	assert.Equal(t, nodes.NodeScore[0].State.ID(), fmt.Sprintf("%s",[]player{
			X, E, X,
			E, O, E,
			O, E, X,
	}))
}

// errado aqui, deveria escolher
func TestExample6(t *testing.T) {
	game := ticTacGame{
		playerTurn: X,
		board: []player{
			X, O, X,
			E, O, E,
			O, E, X,
		},
	}
	tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 5000})
	nodes := tree.Start(game)
	assert.Equal(t, fmt.Sprintf("%s",[]player{
			X, O, X,
			E, O, X,
			O, E, X,
	}), nodes.NodeScore[0].State.ID())
}