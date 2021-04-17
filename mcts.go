package mcts

import (
	"fmt"
	"math"
	"sort"
	"time"
)

type Node struct {
	score    float64
	nVisited uint
	levelY   int

	state  State
	child  []*Node
	parent *Node

	maxPlays   int
	totalPlays int
}

func (n *Node) rollOut() {
	result := n.state.Simulate()
	n.backPropagate(result)
}

func (n *Node) backPropagate(result SimulationResult) {
	n.nVisited++
	nodePlayer := n.state.Player()
	if nodePlayer == result.Player && nodePlayer == result.Winner {
		n.score += result.Score
	} else if nodePlayer == result.Player && nodePlayer != result.Winner {
		n.score -= result.Score
	} else if nodePlayer != result.Player && nodePlayer == result.Winner {
		n.score += result.Score
	} else if nodePlayer != result.Player && nodePlayer != result.Winner {
		n.score -= result.Score
	}
	if n.parent == nil {
		return
	}
	n.parent.backPropagate(result)
}

func (n *Node) expand() (*Node, bool) {
	if n.maxPlays == -1 {
		n.maxPlays = n.state.MaxPlays()
	}
	if n.maxPlays == n.totalPlays {
		return n, false
	}
	state := n.state.Expand(n.totalPlays)
	n.totalPlays++
	if state == nil {
		return nil, false
	}

	// is duplicated
	for _, child := range n.child {
		if child.state.ID() == state.ID() {
			return &Node{}, true
		}
	}

	child := &Node{
		state:    state,
		parent:   n,
		maxPlays: -1,
		levelY:   n.levelY + 1,
	}
	n.child = append(n.child, child)
	return child, false
}

func (n *Node) getParentNVisited() uint {
	if n.parent == nil {
		return n.nVisited
	}
	return n.parent.getParentNVisited()
}

func (n *Node) selectByState(state State) *Node {
	if n.state.ID() == state.ID() {
		return n
	}
	for _, child := range n.child {
		if child.state.ID() == state.ID() {
			return child
		}
		childNode := child.selectByState(state)
		if childNode == nil {
			continue
		}
		if childNode.state.ID() == state.ID() {
			return childNode
		}
	}
	return nil
}

func (n *Node) selection(policy PolicyFunc) *Node {
	for {
		if n.child == nil {
			return n
		}
		if n.maxPlays == -1 || n.totalPlays < n.maxPlays {
			return n
		}
		selectedNodes := getNodeScore(n.child, policy)
		sort.SliceStable(selectedNodes, func(i, j int) bool {
			if selectedNodes[i].score > selectedNodes[j].score {
				return true
			} else if selectedNodes[i].score < selectedNodes[j].score {
				return false
			} else {
				return selectedNodes[i].node.nVisited < selectedNodes[j].node.nVisited
			}
		})
		for _, selectedNode := range selectedNodes {
			node := selectedNode.node.selection(policy)
			if node == nil {
				continue
			}
			return node
		}
		return n
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

type PolicyFunc func(total float64, nVisited, NVisited uint) float64

func defaultPolicyFunc() PolicyFunc {
	return func(total float64, nVisited, NVisited uint) float64 {
		exploitation := total / float64(nVisited)
		exploration := math.Sqrt(2 * math.Log(float64(NVisited)) / float64(nVisited))
		sum := exploitation + exploration
		return math.Round(sum*100) / 100
	}
}

type State interface {
	Simulate() SimulationResult
	Expand(idx int) State
	MaxPlays() int
	Player() string
	ID() string
}

type SimulationResult struct {
	Score  float64
	Winner string
	Player string
}

type MonteCarloTree struct {
	policy            PolicyFunc
	node              *Node
	maxInteractions   uint
	totalInteractions uint
}

type FinalScore struct {
	Iterations uint
	NodeScore  []nodeFinalScore
}

type nodeFinalScore struct {
	State State
	total uint
}

func (mct *MonteCarloTree) Continue(state State) (FinalScore, error) {
	node := mct.node.selectByState(state)
	if node == nil {
		return FinalScore{}, fmt.Errorf("node not found")
	}
	node.parent = nil
	mct.node = node
	return mct.start()
}

func (mct *MonteCarloTree) Start(initialState State) (FinalScore, error) {
	mct.node = &Node{
		state:    initialState,
		maxPlays: -1,
	}
	return mct.start()
}

func (mct *MonteCarloTree) start() (FinalScore, error) {
	interactions := uint(0)
	for {
		node := mct.node.selection(mct.policy)

		childNode, duplicatedNode := node.expand()
		if duplicatedNode {
			return FinalScore{}, fmt.Errorf("duplication node")
		}

		if childNode == nil {
			node.rollOut()
		} else {
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
			total: childNode.nVisited,
			State: childNode.state,
		})
	}

	sort.SliceStable(ndScore, func(i, j int) bool {
		return ndScore[i].total > ndScore[j].total
	})

	return FinalScore{
		Iterations: mct.totalInteractions,
		NodeScore:  ndScore,
	}, nil
}

type MonteCarloTreeConfig struct {
	MaxTimeout    *time.Duration
	MaxIterations uint
}

func NewMonteCarloTree(config MonteCarloTreeConfig) MonteCarloTree {
	if config.MaxIterations == 0 || config.MaxTimeout == nil {
		config.MaxIterations = 1000
	}
	return MonteCarloTree{
		policy:          defaultPolicyFunc(),
		maxInteractions: config.MaxIterations,
	}
}
