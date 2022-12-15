package main

import (
	"github.com/BigJk/ramen/concolor"
	"github.com/BigJk/ramen/console"
)

// Engine acts as a store of game data that can easily be passed to functions that need it
type Engine struct {
	entities *EntityList
	player   *Entity
	gameMap  *GameMap
}

func (e *Engine) Render(con *console.Console) {
	e.gameMap.Render(con)
	e.entities.Render(con)
}

func NewEngine(player *Entity, gameMap *GameMap) *Engine {
	engine := &Engine{
		entities: &EntityList{},
		gameMap:  gameMap,
		player:   player,
	}

	engine.entities.Add(player)
	engine.entities.Add(NewEntity((ScreenW/2)-5, ScreenH/2, '@', concolor.RGB(255, 255, 0)))

	return engine
}
