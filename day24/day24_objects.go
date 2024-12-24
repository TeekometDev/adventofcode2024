package main

type Rule struct {
	in1 string
	in2 string
	out string
	op  Operation
}

type Operation int

const (
	AND Operation = iota
	OR
	XOR
)
