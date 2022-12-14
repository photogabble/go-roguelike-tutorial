package main

type EscapeAction struct {
}

type MovementAction struct {
	dx int
	dy int
}

func NewMovementAction(dx, dy int) MovementAction {
	return MovementAction{
		dx: dx,
		dy: dy,
	}
}

func NewEscapeAction() EscapeAction {
	return EscapeAction{}
}
