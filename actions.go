package main

import "os"

type Action interface {
	Perform(engine *Engine, entity *Entity)
}

type EscapeAction struct {
}

func (esc *EscapeAction) Perform(engine *Engine, entity *Entity) {
	os.Exit(0)
}

type MovementAction struct {
	dx int
	dy int
}

func (move *MovementAction) Perform(engine *Engine, entity *Entity) {
	dX := entity.X + move.dx
	dY := entity.Y + move.dy

	if !engine.gameMap.InBounds(dX, dY) {
		return // Destination out of bounds
	}

	if !engine.gameMap.GetTile(dX, dY).Walkable {
		return // Destination is blocked by a tile
	}

	entity.Move(move.dx, move.dy)
}

func NewMovementAction(dx, dy int) *MovementAction {
	return &MovementAction{
		dx: dx,
		dy: dy,
	}
}

func NewEscapeAction() *EscapeAction {
	return &EscapeAction{}
}
