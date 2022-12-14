package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func EventHandler() interface{} {
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		return NewMovementAction(0, -1)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		return NewMovementAction(0, 1)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		return NewMovementAction(-1, 0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		return NewMovementAction(1, 0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return NewEscapeAction()
	}

	return nil
}
