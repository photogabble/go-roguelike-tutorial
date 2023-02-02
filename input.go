package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// EventHandler listens for ebiten input events and translates them to IAction's.
func EventHandler() IAction {
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		return NewBumpAction(0, -1)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		return NewBumpAction(0, 1)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		return NewBumpAction(-1, 0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		return NewBumpAction(1, 0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return NewEscapeAction()
	}

	return nil
}
