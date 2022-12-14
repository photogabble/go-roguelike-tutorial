package main

import (
	"github.com/BigJk/ramen/console"
	"github.com/BigJk/ramen/font"
	"github.com/hajimehoshi/ebiten/v2"
	"os"
)

const (
	ScreenW = 80
	ScreenH = 50
)

func main() {
	rootConsole, err := console.New(ScreenW, ScreenH, font.DefaultFont, "Yet Another Roguelike Tutorial")
	if err != nil {
		panic(err)
	}

	player := NewPlayer()

	// Update loop, executed 60 times a second, unaffected by FPS
	rootConsole.SetTickHook(func(timeElapsed float64) error {
		action := EventHandler()

		if action != nil {
			switch action.(type) {
			case MovementAction:
				movement := action.(MovementAction)
				player.X += movement.dx
				player.Y += movement.dy
				break
			case EscapeAction:
				os.Exit(0)
			}
		}

		return nil
	})

	// Draw loop, executed before each frame is drawn to the screen
	rootConsole.SetPreRenderHook(func(screen *ebiten.Image, timeDelta float64) error {
		rootConsole.ClearAll() // Clear console
		rootConsole.Print(player.X, player.Y, "@")
		return nil
	})

	// Start the console with a scale factor of 2
	rootConsole.Start(2)
}
