package main

import (
	"github.com/BigJk/ramen/concolor"
	"github.com/BigJk/ramen/console"
	"github.com/BigJk/ramen/font"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenW = 80
	ScreenH = 50
)

func main() {
	InitTileAtlas()

	rootConsole, err := console.New(ScreenW, ScreenH, font.DefaultFont, "Yet Another Roguelike Tutorial")
	if err != nil {
		panic(err)
	}

	player := NewEntity(ScreenW/2, ScreenH/2, '@', concolor.RGB(255, 69, 0))

	engine := NewEngine(
		player,
		GenerateDungeon(80, 45, 30, 6, 30, player),
	)

	// Update loop, executed 60 times a second, unaffected by FPS
	rootConsole.SetTickHook(func(timeElapsed float64) error {
		action := EventHandler()

		if action != nil {
			action.Perform(engine, engine.player)
		}

		return nil
	})

	// Draw loop, executed before each frame is drawn to the screen
	rootConsole.SetPreRenderHook(func(screen *ebiten.Image, timeDelta float64) error {
		rootConsole.ClearAll() // Clear console
		engine.Render(rootConsole)
		return nil
	})

	// Start the console with a scale factor of 2
	rootConsole.Start(2)
}
