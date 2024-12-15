package main

type Movement int

const (
	Up Movement = iota
	Right
	Down
	Left
)

type FieldObject int

const (
	Empty FieldObject = iota
	Boundary
	BoxLeft
	BoxRight
)

type Robot struct {
	x_pos        int
	y_pos        int
	movementList []Movement
}

type Warehouse struct {
	x_size int
	y_size int
}

type Position struct {
	x_pos int
	y_pos int
}
