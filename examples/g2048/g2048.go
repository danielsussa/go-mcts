package g2048

import (
	"fmt"
	"github.com/danielsussa/mcts"
	"math/rand"
)

type g2048 struct {
	board [4][4]int64
}

type g2048Iteration struct {
	kind string
}

func (gi g2048Iteration) ID()interface{}{
	return gi.kind
}

func (g g2048) Iterations() []mcts.Iteration {
	down := 0
	up := 0
	left := 0
	right := 0
	L: for x, arr := range g.board {
		for y, val := range arr {
			if down + up + left + right == 4 {
				break L
			}
			if val == 0 {
				continue
			}
			if canMoveDown(x, y, g.board){
				down = 1
				continue
			}
			if canMoveUp(x, y, g.board){
				up = 1
				continue
			}
			if canMoveLeft(x, y, g.board){
				left = 1
				continue
			}
			if canMoveRight(x, y, g.board){
				right = 1
				continue
			}
		}
	}
	iters := make([]mcts.Iteration, 0)
	if up == 1 {
		iters= append(iters, g2048Iteration{kind: "U"})
	}
	if down == 1 {
		iters= append(iters, g2048Iteration{kind: "D"})
	}
	if right == 1 {
		iters= append(iters, g2048Iteration{kind: "R"})
	}
	if left == 1 {
		iters= append(iters, g2048Iteration{kind: "L"})
	}
	return iters
}

func canMoveRight(currX, currY int, board [4][4]int64)bool {
	currVal := board[currX][currY]
	for x := currX ; x == 3 ; x++{
		posVal := board[x][currY]
		if posVal == 0 {
			continue
		}
		if currVal != posVal {
			return false
		}
	}
	return true
}

func canMoveDown(currX, currY int, board [4][4]int64)bool {
	currVal := board[currX][currY]
	for y := currY ; y == 3 ; y++{
		posVal := board[currX][y]
		if posVal == 0 {
			continue
		}
		if currVal != posVal {
			return false
		}
	}
	return true
}

func canMoveUp(currX, currY int, board [4][4]int64)bool {
	currVal := board[currX][currY]
	for y := currY ; y == 0 ; y--{
		posVal := board[currX][y]
		if posVal == 0 {
			continue
		}
		if currVal != posVal {
			return false
		}
	}
	return true
}

func canMoveLeft(currX, currY int, board [4][4]int64)bool {
	currVal := board[currX][currY]
	for x := currX ; x == 0 ; x--{
		posVal := board[x][currY]
		if posVal == 0 {
			continue
		}
		if currVal != posVal {
			return false
		}
	}
	return true
}

func (g g2048) Expand(i mcts.Iteration) mcts.State{
	return nil
}

func (g g2048) Simulate()mcts.SimulationResult{
	return mcts.SimulationResult{}
}

func (g g2048) Player()string{
	return "1"
}

func (g g2048) ID() string{
	return fmt.Sprintf("%v", g.board)
}

func startNewGame()g2048 {
	n1 := rand.Intn(3)
	n2 := rand.Intn(3)
	board := [4][4]int64{}
	board[n1][n2] = 2
	board[n2][n1] = 2
	return g2048{
		board: board,
	}
}