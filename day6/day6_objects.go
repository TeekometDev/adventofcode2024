package main

type Direction int

const (
	Upwards Direction = iota
	Right
	Downwards
	Left
)

type Guard struct {
	row    int
	col    int
	facing Direction
}

type DirectionTile struct {
	up    bool
	down  bool
	left  bool
	right bool
}
