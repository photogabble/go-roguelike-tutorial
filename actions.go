package main

import (
	"fmt"
	"os"
)

type IAction interface {
	Perform(engine *Engine, entity *Entity)
}

type Action struct {
	Entity *Entity
}

func (action *Action) Perform(engine *Engine, entity *Entity) {
	// "Abstract" function to be implemented by sub-actions
}

type EscapeAction struct {
}

func (esc *EscapeAction) Perform(engine *Engine, entity *Entity) {
	os.Exit(0)
}

// ActionWithDirection is a base struct for Actions with destination
type ActionWithDirection struct {
	dx int
	dy int
}

type MovementAction struct {
	ActionWithDirection
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

	entityAtDestination := engine.gameMap.entities.AtLocation(dX, dY)

	if entityAtDestination != nil && entityAtDestination.blocksMovement == true {
		return // Destination blocked by Entity
	}

	entity.Move(move.dx, move.dy)
}

type MeleeAction struct {
	ActionWithDirection
}

func (action *MeleeAction) Perform(engine *Engine, entity *Entity) {
	dX := entity.X + action.dx
	dY := entity.Y + action.dy

	target := engine.gameMap.entities.AtLocation(dX, dY)

	if target != nil {
		fmt.Println(fmt.Sprintf("You kick the %s, much to its annoyance!", target.name))
	}
}

type BumpAction struct {
	ActionWithDirection
}

func (action *BumpAction) Perform(engine *Engine, entity *Entity) {
	dX := entity.X + action.dx
	dY := entity.Y + action.dy

	target := engine.gameMap.entities.AtLocation(dX, dY)

	if target == nil {
		NewMovementAction(action.dx, action.dy).Perform(engine, entity)
	} else {
		NewMeleeAction(action.dx, action.dy).Perform(engine, entity)
	}
}

func NewMeleeAction(dx, dy int) *MeleeAction {
	return &MeleeAction{
		ActionWithDirection{
			dx: dx,
			dy: dy,
		},
	}
}

func NewMovementAction(dx, dy int) *MovementAction {
	return &MovementAction{
		ActionWithDirection{
			dx: dx,
			dy: dy,
		},
	}
}

func NewBumpAction(dx, dy int) *BumpAction {
	return &BumpAction{
		ActionWithDirection{
			dx: dx,
			dy: dy,
		},
	}
}

func NewEscapeAction() *EscapeAction {
	return &EscapeAction{}
}
