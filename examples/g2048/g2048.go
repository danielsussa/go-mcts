package g2048

import (
	"fmt"
	"github.com/danielsussa/mcts"
	"math/rand"
)

type g2048 struct {
	board [][]int
	score int
}

type g2048Iteration struct {
	kind string
}

func (gi g2048Iteration) ID()interface{}{
	return gi.kind
}

func (g g2048) Iterations() []mcts.Iteration {
	iterations := getAllIterations(g.board)
	return iterations
}

func (g g2048) Simulate()mcts.SimulationResult{
	score := 0
	L: for _, cord := range getFreePlaces(g.board){
		board := copy2DArr(g.board)
		addNumberOnBoardCord(cord, board)

		for i := 0 ; i < 5 ; i++{
			//print2048(board, score)
			allIterations := getAllIterations(board)
			if len(allIterations) == 0 {
				score = 0
				break L
			}
			nextIteration := allIterations[rand.Intn(len(allIterations))].(g2048Iteration)
			switch nextIteration.kind {
			case "D":
				score += computeDown(board)
			case "U":
				score += computeUp(board)
			case "L":
				score += computeLeft(board)
			case "R":
				score += computeRight(board)
			}
			// add random move
			//print2048(board, score)
			addNumberOnBoard(board)
		}
	}


	return mcts.SimulationResult{
		Score:  float64(score),
		Winner: "1",
		Player: "1",
	}
}

func (g g2048) Player()string{
	return "1"
}

func (g g2048) ID() string{
	return fmt.Sprintf("%v", g.board)
}

func getAllIterations(board [][]int) []mcts.Iteration{
	down := 0
	up := 0
	left := 0
	right := 0
	for y := 0 ; y < 4 ; y++ {
		if down + up + left + right == 4 {
			break
		}
		for x := 0; x < 4; x++ {
			if board[x][y] == 0 {
				continue
			}
			if canMoveDown(x, y, board).canMove{
				down = 1
			}
			if canMoveUp(x, y, board).canMove{
				up = 1
			}
			if canMoveLeft(x, y, board).canMove{
				left = 1
			}
			if canMoveRight(x, y, board).canMove{
				right = 1
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

type newPosition struct {
	canMove bool
	x int
	y int
}

func canMoveRight(currX, currY int, board [][]int)newPosition {
	posEval := newPosition{}
	currVal := board[currX][currY]
	if currVal == 0 {
		return newPosition{}
	}
	for x := currX ; x < 4 ; x++{
		if currX == x {
			continue
		}
		posVal := board[x][currY]
		if posVal == 0 || posVal == currVal{
			posEval = newPosition{
				canMove: true,
				x:       x,
				y:       currY,
			}
		}
	}
	return posEval
}

func canMoveDown(currX, currY int, board [][]int)newPosition {
	posEval := newPosition{}
	currVal := board[currX][currY]
	if currVal == 0 {
		return newPosition{}
	}
	for y := currY ; y < 4 ; y++{
		if currY == y {
			continue
		}
		posVal := board[currX][y]
		if posVal == 0 || posVal == currVal{
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
	for y := currY ; y >= 0 ; y--{
		if currY == y {
			continue
		}
		posVal := board[currX][y]
		if posVal == 0 || posVal == currVal{
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
	for x := currX ; x >= 0 ; x--{
		if currX == x {
			continue
		}
		posVal := board[x][currY]
		if posVal == 0 || posVal == currVal{
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
	fmt.Println(fmt.Sprintf("------- %v --------", score))
	for y := 0 ; y < 4 ; y++ {
		fmt.Println()
		for x := 0 ; x < 4 ; x++ {
			fmt.Print(fmt.Sprintf("%-6d", board[x][y]))
		}
	}
	fmt.Println()
}

func (g g2048) Expand(i mcts.Iteration) mcts.State{
	board := copy2DArr(g.board)
	score := 0
	if i.ID().(string) == "D"{
		score += computeDown(board)
	}else if i.ID().(string) == "U"{
		score += computeUp(board)
	}else if i.ID().(string) == "R"{
		score += computeRight(board)
	}else if i.ID().(string) == "L"{
		score += computeLeft(board)
	}
	return g2048{board: board, score: g.score + score}
}

type coordinate struct {
	x int
	y int
}

func getFreePlaces(board [][]int)[]coordinate{
	freePlaces := make([]coordinate, 0)
	for y := 3 ; y >= 0 ; y-- {
		for x := 0 ; x < 4 ; x++ {
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

func addNumberOnBoard(board [][]int){
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

func addNumberOnBoardCord(cord coordinate, board [][]int){
	fRand := rand.Float64()
	val := 2
	if fRand >= 0.9 {
		val = 4
	}
	board[cord.x][cord.y] = val
}

func computeDown(board [][]int)int{
	score := 0
	for y := 3 ; y >= 0 ; y-- {
		for x := 0 ; x < 4 ; x++ {
			score += executeCompute("D", x, y, board)
		}
	}
	return score
}

func computeUp(board [][]int)int{
	score := 0
	for y := 0 ; y < 4 ; y++ {
		for x := 0 ; x < 4 ; x++ {
			score += executeCompute("U", x, y, board)
		}
	}
	return score
}

func computeRight(board [][]int)int{
	score := 0
	for y := 0 ; y < 4 ; y++ {
		for x := 3 ; x >= 0 ; x-- {
			score += executeCompute("R", x, y, board)
		}
	}
	return score
}

func computeLeft(board [][]int)int{
	score := 0
	for y := 0 ; y < 4 ; y++ {
		for x := 0 ; x < 4 ; x++ {
			score += executeCompute("L", x, y, board)
		}
	}
	return score
}

func executeCompute(kind string, x,y int, board [][]int)int {
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
	}else {
		board[newPos.x][newPos.y] += posVal
	}
	board[x][y] = 0
	return 2 * newPosVal
}

func copy2DArr(src [][]int)[][]int{
	duplicate := make([][]int, len(src))
	for i := range src {
		duplicate[i] = make([]int, len(src[i]))
		copy(duplicate[i], src[i])
	}
	return duplicate
}

func startNewGame()g2048 {
	board := make([][]int, 4)
	board[0] = make([]int, 4)
	board[1] = make([]int, 4)
	board[2] = make([]int, 4)
	board[3] = make([]int, 4)
	return g2048{
		board: board,
	}
}