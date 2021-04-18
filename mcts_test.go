package mcts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPolicyFunc(t *testing.T) {
	f := defaultPolicyFunc()

	assert.Equal(t, 0.79, f(-1, 1, 5))
	assert.Equal(t, 2.79, f(1, 1, 5))
	assert.Equal(t, 1.48, f(0, 2, 9))
	assert.Equal(t, 2.48, f(2, 2, 9))
	assert.Equal(t, 0.89, f(-1, 1, 6))
	assert.Equal(t, 2.52, f(2, 2, 10))
}

// precisa criar um otimo teste de selecao!!
func TestNodeSelection(t *testing.T) {
	maxPlays := uint(6)
	parent := &Node{
		score:            6,
		nVisited:         14,
		levelY:           0,
		state:            nil,
		child:            []*Node{},
		parent:           nil,
		iterations:       &maxPlays,
		currIterationIdx: 6,
	}

	c1 := &Node{
		score:            3,
		nVisited:         6,
		levelY:           1,
		parent:           parent,
		iterations:       nil,
		currIterationIdx: 0,
	}

	c2 := &Node{
		score:            3,
		nVisited:         5,
		levelY:           1,
		parent:           parent,
		iterations:       nil,
		currIterationIdx: 0,
	}

	c3 := &Node{
		score:            3,
		nVisited:         3,
		levelY:           1,
		parent:           parent,
		iterations:       nil,
		currIterationIdx: 0,
	}

	parent.child = append(parent.child, c1)
	parent.child = append(parent.child, c2)
	parent.child = append(parent.child, c3)

	selectedNode := parent.selection(defaultPolicyFunc())
	assert.Equal(t, selectedNode, c3)

}

func TestNodeSelection2(t *testing.T) {
	maxPlays := uint(6)
	parent := &Node{
		nVisited:         8,
		state:            nil,
		child:            []*Node{},
		parent:           nil,
		iterations:       &maxPlays,
		currIterationIdx: 6,
	}

	l1N1 := &Node{score: -1, nVisited: 1, parent: parent}
	l1N2 := &Node{score: 1, nVisited: 1, parent: parent}
	l1N3 := &Node{score: 0, nVisited: 2, parent: parent}
	l1N4 := &Node{score: 2, nVisited: 2, parent: parent}
	l1N5 := &Node{score: -1, nVisited: 1, parent: parent}
	l1N6 := &Node{score: 0, nVisited: 1, parent: parent}

	l2N1 := &Node{score: 1, nVisited: 1, parent: parent}
	l1N3.child = append(l1N3.child, l2N1)

	l2N2 := &Node{score: -1, nVisited: 1, parent: parent}
	l1N4.child = append(l1N4.child, l2N2)

	parent.child = append(parent.child, l1N1, l1N2, l1N3, l1N4, l1N5, l1N6)

	nodeScore := getNodeScore(parent.child,defaultPolicyFunc())

	assert.Equal(t, nodeScore[0].score, 1.44)

}