package main

import (
	"fmt"
	"github.com/BigJk/ramen/console"
)

// Engine acts as a store of game data that can easily be passed to functions that need it
type Engine struct {
	player  *Entity
	gameMap *GameMap
}

func (e *Engine) Render(con *console.Console) {
	e.gameMap.Render(con)
}

// HandleEnemyTurns performs enemy entities actions
func (e *Engine) HandleEnemyTurns() {
	for _, entity := range e.gameMap.entities.Entities {
		if entity == e.player {
			continue
		}

		fmt.Println(fmt.Sprintf("The %s wonders when it will get to take a real turn", entity.name))
	}
}

// HandlePlayerTurn performs players action
func (e *Engine) HandlePlayerTurn(action IAction) {
	if action == nil {
		return
	}

	action.Perform(e, e.player)

	e.gameMap.UpdateFov()
	e.HandleEnemyTurns()
}

func NewEngine() *Engine {
	engine := &Engine{
		player: Player.Spawn(ScreenW/2, ScreenH/2, nil),
	}

	GenerateDungeon(80, 45, 30, 6, 30, 2, engine)

	// Init FOV
	engine.gameMap.UpdateFov()

	return engine
}
