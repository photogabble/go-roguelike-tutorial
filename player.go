package main

type Player struct {
	X int
	Y int
}

func NewPlayer() *Player {
	return &Player{
		X: ScreenW / 2,
		Y: ScreenH / 2,
	}
}
