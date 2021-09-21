package g2048

import (
	"fmt"
	"github.com/danielsussa/mcts"
	"math/rand"
)

type g2048 struct {
	board [][]int
	score int
	stats g2048stats
}

type g2048stats struct {
	statistics [][]int
	iterations int
}

func (gs g2048stats) print() {
	fmt.Println(fmt.Sprintf("------- STATS --------"))
	for y := 0; y < 4; y++ {
		fmt.Println()
		for x := 0; x < 4; x++ {
			fmt.Print(fmt.Sprintf("%2f ", float64(gs.statistics[x][y])/float64(gs.iterations)))
		}
	}
	fmt.Println()
}

func (g g2048) Copy() mcts.State {
	return g2048{
		board: copy2DArr(g.board),
		score: g.score,
		stats: g2048stats{
			statistics: copy2DArr(g.stats.statistics),
			iterations: g.stats.iterations,
		},
	}
}

func (g g2048) Iterations() []interface{} {
	iterations := getAllIterations(g.board)
	return iterations
}

func (g g2048) Simulate() float64 {
	score := g.score

	for i := 0; i < 3; i++ {
		//print2048(board, score)
		allIterations := getAllIterations(g.board)
		if len(allIterations) == 0 {
			score = 0
			break
		}
		switch allIterations[rand.Intn(len(allIterations))] {
		case "D":
			score += computeDown(g.board)
		case "U":
			score += computeUp(g.board)
		case "L":
			score += computeLeft(g.board)
		case "R":
			score += computeRight(g.board)
		}
		// add random move
		//print2048(board, score)
		addNumberOnBoard(g.board)
	}

	return float64(score)
}

func (g g2048) ID() string {
	return fmt.Sprintf("%v", g.board)
}

func getAllIterations(board [][]int) []interface{} {
	down := 0
	up := 0
	left := 0
	right := 0
	for y := 0; y < 4; y++ {
		if down+up+left+right == 4 {
			break
		}
		for x := 0; x < 4; x++ {
			if board[x][y] == 0 {
				continue
			}
			if canMoveDown(x, y, board).canMove {
				down = 1
			}
			if canMoveUp(x, y, board).canMove {
				up = 1
			}
			if canMoveLeft(x, y, board).canMove {
				left = 1
			}
			if canMoveRight(x, y, board).canMove {
				right = 1
			}
		}
	}

	iters := make([]interface{}, 0)
	if up == 1 {
		iters = append(iters, "U")
	}
	if down == 1 {
		iters = append(iters, "D")
	}
	if right == 1 {
		iters = append(iters, "R")
	}
	if left == 1 {
		iters = append(iters, "L")
	}
	return iters
}

type newPosition struct {
	canMove bool
	x       int
	y       int
}

func canMoveRight(currX, currY int, board [][]int) newPosition {
	posEval := newPosition{}
	currVal := board[currX][currY]
	if currVal == 0 {
		return newPosition{}
	}
	for x := currX; x < 4; x++ {
		if currX == x {
			continue
		}
		posVal := board[x][currY]
		if posVal == 0 || posVal == currVal {
			posEval = newPosition{
				canMove: true,
				x:       x,
				y:       currY,
			}
		}
	}
	return posEval
}

func canMoveDown(currX, currY int, board [][]int) newPosition {
	posEval := newPosition{}
	currVal := board[currX][currY]
	if currVal == 0 {
		return newPosition{}
	}
	for y := currY; y < 4; y++ {
		if currY == y {
			continue
		}
		posVal := board[currX][y]
		if posVal == 0 || posVal == currVal {
			posEval = newPosition{
				canMove: true,
				x:       currX,
				y:       y,
			}
		}
	}
	return posEval
}

func canMoveUp(currX, currY int, board [][]int) newPosition {
	posEval := newPosition{}
	currVal := board[currX][currY]
	if currVal == 0 {
		return newPosition{}
	}
	for y := currY; y >= 0; y-- {
		if currY == y {
			continue
		}
		posVal := board[currX][y]
		if posVal == 0 || posVal == currVal {
			posEval = newPosition{
				canMove: true,
				x:       currX,
				y:       y,
			}
		}
	}
	return posEval
}

func canMoveLeft(currX, currY int, board [][]int) newPosition {
	posEval := newPosition{}
	currVal := board[currX][currY]
	if currVal == 0 {
		return newPosition{}
	}
	for x := currX; x >= 0; x-- {
		if currX == x {
			continue
		}
		posVal := board[x][currY]
		if posVal == 0 || posVal == currVal {
			posEval = newPosition{
				canMove: true,
				x:       x,
				y:       currY,
			}
		}
	}
	return posEval
}

