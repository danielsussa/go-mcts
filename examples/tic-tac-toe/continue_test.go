package tictactoe

import (
	"fmt"
	"github.com/danielsussa/mcts"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)



func TestWithUniqueStart(t *testing.T) {
	rand.Seed(0)
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

func TestWithContinue(t *testing.T) {
	game := ticTacGame{
		playerTurn: X,
		board: []player{
			E, E, E,
			E, E, E,
			E, E, E,
		},
	}
	tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})

	// AI move
	nodes, err := tree.Start(game)
	assert.NoError(t, err)
	game = nodes.NodeScore[0].State.(ticTacGame)

	assert.Equal(t, game.ID(), fmt.Sprintf("%s",[]player{
			E, E, E,
			E, X, E,
			E, E, E,
	}))

	// my move
	game = game.newWithMove(2, O)
	assert.Equal(t, game.ID(), fmt.Sprintf("%s",[]player{
			E, E, O,
			E, X, E,
			E, E, E,
	}))

	// AI move
	nodes, err = tree.Continue(game)
	assert.NoError(t, err)
	game = nodes.NodeScore[0].State.(ticTacGame)

	assert.Equal(t, game.ID(), fmt.Sprintf("%s",[]player{
			X, E, O,
			E, X, E,
			E, E, E,
	}))

}


func TestMultipleGame(t *testing.T) {
	winners := map[player]int{
		X: 0, O: 0, E: 0,
	}
	for i := 0; i < 100; i++{
		winners[newGame()]++
	}
	fmt.Println(winners)
}

func newGame()player{
	game := ticTacGame{
		playerTurn: O,
		board: []player{
			E, E, E,
			E, E, E,
			E, E, E,
		},
	}

	for {
		move := false
		game, move = game.newWithRandomMove()
		if !move || game.winner() != E{
			if game.winner() == O {
				return game.winner()
			}
			return game.winner()
		}
		//game.print()


		tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
		nodeScore, _ := tree.Start(game)

		if len(nodeScore.NodeScore) == 0 {
			return game.winner()
		}
		game = game.newWithMove(nodeScore.NodeScore[0].State.(ticTacGame).lastMove, X)

		if game.winner() != E {
			return game.winner()
		}
		//game.print()
	}
}