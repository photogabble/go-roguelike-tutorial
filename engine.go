package main

import (
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

func NewEngine(player *Entity, gameMap *GameMap) *Engine {
	engine := &Engine{
		gameMap: gameMap,
		player:  player,
	}

	// Init FOV
	gameMap.UpdateFov(player)

	return engine
}