func print2048(board [][]int, score int) {
	fmt.Print("\033[H\033[2J")
	fmt.Println(fmt.Sprintf("------- %v --------", score))
	for y := 0; y < 4; y++ {
		fmt.Println()
		for x := 0; x < 4; x++ {
			fmt.Print(fmt.Sprintf("%-6d", board[x][y]))
		}
	}
	fmt.Println()
}

func (g g2048) Expand(i interface{}) mcts.State {
	score := 0
	if i.(string) == "D" {
		score += computeDown(g.board)
	} else if i.(string) == "U" {
		score += computeUp(g.board)
	} else if i.(string) == "R" {
		score += computeRight(g.board)
	} else if i.(string) == "L" {
		score += computeLeft(g.board)
	}
	return g2048{board: g.board, score: g.score + score, stats: g2048stats{
		statistics: addStatistic(g.board, g.stats.statistics),
		iterations: g.stats.iterations + 1,
	}}
}

func addStatistic(board [][]int, s [][]int) [][]int {
	copyStatistic := copy2DArr(s)
	m := struct {
		y      int
		x      int
		number int
	}{}
	for y := 3; y >= 0; y-- {
		for x := 0; x < 4; x++ {
			if board[x][y] > m.number {
				m.number = board[x][y]
				m.x = x
				m.y = y
			}
		}
	}
	copyStatistic[m.x][m.y]++
	return copyStatistic
}

type coordinate struct {
	x int
	y int
}

func getFreePlaces(board [][]int) []coordinate {
	freePlaces := make([]coordinate, 0)
	for y := 3; y >= 0; y-- {
		for x := 0; x < 4; x++ {
			if board[x][y] == 0 {
				freePlaces = append(freePlaces, coordinate{
					x: x,
					y: y,
				})
			}
		}
	}
	return freePlaces
}

func addNumberOnBoard(board [][]int) {
	freePlaces := getFreePlaces(board)
	if len(freePlaces) == 0 {
		return
	}

	freePlace := freePlaces[rand.Intn(len(freePlaces))]
	fRand := rand.Float64()
	val := 2
	if fRand >= 0.9 {
		val = 4
	}
	board[freePlace.x][freePlace.y] = val
}

func addNumberOnBoardCord(cord coordinate, board [][]int) {
	fRand := rand.Float64()
	val := 2
	if fRand >= 0.9 {
		val = 4
	}
	board[cord.x][cord.y] = val
}

func computeDown(board [][]int) int {
	score := 0
	for y := 3; y >= 0; y-- {
		for x := 0; x < 4; x++ {
			score += executeCompute("D", x, y, board)
		}
	}
	return score
}

func computeUp(board [][]int) int {
	score := 0
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			score += executeCompute("U", x, y, board)
		}
	}
	return score
}

func computeRight(board [][]int) int {
	score := 0
	for y := 0; y < 4; y++ {
		for x := 3; x >= 0; x-- {
			score += executeCompute("R", x, y, board)
		}
	}
	return score
}

func computeLeft(board [][]int) int {
	score := 0
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			score += executeCompute("L", x, y, board)
		}
	}
	return score
}

func executeCompute(kind string, x, y int, board [][]int) int {
	posVal := board[x][y]
	if posVal == 0 {
		return 0
	}
	newPos := newPosition{}
	switch kind {
	case "D":
		newPos = canMoveDown(x, y, board)
	case "U":
		newPos = canMoveUp(x, y, board)
	case "R":
		newPos = canMoveRight(x, y, board)
	case "L":
		newPos = canMoveLeft(x, y, board)
	}
	if !newPos.canMove {
		return 0
	}
	newPosVal := board[newPos.x][newPos.y]
	if newPosVal == 0 {
		board[newPos.x][newPos.y] = posVal
	} else {
		board[newPos.x][newPos.y] += posVal
	}
	board[x][y] = 0
	return 2 * newPosVal
}

func copy2DArr(src [][]int) [][]int {
	duplicate := make([][]int, len(src))
	for i := range src {
		duplicate[i] = make([]int, len(src[i]))
		copy(duplicate[i], src[i])
	}
	return duplicate
}

func startNewGame() g2048 {
	return g2048{
		board: newArr(),
		stats: g2048stats{
			statistics: newArr(),
			iterations: 0,
		},
	}
}

func newArr() [][]int {
	board := make([][]int, 4)
	board[0] = make([]int, 4)
	board[1] = make([]int, 4)
	board[2] = make([]int, 4)
	board[3] = make([]int, 4)
	return board
}
