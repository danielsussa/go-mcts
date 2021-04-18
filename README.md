# Monte Carlo Tree Search

Copyright (c) 2021 Daniel Kanczuk

Easy to implement alghoritm creating a game with interface:

```go
type  State  interface {
	Simulate() SimulationResult
	Expand(idx int) State
	Iterations() int
	Player() string
	ID() string
}
```
**Simulate:** Play a random game based on current state.

**Iterations:** Total iterations of current state. 

**Expand:** Node expansion based on index of previews Iteration count

**Player:** Player who own the next move decision

**ID:** State identifier.

This implementation is based on example of [tic-tac-toe AI](https://vgarciasc.github.io/mcts-viz) and it was implemented with one expasion per iteration instead of expand all, to prevent memory overhead.

Check some examples here: