package survivor

import (
	"github.com/danielsussa/mcts"
)

// simulate (player action and turn)
func (g Game) Simulate() float64 {
	start := g.CurrDate
	for g.PlayRandom() {
	}
	end := g.CurrDate
	score := end.Sub(start).Hours() / 1000
	return score
}

func (g Game) Score() float64 {
	return g.CurrDate.Sub(startDate).Hours() / 100
}

func (g Game) Iterations() []interface{} {
	listActions := g.Player.getPossibleActions()

	iters := make([]interface{}, 0)
	for _, act := range listActions {
		iters = append(iters, act)
	}

	return iters
}

func (g Game) ID() string {
	return string(g.Player.SelectedAction)
}

func (g Game) Expand(i interface{}) mcts.State {
	action := i.(action)
	g.Play(action)
	return g
}

func (g Game) Copy() mcts.State {
	itemsCopy := make([]item, len(g.Player.Items))
	copy(itemsCopy, g.Player.Items)
	return Game{
		CurrDate: g.CurrDate,
		Turn:     g.Turn,
		Player: &Player{
			Life:           g.Player.Life,
			Hunger:         g.Player.Hunger,
			SelectedAction: g.Player.SelectedAction,
			Items:          itemsCopy,
			CurrentPlace:   g.Player.CurrentPlace,
		},
	}
}
