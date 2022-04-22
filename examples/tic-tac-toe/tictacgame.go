package tictactoe

import (
	"fmt"
	"github.com/danielsussa/mcts"
	"math/rand"
)

type player string

const (
	E player = "E"
	X player = "X"
	O player = "O"
)

type ticTacGame struct {
	board      []player
	playerTurn player
	lastMove   int
}

// until final game & result
func (t ticTacGame) Simulate() float64 {
	return t.simulate(t.getPlayerTurn())
}

func (t ticTacGame) Copy() mcts.State {
	newBoard := make([]player, len(t.board))
	copy(newBoard, t.board)
	return ticTacGame{
		board:      newBoard,
		playerTurn: t.playerTurn,
		lastMove:   t.lastMove,
	}
}

func (t ticTacGame) simulate(startPlayer player) float64 {
	p := startPlayer
	if t.winner() != E {
		return t.winner().toScore()
	}
	moved := t.randomMove(p)
	if !moved {
		return t.winner().toScore()
	}

	playerWinner := E
	for {
		playerWinner = t.winner()
		if playerWinner != E {
			break
		}
		p = nextPlayer(t)
		moved = t.randomMove(p)
		if !moved {
			return 0
		}
	}
	return playerWinner.toScore()
}

func nextPlayer(game ticTacGame) player {
	if game.playerTurn == E {
		return X
	}
	if game.playerTurn == X {
		return O
	}
	return X
}
func (p player) toScore() float64 {
	switch p {
	case X:
		return 1
	case E:
		return 0
	case O:
		return -1
	}
	return 0
}

func (t ticTacGame) winner() player {
	b := t.board
	if b[0] == b[1] && b[0] == b[2] && b[0] != E {
		return b[0]
	}
	if b[3] == b[4] && b[3] == b[5] && b[3] != E {
		return b[3]
	}
	if b[6] == b[7] && b[6] == b[8] && b[6] != E {
		return b[6]
	}

	if b[0] == b[3] && b[0] == b[6] && b[0] != E {
		return b[0]
	}
	if b[1] == b[4] && b[1] == b[7] && b[1] != E {
		return b[1]
	}
	if b[2] == b[5] && b[2] == b[8] && b[2] != E {
		return b[2]
	}

	if b[0] == b[4] && b[0] == b[8] && b[0] != E {
		return b[0]
	}
	if b[2] == b[4] && b[2] == b[6] && b[2] != E {
		return b[2]
	}
	return E
}

func (t ticTacGame) Iterations() []any {
	iters := make([]any, 0)
	for idx, place := range t.board {
		if place == E {
			iters = append(iters, idx)
		}
	}
	return iters
}

func (t ticTacGame) Expand(id any) mcts.State {
	t.move(id.(int), nextPlayer(t))
	return t
}

func (t ticTacGame) ID() string {
	return fmt.Sprintf("%v", t.board)
}

func (t ticTacGame) newWithNextPlayer() ticTacGame {
	newBoard := make([]player, len(t.board))
	copy(newBoard, t.board)
	return ticTacGame{
		board:      newBoard,
		playerTurn: nextPlayer(t),
	}
}

func (t ticTacGame) randomMove(p player) bool {
	free := make([]int, 0)

	for idx, place := range t.board {
		if place == E {
			free = append(free, idx)
		}
	}
	if len(free) == 0 {
		return false
	}
	place := free[rand.Intn(len(free))]
	t.board[place] = p
	return true
}

func (t ticTacGame) print() {
	fmt.Println("----------------------")
	fmt.Println(fmt.Sprintf("%s|%s|%s", t.board[0], t.board[1], t.board[2]))
	fmt.Println(fmt.Sprintf("%s|%s|%s", t.board[3], t.board[4], t.board[5]))
	fmt.Println(fmt.Sprintf("%s|%s|%s", t.board[6], t.board[7], t.board[8]))
}

func (t ticTacGame) getPlayerTurn() player {
	if t.playerTurn == E {
		return nextPlayer(t)
	}
	return t.playerTurn
}

func (t ticTacGame) move(idx int, p player) {
	t.board[idx] = p
}
