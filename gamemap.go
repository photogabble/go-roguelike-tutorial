package main

import (
	"github.com/BigJk/ramen/console"
	"github.com/BigJk/ramen/t"
	"github.com/norendren/go-fov/fov"
)

type GameMap struct {
	width    int
	height   int
	tiles    []*Tile
	explored []bool // Tile positions that have been explored
	fov      *fov.View
	entities *EntityList
	engine   *Engine
}

// InBounds returns true if x and y are inside the bounds of this map
func (m *GameMap) InBounds(x, y int) bool {
	return 0 <= x && x < m.width && 0 <= y && y < m.height
}

// IsOpaque returns true if a tile blocks the players FOV
func (m *GameMap) IsOpaque(x, y int) bool {
	return m.tiles[y*m.width+x].Transparent == false
}

// IsExplored returns true if a tile has been explored
func (m *GameMap) IsExplored(x, y int) bool {
	return m.explored[y*m.width+x]
}

// SetExplored sets a tile as explored
func (m *GameMap) SetExplored(x, y int) {
	m.explored[y*m.width+x] = true
}

// UpdateFov recomputes the visible are based upon the players point of view.
func (m *GameMap) UpdateFov() {
	m.fov.Compute(m, m.engine.player.X, m.engine.player.Y, 6)
}

func (m *GameMap) Render(con *console.Console) {
	var tGraphic TileGraphic

	for x := 0; x < m.width; x++ {
		for y := 0; y < m.height; y++ {
			tile := m.GetTile(x, y)

			if m.fov.IsVisible(x, y) {
				m.SetExplored(x, y)
				tGraphic = tile.Light
			} else if m.IsExplored(x, y) {
				tGraphic = tile.Dark
			} else {
				tGraphic = SHROUD
			}

			con.Transform(x, y, t.Foreground(tGraphic.FgColor), t.Background(tGraphic.BgColor), t.Char(tGraphic.Char))
		}
	}

	m.entities.Render(m, con)
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

func NewGameMap(width, height int, entities *EntityList, engine *Engine) *GameMap {
	m := &GameMap{
		width:    width,
		height:   height,
		tiles:    make([]*Tile, width*height),
		explored: make([]bool, width*height),
		fov:      fov.New(),
		entities: entities,
		engine:   engine,
	}

	for i := range m.tiles {
		m.tiles[i] = TileAtlas["Wall"]
		m.explored[i] = false
	}

	return m
}
