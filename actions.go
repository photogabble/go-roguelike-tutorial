package main

import (
	"fmt"
	"os"
)

type IAction interface {
	Perform()
}

// Action is the base struct, intended to be extended by structs that implement the Perform function of IAction
type Action struct {
	entity *Entity
}

// engine is a shortcut to Engine
func (action *Action) engine() *Engine {
	return action.entity.dungeon.engine
}

func (action *Action) Perform() {
	// "Abstract" function to be implemented by sub-actions
}

type EscapeAction struct {
}

func (esc *EscapeAction) Perform() {
	os.Exit(0)
}

// ActionWithDirection is a base struct for Actions with destination
type ActionWithDirection struct {
	Action
	dx int
	dy int
}

func (action *ActionWithDirection) DestinationXY() (int, int) {
	return action.entity.X + action.dx, action.entity.Y + action.dy
}

// BlockingEntity returns the blocking Entity at this Actions destination
func (action *ActionWithDirection) BlockingEntity() *Entity {
	dX, dY := action.DestinationXY()
	entityAtDestination := action.engine().gameMap.entities.AtLocation(dX, dY)

	if entityAtDestination != nil && entityAtDestination.blocksMovement == true {
		return entityAtDestination
	}

	return nil
}

type MovementAction struct {
	ActionWithDirection
}

func (move *MovementAction) Perform() {
	dX, dY := move.DestinationXY()

	if !move.engine().gameMap.InBounds(dX, dY) {
		return // Destination out of bounds
	}

	if !move.engine().gameMap.GetTile(dX, dY).Walkable {
		return // Destination is blocked by a tile
	}

	if move.BlockingEntity() != nil {
		return // Destination blocked by Entity
	}

	move.entity.Move(move.dx, move.dy)
}

type MeleeAction struct {
	ActionWithDirection
}

// Perform this MeleeAction on Entity
func (action *MeleeAction) Perform() {
	target := action.BlockingEntity()

	if target != nil {
		fmt.Println(fmt.Sprintf("You kick the %s, much to its annoyance!", target.name))
	}
}

type BumpAction struct {
	ActionWithDirection
}

func (action *BumpAction) Perform() {
	target := action.BlockingEntity()

	if target == nil {
		NewMovementAction(action.entity, action.dx, action.dy).Perform()
	} else {
		NewMeleeAction(action.entity, action.dx, action.dy).Perform()
	}
}

func NewMeleeAction(entity *Entity, dx, dy int) *MeleeAction {
	return &MeleeAction{
		ActionWithDirection{
			Action: Action{entity},
			dx:     dx,
			dy:     dy,
		},
	}
}

func NewMovementAction(entity *Entity, dx, dy int) *MovementAction {
	return &MovementAction{
		ActionWithDirection{
			Action: Action{entity},
			dx:     dx,
			dy:     dy,
		},
	}
}

func NewBumpAction(entity *Entity, dx, dy int) *BumpAction {
	return &BumpAction{
		ActionWithDirection{
			Action: Action{entity},
			dx:     dx,
			dy:     dy,
		},
	}
}

func NewEscapeAction() *EscapeAction {
	return &EscapeAction{}
}
