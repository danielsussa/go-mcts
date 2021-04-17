package tictactoe

import (
	"fmt"
	"github.com/danielsussa/mcts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExample1(t *testing.T) {
	game := ticTacGame{
		playerTurn: O,
		board: []player{
			O, E, E,
			E, E, E,
			E, E, E,
		},
	}
	tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
	nodes, err := tree.Start(game)
	assert.NoError(t, err)
	assert.Equal(t, nodes.NodeScore[0].State.ID(), fmt.Sprintf("%s",[]player{
			O, E, E,
			E, X, E,
			E, E, E,
	}))
}

func TestExample2(t *testing.T) {
	game := ticTacGame{
		playerTurn: O,
		board: []player{
			O, E, O,
			E, X, E,
			E, E, E,
		},
	}
	tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
	nodes, err := tree.Start(game)
	assert.NoError(t, err)
	assert.Equal(t, nodes.NodeScore[0].State.ID(), fmt.Sprintf("%s",[]player{
			O, X, O,
			E, X, E,
			E, E, E,
	}))
}

func TestExample4(t *testing.T) {
	game := ticTacGame{
		playerTurn: O,
		board: []player{
			O, X, O,
			X, X, O,
			E, O, E,
		},
	}
	tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
	nodes, err := tree.Start(game)
	assert.NoError(t, err)
	assert.Equal(t, nodes.NodeScore[0].State.ID(), fmt.Sprintf("%s",[]player{
			O, X, O,
			X, X, O,
			E, O, X,
	}))
}

func TestExample5(t *testing.T) {
	game := ticTacGame{
		playerTurn: O,
		board: []player{
			X, E, E,
			E, O, E,
			O, E, X,
		},
	}
	tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
	nodes, err := tree.Start(game)
	assert.NoError(t, err)
	assert.Equal(t, nodes.NodeScore[0].State.ID(), fmt.Sprintf("%s",[]player{
			X, E, X,
			E, O, E,
			O, E, X,
	}))
}

func TestExample6(t *testing.T) {
	game := ticTacGame{
		playerTurn: O,
		board: []player{
			X, O, X,
			E, O, E,
			O, E, X,
		},
	}
	tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
	nodes, err := tree.Start(game)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%s",[]player{
			X, O, X,
			E, O, X,
			O, E, X,
	}), nodes.NodeScore[0].State.ID())
}

func TestExample7(t *testing.T) {
	game := ticTacGame{
		playerTurn: O,
		board: []player{
			O, X, O,
			E, X, E,
			O, E, X,
		},
	}
	tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
	nodes, err := tree.Start(game)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%s",[]player{
			O, X, O,
			E, X, E,
			O, X, X,
	}), nodes.NodeScore[0].State.ID())
}


func TestExample8(t *testing.T) {
	game := ticTacGame{
		playerTurn: O,
		board: []player{
			E, E, E,
			O, X, E,
			E, E, O,
		},
	}
	tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
	nodes, err := tree.Start(game)
	assert.NoError(t, err)
	gameBestMove := nodes.NodeScore[0].State.(ticTacGame)
	assert.Equal(t, gameBestMove.board[1] == X || gameBestMove.board[7] == X, true)
}

func TestExample9(t *testing.T) {
	game := ticTacGame{
		playerTurn: O,
		board: []player{
			E, E, E,
			O, X, O,
			E, X, O,
		},
	}
	tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
	nodes, err := tree.Start(game)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%s",[]player{
			E, X, E,
			O, X, O,
			E, X, O,
	}), nodes.NodeScore[0].State.ID())
}

func TestExample10(t *testing.T) {
	game := ticTacGame{
		playerTurn: O,
		board: []player{
			E, E, E,
			O, X, O,
			O, X, O,
		},
	}
	tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
	nodes, err := tree.Start(game)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%s",[]player{
		E, X, E,
		O, X, O,
		O, X, O,
	}), nodes.NodeScore[0].State.ID())
}

func TestExample11(t *testing.T) {
	game := ticTacGame{
		playerTurn: E,
		board: []player{
			E, E, E,
			E, E, E,
			E, E, E,
		},
	}
	tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
	nodes, err := tree.Start(game)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%s",[]player{
			E, E, E,
			E, X, E,
			E, E, E,
	}), nodes.NodeScore[0].State.ID())
}