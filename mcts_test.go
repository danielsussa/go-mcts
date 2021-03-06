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

func TestNodeSelection(t *testing.T) {
	parent := &Node{
		score:            6,
		nVisited:         14,
		levelY:           0,
		state:            nil,
		child:            []*Node{},
		parent:           nil,
		iterations:       []interface{}{1, 2, 3, 4, 5, 6},
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
	{
		parent := &Node{nVisited: 8, child: []*Node{}}

		l1N1 := &Node{score: -1, nVisited: 1, parent: parent}
		l1N2 := &Node{score: 0, nVisited: 3, parent: parent}
		l1N3 := &Node{score: -2, nVisited: 2, parent: parent}
		l1N4 := &Node{score: -2, nVisited: 2, parent: parent}

		parent.child = append(parent.child, l1N1, l1N2, l1N3, l1N4)

		nodeScore := getNodeScore(parent, defaultPolicyFunc())

		assert.Equal(t, nodeScore[0].score, 1.18)
	}
	{
		parent := &Node{nVisited: 9, child: []*Node{}}

		l1N1 := &Node{score: -1, nVisited: 1, parent: parent}
		l1N2 := &Node{score: 0, nVisited: 4, parent: parent}
		l1N3 := &Node{score: -2, nVisited: 2, parent: parent}
		l1N4 := &Node{score: -2, nVisited: 2, parent: parent}

		parent.child = append(parent.child, l1N1, l1N2, l1N3, l1N4)

		nodeScore := getNodeScore(parent, defaultPolicyFunc())

		assert.Equal(t, nodeScore[0].score, 1.10)
	}
	{
		parent := &Node{nVisited: 10, child: []*Node{}}

		l1N1 := &Node{score: -2, nVisited: 2, parent: parent}
		l1N2 := &Node{score: 0, nVisited: 4, parent: parent}
		l1N3 := &Node{score: -2, nVisited: 2, parent: parent}
		l1N4 := &Node{score: -2, nVisited: 2, parent: parent}

		parent.child = append(parent.child, l1N1, l1N2, l1N3, l1N4)

		nodeScore := getNodeScore(parent, defaultPolicyFunc())

		assert.Equal(t, nodeScore[0].score, 1.07)
	}
	{
		parent := &Node{nVisited: 11, child: []*Node{}}

		l1N1 := &Node{score: -2, nVisited: 2, parent: parent}
		l1N2 := &Node{score: 1, nVisited: 5, parent: parent}
		l1N3 := &Node{score: -2, nVisited: 2, parent: parent}
		l1N4 := &Node{score: -2, nVisited: 2, parent: parent}

		parent.child = append(parent.child, l1N1, l1N2, l1N3, l1N4)

		nodeScore := getNodeScore(parent, defaultPolicyFunc())

		assert.Equal(t, nodeScore[0].score, 1.18)
	}

}

func TestNormalize(t *testing.T) {
	assert.Equal(t, 0.5, normalize(6, 3, 9))
}
