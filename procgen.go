package main

import (
	"math/rand"
)

// RectangularRoom contains data relating to a rectangular room
type RectangularRoom struct {
	X1 int
	X2 int
	Y1 int
	Y2 int
}

// Center returns the (x,y) co-ordinates for the center of the room
func (room *RectangularRoom) Center() (int, int) {
	return (room.X1 + room.X2) / 2, (room.Y1 + room.Y2) / 2
}

// Inner returns the inner area of the room as a 2D array index
func (room *RectangularRoom) Inner() []Vector2i {
	locations := make([]Vector2i, (room.X2-(room.X1+1))*(room.Y2-(room.Y1+1)))
	i := 0

	for x := room.X1 + 1; x < room.X2; x++ {
		for y := room.Y1 + 1; y < room.Y2; y++ {
			locations[i] = Vector2i{
				X: x,
				Y: y,
			}
			i++
		}
	}

	return locations
}

// Intersects returns true if this room overlaps with another Rectangular Room
func (room *RectangularRoom) Intersects(other RectangularRoom) bool {
	return room.X1 <= other.X2 && room.X2 >= other.X1 && room.Y1 <= other.Y2 && room.Y2 >= other.Y1
}

// NewRectangularRoom returns a RectangularRoom pointer
func NewRectangularRoom(x, y, width, height int) *RectangularRoom {
	return &RectangularRoom{
		X1: x,
		Y1: y,
		X2: x + width,
		Y2: y + height,
	}
}

// RectangularRoomList is a helper struct for running Intersects on a list of RectangularRoom
type RectangularRoomList struct {
	Rooms []*RectangularRoom
}

// Add adds a RectangularRoom to the list
func (list *RectangularRoomList) Add(room *RectangularRoom) {
	list.Rooms = append(list.Rooms, room)
}

// Intersects loops over the list of RectangularRoom and returns true if any intersect with the input room.
func (list *RectangularRoomList) Intersects(room *RectangularRoom) bool {
	for _, other := range list.Rooms {
		if other.Intersects(*room) {
			return true
		}
	}

	return false
}

// Last retuns the last RectangularRoom to be added to the list or nil
func (list *RectangularRoomList) Last() *RectangularRoom {
	return list.Rooms[len(list.Rooms)-1]
}

// TunnelBetween returns an L-shaped tunnel generated by Bresenham's line drawing algorithm between start and end
// points as a list of Vector2i
func TunnelBetween(a, b RectangularRoom) []Vector2i {
	var cX, cY int
	x1, y1 := a.Center()
	x2, y2 := b.Center()

	if rand.Intn(100) < 50 {
		// Move Horizontally, then vertically.
		cX = x2
		cY = y1
	} else {
		// Move Vertically, then horizontally
		cX = x1
		cY = y2
	}

	// Concatenate the two slices for the L-shaped tunnel into one slice for digging.
	return append(BresenhamLine(Vector2i{cX, cY}, Vector2i{x2, y2}), BresenhamLine(Vector2i{x1, y1}, Vector2i{cX, cY})...)
}

// GenerateDungeon returns a fresh GameMap containing a generated dungeon made up of rectangular rooms joined by passageways.
func GenerateDungeon(width, height, maxRooms, minRoomSize, maxRoomSize int, player *Entity) *GameMap {
	entities := &EntityList{}
	entities.Add(player)

	dungeon := NewGameMap(width, height, entities)
	rooms := &RectangularRoomList{}

	for r := 0; r < maxRooms; r++ {
		roomWidth := rng(minRoomSize, maxRoomSize)
		roomHeight := rng(minRoomSize, maxRoomSize)

		x := rng(0, dungeon.width-roomWidth-1)
		y := rng(0, dungeon.height-roomHeight-1)

		newRoom := NewRectangularRoom(x, y, roomWidth, roomHeight)

		// Check room doesn't intersect with any other rooms
		if rooms.Intersects(newRoom) {
			continue
		}

		// Dig out this rooms inner area
		dungeon.SetArea(newRoom.Inner(), TileAtlas["Floor"])

		// If this is the first room, then place the player Entity in its center
		if len(rooms.Rooms) == 0 {
			player.X, player.Y = newRoom.Center()
		} else {
			// Dig out a tunnel between this room and the previous room
			dungeon.SetArea(TunnelBetween(*rooms.Last(), *newRoom), TileAtlas["Floor"])
		}

		// Append new room to list
		rooms.Add(newRoom)
	}

	return dungeon
}
