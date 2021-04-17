package tictactoe

import (
	"fmt"
	"github.com/danielsussa/mcts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithUniqueStart(t *testing.T) {
	// iteration 1
	{
		game := ticTacGame{
			playerTurn: X,
			board: []player{
				E, E, E,
				E, E, E,
				E, E, E,
			},
		}
		tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
		nodes, err := tree.Start(game)
		assert.NoError(t, err)
		assert.Equal(t, nodes.NodeScore[0].State.ID(), fmt.Sprintf("%s",[]player{
			E, E, E,
			E, X, E,
			E, E, E,
		}))
	}

	// iteration 2
	{
		game := ticTacGame{
			playerTurn: X,
			board: []player{
				E, E, O,
				E, X, E,
				E, E, E,
			},
		}
		tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
		nodes, err := tree.Start(game)
		assert.NoError(t, err)
		assert.Equal(t, nodes.NodeScore[0].State.ID(), fmt.Sprintf("%s",[]player{
			E, E, O,
			E, X, E,
			E, E, X,
		}))
	}

	// iteration 3
	{
		game := ticTacGame{
			playerTurn: X,
			board: []player{
				O, E, O,
				E, X, E,
				E, E, X,
			},
		}
		tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
		nodes, err := tree.Start(game)
		assert.NoError(t, err)
		assert.Equal(t, nodes.NodeScore[0].State.ID(), fmt.Sprintf("%s",[]player{
			O, X, O,
			E, X, E,
			E, E, X,
		}))
	}

	// iteration 4
	{
		game := ticTacGame{
			playerTurn: X,
			board: []player{
				O, X, O,
				E, X, E,
				O, E, X,
			},
		}
		tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
		nodes, err := tree.Start(game)
		assert.NoError(t, err)
		assert.Equal(t, nodes.NodeScore[0].State.ID(), fmt.Sprintf("%s",[]player{
			O, X, O,
			X, X, E,
			O, E, X,
		}))
	}

	// iteration 5
	{
		game := ticTacGame{
			playerTurn: X,
			board: []player{
				O, X, O,
				X, X, O,
				O, E, X,
			},
		}
		tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
		nodes, err := tree.Start(game)
		assert.NoError(t, err)
		assert.Equal(t, nodes.NodeScore[0].State.ID(), fmt.Sprintf("%s",[]player{
			O, X, O,
			X, X, O,
			O, X, X,
		}))
	}

}
