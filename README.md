# Monte Carlo Tree Search

Copyright (c) 2021 Daniel Kanczuk

Easy to implement alghoritm creating a game with interface:

```go
type  State  interface {
	Simulate() SimulationResult
	Expand(Iteration) State
	Iterations() []Iteration
	Copy() State
	ID() string
}
```
**Simulate:** Play a random game based on current state.

**Iterations:** Total iterations of current state. 

**Expand:** Node expansion based on an previews Iteration sent

**Copy:** Deep copy of State. When implementing this interface, take care of copying objets in a way that doesn't carry any reference from the other one

**ID:** State identifier.It's a way to prevent duplicated instances

An example of implement:

```go
game := yourCustomGame{}
tree := mcts.NewMonteCarloTree(mcts.MonteCarloTreeConfig{MaxIterations: 256})

// the command start will block until process end
node, err := tree.Start(game)
if err != nil {
    panic(err)
}
// the new state of game, is the result of tree processing
newGameState := node.NodeScore[0].State.(yourCustomGame)

// you can use the newGameState again in monte carlo tree
// inner loop may be necessary to do it!
```

This implementation is based on example of [tic-tac-toe AI](https://vgarciasc.github.io/mcts-viz) and it was implemented with one expasion per iteration instead of expand all, to prevent memory overhead.

Check some examples here:

* [tic tac toe](https://github.com/danielsussa/go-mcts/blob/master/examples/tic-tac-toe/tictacgame.go)
* [game 2048](https://github.com/danielsussa/go-mcts/blob/master/examples/g2048/g2048.go)
