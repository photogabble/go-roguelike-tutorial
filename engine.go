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
func (e *Engine) HandlePlayerTurn(action Action) {
	if action == nil {
		return
	}

	action.Perform(e, e.player)

	e.gameMap.UpdateFov(e.player)
	e.HandleEnemyTurns()
}

func NewEngine(player *Entity, gameMap *GameMap) *Engine {
	engine := &Engine{
		gameMap: gameMap,
		player:  player,
	}

	// Init FOV
	gameMap.UpdateFov(player)

	return engine
}
