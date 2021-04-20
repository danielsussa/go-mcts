package tictactoe

import (
	"fmt"
	"github.com/danielsussa/mcts"
	"testing"
)


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

	p := O
	for {
		move := game.randomMove(p)
		if !move || game.winner() != E{
			if game.winner() == O {
				return game.winner()
			}
			return game.winner()
		}
		//game.print()

		p = nextPlayer(game)


		tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 1000})
		nodeScore, _ := tree.Start(game)

		if len(nodeScore.NodeScore) == 0 {
			return game.winner()
		}
		game.move(nodeScore.NodeScore[0].State.(ticTacGame).lastMove, X)

		if game.winner() != E {
			return game.winner()
		}
		//game.print()
	}
}