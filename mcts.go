package mcts

import (
	"math"
	"sort"
	"time"
)

type Node struct {
	score    float64
	scoreL   float64
	leafs    int64
	nVisited int64
	levelY   int
	levelX   int

	state  State
	child  []*Node // if nil it's a leaf node
	parent *Node // if nil it's a root node

	isLeaf bool
}


func (n *Node) rollOut() {
	score := n.state.Simulate()
	n.backPropagate(score)
}

func (n *Node) backPropagate(score float64) {
	n.nVisited++
	n.score += score
	if n.parent == nil {
		return
	}
	n.parent.backPropagate(score)
}

func (n *Node) backPropagateL(score float64) {
	n.scoreL += score
	n.leafs++
	if n.parent == nil {
		return
	}
	n.parent.backPropagateL(score)
}

func (n *Node) expand() bool{
	states := n.state.Expand()
	if states == nil || len(states) == 0 {
		n.isLeaf = true
		n.scoreL = n.state.Simulate()
		n.parent.backPropagateL(n.scoreL)
		return false
	}
	for idx, state := range states {
		n.child = append(n.child, &Node{
			state:  state,
			parent: n,
			levelY: n.levelY + 1,
			levelX: idx,
		})
	}
	return true
}

func(n *Node) getParentNVisited() int64 {
	if n.parent == nil {
		return n.nVisited
	}
	return n.parent.getParentNVisited()
}

func(n *Node) getByState(state State) *Node {
	if n.state.ID() == state.ID() {
		return n
	}
	for _, child := range n.child {
		node := child.getByState(state)
		if node == nil {
			continue
		}
		return node
	}
	return nil
}

func(n *Node) finalScore() float64 {
	exploitation := n.score / float64(n.nVisited)
	exploration := math.Sqrt(math.Log(float64(n.getParentNVisited())) / float64(n.nVisited))
	return exploitation + exploration
}
func(n *Node) selection(policy PolicyFunc) *Node {
	for {
		if n.child == nil {
			return n
		}
		selectedNodes := getNodeScore(n.child, policy)
		if len(selectedNodes) == 0 {
			return nil
		}
		sort.SliceStable(selectedNodes, func(i, j int) bool {
			return selectedNodes[i].score > selectedNodes[j].score
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
		if child.isLeaf {
			continue
		}
		nodesScore = append(nodesScore, nodeScore{
			node:  child,
			score: policy(child.score, child.nVisited, child.getParentNVisited()),
		})
	}
	return nodesScore
}

type PolicyFunc func(total float64, nVisited, NVisited int64)float64

func defaultPolicyFunc() PolicyFunc {
	return func(total float64, nVisited , NVisited int64) float64 {
		if nVisited == 0 {
			return math.Inf(1)
		}
		nVisitedF := float64(nVisited)
		sqrtV := math.Sqrt(math.Log(float64(NVisited)) / nVisitedF)
		return math.Round((total + 2.0 * sqrtV) * 100) / 100
	}
}

type State interface {
	Simulate()float64
	Expand()[]State
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
	node   *Node
	State  State
	score  float64
}

func(mct *MonteCarloTree) Continue(state State)(FinalScore, bool){
	node := mct.node.getByState(state)
	if node == nil {
		return FinalScore{}, false
	}
	mct.node = node
	return mct.start(), true
}

func(mct *MonteCarloTree) Start(initialState State)FinalScore{
	mct.node = &Node{
		state:  initialState,
	}
	return mct.start()
}

func(mct *MonteCarloTree) start()FinalScore{
	iterateUntilEnd := false
	interactions := uint(0)
	for {
		node := mct.node.selection(mct.policy)
		if node == nil {
			iterateUntilEnd = true
			break
		}
		if node.parent == nil && node.child == nil {
			node.expand()
		}else if node.nVisited == 0 {
			node.rollOut()
		}else if node.nVisited == 1{
			node.expand()
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
			node:   childNode,
			score:  defaultPolicyFunc()(childNode.score, childNode.nVisited, childNode.getParentNVisited()),
			State:  childNode.state,
		})
	}

	sort.SliceStable(ndScore, func(i, j int) bool {
		return ndScore[i].node.nVisited > ndScore[j].node.nVisited
	})

	return FinalScore{
		Iterations: mct.totalInteractions,
		UntilEnd:   iterateUntilEnd,
		NodeScore:  ndScore,
	}
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