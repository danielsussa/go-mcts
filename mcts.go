package mcts

import (
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

	isLeaf   bool
	maxPlays   *uint
	totalPlays uint
}

func (n *Node) getMaxPlays() uint{
	return *n.maxPlays
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

func (n *Node) expand() *Node {
	if n.maxPlays == nil {
		maxPlays := n.state.MaxPlays()
		n.maxPlays = &maxPlays
	}
	if *n.maxPlays == n.totalPlays {
		return n
	}
	state := n.state.Expand(n.totalPlays)
	n.totalPlays++
	if state == nil {
		n.isLeaf = true
		return nil
	}

	child := &Node{
		state:  state,
		parent: n,
		levelY: n.levelY + 1,
	}
	n.child = append(n.child, child)
	return child
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

func(n *Node) selection(policy PolicyFunc) *Node {
	for {
		if n.child == nil {
			return n
		}
		if n.maxPlays == nil || *n.maxPlays != n.totalPlays {
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
	Expand(idx uint)State
	MaxPlays()uint
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
		childNode := node.expand()
		if childNode == nil {
			continue
		}
		childNode.rollOut()

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
		return ndScore[i].score > ndScore[j].score
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