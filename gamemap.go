package main

import (
	"fmt"
	"github.com/BigJk/ramen/console"
	"github.com/BigJk/ramen/t"
)

type GameMap struct {
	width  int
	height int
	tiles  []*Tile
}

// InBounds returns true if x and y are inside the bounds of this map
func (m *GameMap) InBounds(x, y int) bool {
	return 0 <= x && x < m.width && 0 <= y && y < m.height
}

func (m *GameMap) Render(con *console.Console) {

	for x := 0; x < m.width; x++ {
		for y := 0; y < m.height; y++ {
			tile := m.GetTile(x, y)
			con.Transform(x, y, t.Foreground(tile.Dark.FgColor), t.Background(tile.Dark.BgColor), t.Char(tile.Dark.Char))
		}
	}
}

func (m *GameMap) GetTile(x, y int) *Tile {
	return m.tiles[y*m.width+x]
}

func NewGameMap(width, height int) *GameMap {
	m := &GameMap{
		width:  width,
		height: height,
		tiles:  make([]*Tile, width*height),
	}

	fmt.Println(width*height, len(m.tiles))

	for i := range m.tiles {
		m.tiles[i] = TileAtlas["Floor"]
	}

	m.tiles[100] = TileAtlas["Wall"]

	return m
}
