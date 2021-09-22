package survivor

import (
	"math/rand"
	"sort"
	"time"
)

type Game struct {
	CurrDate time.Time
	Turn     int
	Player   *Player
}

var startDate = time.Date(2000, 1, 10, 8, 0, 0, 0, time.UTC)

func NewGame() Game {
	return Game{
		CurrDate: startDate,
		Turn:     0,
		Player: &Player{
			Life:         100,
			Hunger:       0,
			Items:        []item{},
			CurrentPlace: home,
		},
	}
}

func (g *Game) PlayRandom() bool {
	g.Player.playRandomAction()
	return g.playTurn()
}

func (g *Game) Play(a action) bool {
	actions := g.Player.getPossibleActions()
	if !canPlay(a, actions) {
		panic("cannot play action")
	}
	g.Player.playAction(a)

	return g.playTurn()
}

func canPlay(a action, l []action) bool {
	for _, aa := range l {
		if aa == a {
			return true
		}
	}
	return false
}

func (g *Game) playTurn() bool {
	g.Turn++
	player := g.Player
	if player.Hunger > 100 {
		player.Life -= player.Hunger - 100
	}
	if player.Life <= 0 {
		return false
	}

	player.Hunger++

	switch g.Player.SelectedAction {
	case rest:
		g.CurrDate = g.CurrDate.Add(6 * time.Hour)
	case doNothing:
		g.CurrDate = g.CurrDate.Add(1 * time.Hour)
	case goToFarm:
		g.CurrDate = g.CurrDate.Add(3 * time.Hour)
	case goToForest:
		g.CurrDate = g.CurrDate.Add(4 * time.Hour)
	case huntRabbit:
		g.CurrDate = g.CurrDate.Add(2 * time.Hour)
	case getVegetable:
		g.CurrDate = g.CurrDate.Add(20 * time.Minute)
	}

	damageFactor := 1
	if g.CurrDate.Hour() > 21 || g.CurrDate.Hour() < 6 {
		damageFactor = 2
	}
	switch g.Player.CurrentPlace {
	case forest:
		// throw dice to take damage
		if rand.Intn(100) > 65 {
			g.Player.Life -= 20 * damageFactor
		}
	case farm:
		// throw dice to take damage
		if rand.Intn(100) > 95 {
			g.Player.Life -= 20 * damageFactor
		}
	}
	return true
}

type Player struct {
	Life   int
	Hunger int

	SelectedAction action
	Items          []item
	CurrentPlace   place
}

type action string

const (
	run       action = "run"
	doNothing action = "doNothing"

	goToFarm   action = "goToFarm"
	goToHome   action = "goToHome"
	goToForest action = "goToForest"

	getVegetable action = "getVegetable"
	huntRabbit   action = "huntRabbit"
	useItem      action = "useItem"
	rest         action = "rest"
)

func (p *Player) getPossibleActions() []action {
	mapActions := map[action]bool{
		run:        true,
		doNothing:  true,
		goToFarm:   true,
		goToHome:   true,
		goToForest: true,

		getVegetable: false,
		useItem:      false,
		rest:         false,
		huntRabbit:   false,
	}

	if p.CurrentPlace == home {
		mapActions[goToHome] = false
		mapActions[rest] = true
	}

	if p.CurrentPlace == forest {
		mapActions[goToForest] = false
		mapActions[huntRabbit] = true
	}

	if p.CurrentPlace == farm {
		mapActions[getVegetable] = true
		mapActions[goToFarm] = false
	}

	if len(p.Items) >= 5 {
		mapActions[getVegetable] = false
	}

	if len(p.Items) > 0 {
		mapActions[useItem] = true
	}

	var listActions []action
	for act, canDo := range mapActions {
		if canDo {
			listActions = append(listActions, act)
		}
	}
	sort.Slice(listActions, func(i, j int) bool {
		return listActions[i] < listActions[j]
	})
	return listActions
}

func (p *Player) playRandomAction() action {
	actions := p.getPossibleActions()
	idx := rand.Intn(len(actions))
	selectedAction := actions[idx]
	p.playAction(selectedAction)
	return selectedAction
}

func (p *Player) playAction(a action) {
	switch a {
	case doNothing:
		p.doNothing()
	case run:
		p.run()
	case goToFarm:
		p.goToFarm()
	case goToHome:
		p.goToHome()
	case goToForest:
		p.goToForest()
	case huntRabbit:
		p.huntRabbit()
	case getVegetable:
		p.getVegetable()
	case useItem:
		p.useItem()
	case rest:
		p.rest()
	}
}

// ACTIONS

func (p *Player) doNothing() {
	p.setAction(doNothing)
}

func (p *Player) run() {
	p.setAction(run)
	p.Hunger += 5
}

func (p *Player) goToFarm() {
	p.setAction(goToFarm)
	p.CurrentPlace = farm
	p.Hunger += 5
}

func (p *Player) goToHome() {
	p.setAction(goToHome)
	p.CurrentPlace = home
	p.Hunger += 5
}

func (p *Player) goToForest() {
	p.setAction(goToForest)
	p.CurrentPlace = forest
	p.Hunger += 8
}

func (p *Player) getVegetable() {
	p.setAction(getVegetable)

	p.Items = append(p.Items, vegetable)
}

func (p *Player) huntRabbit() {
	p.setAction(huntRabbit)

	p.Items = append(p.Items, rabbit)
}

func (p *Player) useItem() {
	p.setAction(useItem)

	item := p.Items[0]
	p.Items = p.Items[1:]

	switch item {
	case vegetable:
		p.Hunger -= 10
	case rabbit:
		p.Hunger -= 25
	}
	if p.Hunger < 0 {
		p.Hunger = 0
	}
}

func (p *Player) rest() {
	p.setAction(rest)
	p.Life += 15
	p.Hunger += 5
	if p.Life > 100 {
		p.Life = 100
	}
}

func (p *Player) setAction(a action) {
	p.SelectedAction = a
}

// to String

func (p *Player) allActionsAsString() string {
	k := ""
	for _, act := range p.SelectedAction {
		k += string(act) + "_"
	}
	return k
}

// items

type item string

const (
	vegetable item = "vegetable"
	rabbit    item = "rabbit"
)

// place

type place string

const (
	home   place = "home"
	farm   place = "farm"
	forest place = "forest"
)
