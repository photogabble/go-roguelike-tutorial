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
	Dark        TileGraphic // Graphics for when this tile is not in FOV
}

type Atlas map[string]*Tile

var TileAtlas = make(Atlas)

func NewTile(walkable, transparent bool, char rune, foregroundColor, backgroundColor concolor.Color) *Tile {
	return &Tile{
		Walkable:    walkable,
		Transparent: transparent,
		Dark: TileGraphic{
			Char:    int(char),
			FgColor: foregroundColor,
			BgColor: backgroundColor,
		},
	}
}

func InitTileAtlas() {
	TileAtlas["Floor"] = NewTile(true, true, ' ', concolor.RGB(255, 255, 255), concolor.RGB(50, 50, 150))
	TileAtlas["Wall"] = NewTile(false, false, ' ', concolor.RGB(255, 255, 255), concolor.RGB(0, 0, 100))
}
