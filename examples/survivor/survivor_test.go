package survivor

import (
	"fmt"
	"github.com/danielsussa/mcts"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSimpleActions(t *testing.T) {
	game := NewGame()
	game.Player.playAction(run)
	game.playTurn()
	fmt.Println(game.Score())
}

func TestSurvivor(t *testing.T) {
	rand.Seed(1)
	game := NewGame()

	for {
		tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 512, SimulationConfig: mcts.SimulationConfig{
			Ratio:    10,
			Strategy: mcts.Avg,
		}})
		nodes, err := tree.Start(game)
		assert.NoError(t, err)

		if len(nodes.NodeScore) == 0 {
			break
		}
		newGame := nodes.NodeScore[0].State.(Game)
		action := newGame.Player.SelectedAction
		game.Player.playAction(action)
		if !game.playTurn() {
			break
		}
		fmt.Println(fmt.Sprintf("[%v] %d action played: %s | place: %s | hunger: %d | life: %d | bag: %v | score: %0.f",
			game.CurrDate.Format("02 Jan 2006 15:04"),
			game.Turn,
			action,
			game.Player.CurrentPlace,
			game.Player.Hunger,
			game.Player.Life,
			game.Player.Items,
			game.Score(),
		))
	}

	fmt.Println("----------------")
	fmt.Println("Score: ", game.Score())
}

func TestSurvivorUnit1(t *testing.T) {
	rand.Seed(1)
	// [21 Jan 2000 23:00] 148 action played: run | place: forest | hunger: 61 | life: -10 | bag: [vegetable rabbit] | score: 279
	game := Game{
		CurrDate: startDate,
		Turn:     0,
		Player: &Player{
			Life:           10,
			Hunger:         55,
			SelectedAction: huntRabbit,
			Items:          []item{vegetable, rabbit},
			CurrentPlace:   forest,
		},
	}

	tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 256})
	nodes, err := tree.Start(game)
	assert.NoError(t, err)

	newGame := nodes.NodeScore[0].State.(Game)
	action := newGame.Player.SelectedAction

	assert.Equal(t, action, goToHome)

	fmt.Println("----------------")
	fmt.Println("Score: ", game.Score())
}

func TestSurvivorUnit2(t *testing.T) {
	rand.Seed(1)
	game := Game{
		CurrDate: startDate,
		Turn:     0,
		Player: &Player{
			SelectedAction: doNothing,
			Life:           100,
			Hunger:         107,
			Items:          []item{},
			CurrentPlace:   home,
		},
	}

	tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 512, SimulationConfig: mcts.SimulationConfig{
		Ratio:    100,
		Strategy: mcts.Avg,
	}})
	nodes, err := tree.Start(game)
	assert.NoError(t, err)

	newGame := nodes.NodeScore[0].State.(Game)
	action := newGame.Player.SelectedAction

	assert.Equal(t, goToForest, action)

	//fmt.Println("----------------")
	//fmt.Println("Score: ", game.Score())
}

func TestSurvivorUnitTmp(t *testing.T) {
	rand.Seed(1)
	game := Game{
		CurrDate: startDate,
		Turn:     0,
		Player: &Player{
			Life:         100,
			Hunger:       107,
			Items:        []item{},
			CurrentPlace: home,
		},
	}

	{
		assert.True(t, game.Play(rest))
	}
	{
		assert.True(t, game.Play(goToForest))
	}
	{
		assert.True(t, game.Play(huntRabbit))
	}
	{
		assert.True(t, game.Play(goToHome))
	}
	{
		assert.True(t, game.Play(useItem))
	}
	{
		assert.True(t, game.Play(doNothing))
	}
	{
		assert.True(t, game.Play(doNothing))
	}
	fmt.Println("score: ", game.Score())

}
