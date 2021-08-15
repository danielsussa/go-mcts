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

Ok.

This implementation is based on example of [tic-tac-toe AI](https://vgarciasc.github.io/mcts-viz) and it was implemented with one expasion per iteration instead of expand all, to prevent memory overhead.

Check some examples here:

* [tic tac toe](https://github.com/danielsussa/go-mcts/blob/master/examples/tic-tac-toe/tictacgame.go)
* [game 2048](https://github.com/danielsussa/go-mcts/blob/master/examples/g2048/g2048.go)
