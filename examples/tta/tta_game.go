package tta

import "github.com/danielsussa/mcts"

type ttaGame struct {

	civilCards []string
	players []player

	currentPlayer player
}

type currentPlayer struct {
	isDirty       bool
	actions       []action
	player        player
}

func (g ttaGame) Iterations() []mcts.Iteration{
	return nil
}


func (g ttaGame) Copy() mcts.State{
	civilCopy := make([]string,len(g.civilCards))
	copy(civilCopy, g.civilCards)
	return ttaGame{
		civilCards: civilCopy,
		players:    g.players,
	}
}

func (g ttaGame) Expand(iter mcts.Iteration) mcts.State{
	return nil
}

func (g ttaGame) Simulate()mcts.SimulationResult{
	game := g.copy()
	//fpsPlayer := game.currentPlayer

	for {
		game.simulateNextAction()
		break
	}

	return mcts.SimulationResult{}
}


func (g ttaGame) ID() string{
	return ""
}

func (g ttaGame) getAllActions(p player) []action {
	return nil
}


func (g ttaGame) getNextPlayer() player {
	return player{}
}

func (g ttaGame) copy() ttaGame{
	civilCopy := make([]string,len(g.civilCards))
	copy(civilCopy, g.civilCards)
	return ttaGame{
		civilCards: civilCopy,
		players:    g.players,
	}
}

func (g *ttaGame) updatePlayerRef(p player) {
	g.currentPlayer = p
	for idx, player := range g.players {
		if player.ID == p.ID {
			g.players[idx] = p
			break
		}
	}
}

func (g ttaGame) simulateNextAction() {

}

type player struct {
	ID      string
	science int
	rocks   int
	food    int
	culture int

	totalActions int
	remainAction int

	civilCards []string
}

// cada expand e uma jogada
// validMove para testar

type action struct {
	kind actionKind
	idx  int
}

type actionKind string

const (
	getCard actionKind = "GET_CARD"
)


func (p player) excludeAction() player {
	return player{}
}

func (p player) actionValidation(act action) bool {
	return false
}

func (p player) canGetCard(idx int, card string) bool {
	actionsRequired := 3
	if idx < 8 {
		actionsRequired = 2
	}
	if idx < 4 {
		actionsRequired = 1
	}
	if p.remainAction < actionsRequired {
		return false
	}
	// validate if can get card
	return true
}

func (p player) playAction(act action) player {
	return player{}
}

func (p player) copy() player{
	civilCopy := make([]string,len(p.civilCards))
	copy(civilCopy, p.civilCards)

	pCopy := p
	pCopy.civilCards = civilCopy
	return pCopy
}

func (p *player) getCard(card string)  {
	p.civilCards = append(p.civilCards, card)
}


const (
	res1 string = "RS1"
)