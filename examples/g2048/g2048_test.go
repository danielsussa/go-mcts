package g2048

import (
	"github.com/danielsussa/mcts"
	"testing"
)

func TestG2048(t *testing.T) {
	game := startNewGame()
	mcts.NewMonteCarloTree()
}
