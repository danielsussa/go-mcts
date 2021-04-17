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
	lastMove   uint
}

func (t ticTacGame) Player()string{
	return string(t.playerTurn)
}

// until final game & result
func (t ticTacGame) Simulate()(float64,string){
	//total := 0.0
	//for i := 0; i < 40 ; i++ {
	//	total += t.simulate()
	//}
	return t.simulate()
}

func (t ticTacGame) simulate()(float64,string){
	if t.winner() != E {
		return t.winner().winnerPlayer()
	}
	game, moved := t.newWithRandomMove()
	if !moved {
		return t.winner().winnerPlayer()
	}

	playerWinner := E
	for {
		playerWinner = game.winner()
		if playerWinner != E {
			break
		}
		game, moved = game.newWithRandomMove()
		if !moved {
			return 0, "E"
		}
	}
	return playerWinner.winnerPlayer()
}
func (p player) winnerPlayer()(float64, string) {
	switch p {
	case X:
		return 1, "X"
	case E:
		return 0, "E"
	case O:
		return -1, "O"
	}
	return 0, "E"
}

func (t ticTacGame)winner()player{
	b := t.board
	if b[0] == b[1] && b[0] == b[2] && b[0] != E{
		return b[0]
	}
	if b[3] == b[4] && b[3] == b[5] && b[3] != E{
		return b[3]
	}
	if b[6] == b[7] && b[6] == b[8] && b[6] != E{
		return b[6]
	}

	if b[0] == b[3] && b[0] == b[6] && b[0] != E{
		return b[0]
	}
	if b[1] == b[4] && b[1] == b[7] && b[1] != E{
		return b[1]
	}
	if b[2] == b[5] && b[2] == b[8] && b[2] != E{
		return b[2]
	}

	if b[0] == b[4] && b[0] == b[8] && b[0] != E{
		return b[0]
	}
	if b[2] == b[4] && b[2] == b[6] && b[2] != E{
		return b[2]
	}
	return E
}

func (t ticTacGame) MaxPlays() uint{
	total := uint(0)
	for _,place := range t.board {
		if place == E {
			total++
		}
	}
	return total
}

func (t ticTacGame) Expand(idx uint) mcts.State{
	if t.winner() != E {
		return nil
	}

	free := make([]uint, 0)
	for idx,place := range t.board {
		if place == E {
			free = append(free, uint(idx))
		}
	}

	return t.newWithMove(free[idx], t.playerTurn)
}

func (t ticTacGame) ID() string{
	return fmt.Sprintf("%v", t.board)
}

func (t ticTacGame) newWithRandomMove()(ticTacGame, bool){
	free := make([]int, 0)

	newBoard := make([]player,len(t.board))
	copy(newBoard, t.board)
	for idx,place := range t.board {
		if place == E {
			free = append(free, idx)
		}
	}
	if len(free) == 0 {
		return ticTacGame{}, false
	}
	place := free[rand.Intn(len(free))]
	newBoard[place] = t.playerTurn
	return ticTacGame{
		board:      newBoard,
		lastMove:   uint(place),
		playerTurn: t.otherTurn(),
	}, true
}

func (t ticTacGame) print() {
	fmt.Println("----------------------")
	fmt.Println(fmt.Sprintf("%s|%s|%s", t.board[0], t.board[1], t.board[2]))
	fmt.Println(fmt.Sprintf("%s|%s|%s", t.board[3], t.board[4], t.board[5]))
	fmt.Println(fmt.Sprintf("%s|%s|%s", t.board[6], t.board[7], t.board[8]))
}

func (t ticTacGame) otherTurn() player{
	if t.playerTurn == X {
		return O
	}else {
		return X
	}
}

func (t ticTacGame)newWithMove(idx uint, p player)ticTacGame{
	newBoard := make([]player,len(t.board))
	copy(newBoard, t.board)
	newBoard[idx] = p

	return ticTacGame{board: newBoard, playerTurn: t.otherTurn(), lastMove: idx}
}

func newTicTacGame()ticTacGame{
	return ticTacGame{
		playerTurn: O,
		board: []player{
			E, E, E,
			E, E, E,
			E, E, E,
		},
	}
}
