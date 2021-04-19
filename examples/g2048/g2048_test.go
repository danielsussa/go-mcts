package g2048

import (
	"github.com/danielsussa/mcts"
	"github.com/stretchr/testify/assert"
	"testing"
)


// 4 0 0 0
// 4 2 0 0
// 4 0 0 0
// 4 2 0 0
func TestGetAllIterations(t *testing.T) {
	g2048 := startNewGame()
	g2048.board[0][0] = 4
	g2048.board[0][1] = 4
	g2048.board[0][2] = 4
	g2048.board[0][3] = 4
	g2048.board[1][1] = 2
	g2048.board[1][3] = 2
	assert.Equal(t, len(getAllIterations(g2048.board)), 3)
}

func TestG2048(t *testing.T) {
	game2048 := startNewGame()
	addNumberOnBoard(game2048.board)

	totalIterations := 0

	for {
		//print2048(game2048.board, game2048.score)

		tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 200})
		nodes, err := tree.Start(game2048)
		assert.NoError(t, err)

		if len(nodes.NodeScore) == 0 {
			break
		}

		if totalIterations % 200 == 0 {
			print2048(game2048.board, game2048.score)
		}

		game2048 = nodes.NodeScore[0].State.(g2048)
		//print2048(game2048.board, game2048.score)
		// add new move
		addNumberOnBoard(game2048.board)
		totalIterations++
	}
	print2048(game2048.board, game2048.score)


}
