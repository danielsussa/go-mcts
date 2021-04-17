package mcts

import (
	"fmt"
	"math"
	"sort"
	"time"
)

type Node struct {
	score    float64
	nVisited int64
	levelY   int

	state  State
	child  []*Node // if nil it's a leaf node
	parent *Node // if nil it's a root node

	maxPlays   *uint
	totalPlays uint
	player     string
}

func (n *Node) getMaxPlays() uint{
	return *n.maxPlays
}

func (n *Node) rollOut() {
	score, winner := n.state.Simulate()
	n.backPropagate(score, winner)
}

func (n *Node) backPropagate(score float64, winner string) {
	n.nVisited++
	if n.player == winner {
		n.score += score
	}else {
		n.score -= score
	}
	if n.parent == nil {
		return
	}
	n.parent.backPropagate(score, winner)
}

func (n *Node) expand() (*Node,bool) {
	if n.maxPlays == nil {
		maxPlays := n.state.MaxPlays()
		n.maxPlays = &maxPlays
	}
	if *n.maxPlays == n.totalPlays {
		return n, true
	}
	state := n.state.Expand(n.totalPlays)
	n.totalPlays++
	if state == nil {
		return nil, true
	}

	// is duplicated
	for _, child := range n.child {
		if child.state.ID() == state.ID() {
			return &Node{}, false
		}
	}

	child := &Node{
		state:  state,
		player: state.Player(),
		parent: n,
		levelY: n.levelY + 1,
	}
	n.child = append(n.child, child)
	return child, true
}

func(n *Node) getParentNVisited() int64 {
	if n.parent == nil {
		return n.nVisited
	}
	return n.parent.getParentNVisited()
}

func(n *Node) selectByState(state State) *Node {
	if n.state.ID() == state.ID() {
		return n
	}
	for _, child := range n.child {
		if child.state.ID() == state.ID() {
			return child
		}
		childNode :=  child.selectByState(state)
		if childNode == nil {
			continue
		}
		if childNode.state.ID() == state.ID() {
			return childNode
		}
	}
	return nil
}

type action string

const (
	selection
)

func(n *Node) selection(policy PolicyFunc) *Node {
	for {
		if n.child == nil {
			return n
		}
		if n.maxPlays == nil || *n.maxPlays != n.totalPlays {
			return n
		}
		selectedNodes := getNodeScore(n.child, policy)
		sort.SliceStable(selectedNodes, func(i, j int) bool {
			return selectedNodes[i].score > selectedNodes[j].score && selectedNodes[i].node.nVisited < selectedNodes[j].node.nVisited
		})
		for _, selectedNode := range selectedNodes {
			node := selectedNode.node.selection(policy)
			if node == nil {
				continue
			}
			return node
		}
		return nil
	}
}

type nodeScore struct {
	node  *Node
	score float64
}

func getNodeScore(childNodes []*Node, policy PolicyFunc) []nodeScore {
	nodesScore := make([]nodeScore, 0)

	for _, child := range childNodes {
		nodesScore = append(nodesScore, nodeScore{
			node:  child,
			score: policy(child.score, child.nVisited, child.parent.nVisited),
		})
	}
	return nodesScore
}

type PolicyFunc func(total float64, nVisited, NVisited int64)float64

func defaultPolicyFunc() PolicyFunc {
	return func(total float64, nVisited , NVisited int64) float64 {
		exploitation := total / float64(nVisited)
		exploration := math.Sqrt(2 * math.Log(float64(NVisited)) / float64(nVisited))
		sum := exploitation + exploration
		return math.Round(sum * 100) / 100

		//nVisitedF := float64(nVisited)
		//sqrtV := math.Sqrt(math.Log(float64(NVisited)) / nVisitedF)
		//
		//return math.Round((total + 2.0 * sqrtV) * 100) / 100
	}
}

//let exploitation = node.data.value / node.data.simulations;
//let exploration = Math.sqrt(2 * Math.log(parent.data.simulations) / node.data.simulations);
//return exploitation + exploration;
//

type State interface {
	Simulate()(float64, string)
	Expand(idx uint)State
	MaxPlays()uint
	Player()string
	ID()string
}

type MonteCarloTree struct {
	policy    PolicyFunc
	node      *Node
	maxInteractions uint
	totalInteractions uint
}

type FinalScore struct {
	Iterations uint
	UntilEnd   bool
	NodeScore  []nodeFinalScore
}

type nodeFinalScore struct {
	State  State
	score  float64
}

func(mct *MonteCarloTree) Continue(state State)(FinalScore, error){
	node := mct.node.selectByState(state)
	if node == nil {
		return FinalScore{}, fmt.Errorf("node not found")
	}
	node.parent = nil
	mct.node = node
	return mct.start()
}

func(mct *MonteCarloTree) Start(initialState State)(FinalScore, error){
	mct.node = &Node{
		state:  initialState,
	}
	return mct.start()
}

func(mct *MonteCarloTree) start()(FinalScore, error){
	iterateUntilEnd := false
	interactions := uint(0)
	for {
		node := mct.node.selection(mct.policy)
		if node == nil {
			iterateUntilEnd = true
			break
		}
		childNode, addNewNode := node.expand()
		if !addNewNode {
			return FinalScore{}, fmt.Errorf("duplication node")
		}

		if childNode == nil {
			node.rollOut()
		}else {
			childNode.rollOut()
		}


		interactions++
		if interactions >= mct.maxInteractions {
			break
		}
	}
	mct.totalInteractions += interactions

	ndScore := make([]nodeFinalScore, 0)
	for _, childNode := range mct.node.child {
		ndScore = append(ndScore, nodeFinalScore{
			score:  float64(childNode.nVisited),
			State:  childNode.state,
		})
	}

	sort.SliceStable(ndScore, func(i, j int) bool {
		return ndScore[i].score > ndScore[j].score
	})

	return FinalScore{
		Iterations: mct.totalInteractions,
		UntilEnd:   iterateUntilEnd,
		NodeScore:  ndScore,
	}, nil
}

type MonteCarloTreeConfig struct {
	MaxTimeout    time.Duration
	MaxIterations uint
}

func NewMonteCarloTree(config MonteCarloTreeConfig) MonteCarloTree {
	if config.MaxIterations == 0 {
		config.MaxIterations = 1000
	}
	return MonteCarloTree{
		policy: defaultPolicyFunc(),
		maxInteractions: config.MaxIterations,
	}
}