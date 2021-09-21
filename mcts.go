package mcts

import (
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

	iterations       []interface{}
	currIterationIdx int
	id               string
}

func (n *Node) rollOut(simConfig SimulationConfig) {
	score := 0.0
	switch simConfig.Strategy {
	case Avg:
		score = n.avgStrategy(simConfig)
	case Min:
		score = n.minStrategy(simConfig)
	case Max:
		score = n.maxStrategy(simConfig)
	default:
		score = n.avgStrategy(simConfig)
	}
	n.backPropagate(score)
}

func (n *Node) avgStrategy(simConfig SimulationConfig) float64 {
	score := 0.0
	for i := 0; i <= simConfig.Ratio; i++ {
		score += n.state.Copy().Simulate()
	}
	return score
}

func (n *Node) minStrategy(simConfig SimulationConfig) float64 {
	minScore := 0.0
	for i := 0; i <= simConfig.Ratio; i++ {
		if i == 0 {
			minScore = n.state.Copy().Simulate()
			if minScore <= 0 {
				break
			}
		} else {
			currScore := n.state.Copy().Simulate()
			if currScore < minScore {
				minScore = currScore
			}
		}
	}
	return minScore
}

func (n *Node) maxStrategy(simConfig SimulationConfig) float64 {
	maxScore := 0.0
	for i := 0; i <= simConfig.Ratio; i++ {
		if i == 0 {
			maxScore = n.state.Copy().Simulate()
		} else {
			currScore := n.state.Copy().Simulate()
			if currScore > maxScore {
				maxScore = currScore
			}
		}
	}
	return maxScore
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
	if n.iterations == nil {
		iteration := n.state.Copy().Iterations()
		if iteration == nil {
			return &Node{}
		}
		n.iterations = n.state.Copy().Iterations()
	}
	if len(n.iterations) == n.currIterationIdx {
		return n
	}
	state := n.state.Copy().Expand(n.iterations[n.currIterationIdx])
	n.currIterationIdx++
	if state == nil {
		return nil
	}

	child := &Node{
		state:      state,
		parent:     n,
		iterations: nil,
		levelY:     n.levelY + 1,
	}
	n.child = append(n.child, child)
	return child
}

func (n *Node) getParentNVisited() uint {
	if n.parent == nil {
		return n.nVisited
	}
	return n.parent.getParentNVisited()
}

func (n *Node) selection(policy PolicyFunc) *Node {
	if n.child == nil {
		return n
	}
	if n.iterations == nil || n.currIterationIdx < len(n.iterations) {
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
	Simulate() float64
	Expand(iter interface{}) State
	Iterations() []interface{}
	Copy() State
}

type MonteCarloTree struct {
	policy            PolicyFunc
	node              *Node
	maxInteractions   uint
	totalInteractions uint
	simulationsConfig SimulationConfig
	firstStateID      string
}

type FinalScore struct {
	Iterations uint
	NodeScore  []nodeFinalScore
	TotalNodes uint
}

type nodeFinalScore struct {
	State State
	total uint
}

func (mct *MonteCarloTree) Start(initialState State) (FinalScore, error) {
	mct.node = &Node{
		state:      initialState.Copy(),
		iterations: nil,
	}
	return mct.start()
}

func (mct *MonteCarloTree) start() (FinalScore, error) {
	interactions := uint(0)
	totalNodes := uint(0)
	for {
		node := mct.node.selection(mct.policy)

		childNode := node.expand()

		if childNode == nil {
			node.rollOut(mct.simulationsConfig)
		} else {
			childNode.rollOut(mct.simulationsConfig)
			totalNodes++
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
		TotalNodes: totalNodes,
		NodeScore:  ndScore,
	}, nil
}

type MonteCarloTreeConfig struct {
	MaxTimeout       *time.Duration
	MaxIterations    uint
	SimulationConfig SimulationConfig
}

type SimulationConfig struct {
	Ratio    int
	Strategy ScoreStrategy
}

type ScoreStrategy string

const (
	Min ScoreStrategy = "min"
	Max ScoreStrategy = "max"
	Avg ScoreStrategy = "avg"
)

func NewMonteCarloTree(config MonteCarloTreeConfig) MonteCarloTree {
	if config.MaxIterations == 0 && config.MaxTimeout == nil {
		config.MaxIterations = 1000
	}
	return MonteCarloTree{
		policy:            defaultPolicyFunc(),
		maxInteractions:   config.MaxIterations,
		simulationsConfig: config.SimulationConfig,
	}
}
