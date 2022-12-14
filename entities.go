package main

import (
	"github.com/BigJk/ramen/concolor"
	"github.com/BigJk/ramen/console"
	"github.com/BigJk/ramen/t"
)

// Entity is a generic struct to represent players, enemies, items,etc.
type Entity struct {
	X     int
	Y     int
	Char  int
	Color concolor.Color
}

// Move the entity by a given amount.
func (entity *Entity) Move(dx, dy int) {
	entity.X += dx
	entity.Y += dy
}

type EntityList struct {
	Entities []*Entity
}

func (list *EntityList) Add(entity *Entity) {
	list.Entities = append(list.Entities, entity)
}

func (list *EntityList) Render(con *console.Console) {
	for _, entity := range list.Entities {
		con.Transform(entity.X, entity.Y, t.Foreground(entity.Color), t.Char(entity.Char))
	}
}

func NewEntity(x, y int, char rune, color concolor.Color) *Entity {
	return &Entity{
		X:     x,
		Y:     y,
		Char:  int(char),
		Color: color,
	}
}
