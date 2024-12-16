package main

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Field struct {
	direction  Direction
	distance   int
	isTarget   bool
	isBoundary bool
}
