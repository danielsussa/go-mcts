package mcts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPolicyFunc(t *testing.T) {
	f := defaultPolicyFunc()
	assert.Equal(t, 21.67, f(20, 1, 2))
	assert.Equal(t, 11.67, f(10, 1, 2))
}
