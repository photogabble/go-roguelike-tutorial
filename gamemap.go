package main

import (
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

// GetTile returns the Tile found at a given position
func (m *GameMap) GetTile(x, y int) *Tile {
	return m.tiles[y*m.width+x]
}

// SetArea sets Tile in the GameMap for a given area
func (m *GameMap) SetArea(area []Vector2i, tile *Tile) {
	for _, loc := range area {
		m.tiles[loc.Y*m.width+loc.X] = tile
	}
}

func NewGameMap(width, height int) *GameMap {
	m := &GameMap{
		width:  width,
		height: height,
		tiles:  make([]*Tile, width*height),
	}

	for i := range m.tiles {
		m.tiles[i] = TileAtlas["Wall"]
	}

	return m
}
