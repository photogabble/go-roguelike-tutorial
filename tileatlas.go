package main

import "github.com/BigJk/ramen/concolor"

type TileGraphic struct {
	Char    int
	FgColor concolor.Color
	BgColor concolor.Color
}

type Tile struct {
	Walkable    bool        // True if this tile can be walked over
	Transparent bool        // True if this tile doesn't block FOV
	Dark        TileGraphic // Graphic for when this tile is not in FOV
	Light       TileGraphic // Graphic for when this tile is in FOV
}

type Atlas map[string]*Tile

var TileAtlas = make(Atlas)

// SHROUD is used when a tile is neither in view nor has been "explored"
var SHROUD = TileGraphic{
	Char:    ' ',
	FgColor: concolor.RGB(255, 255, 255),
	BgColor: concolor.RGB(0, 0, 0),
}

// NewTile is a helper function for defining individual Tile types
func NewTile(walkable, transparent bool, darkGraphic, lightGraphic TileGraphic) *Tile {
	return &Tile{
		Walkable:    walkable,
		Transparent: transparent,
		Dark:        darkGraphic,
		Light:       lightGraphic,
	}
}

func InitTileAtlas() {
	TileAtlas["Floor"] = NewTile(true, true, TileGraphic{int(' '), concolor.RGB(255, 255, 255), concolor.RGB(50, 50, 150)}, TileGraphic{' ', concolor.RGB(255, 255, 255), concolor.RGB(200, 180, 50)})
	TileAtlas["Wall"] = NewTile(false, false, TileGraphic{int(' '), concolor.RGB(255, 255, 255), concolor.RGB(0, 0, 100)}, TileGraphic{' ', concolor.RGB(255, 255, 255), concolor.RGB(130, 110, 50)})
}
