package main

import (
	"math"
)

type Dir int

const (
	Next Dir = iota
	Halt
)

type Pos struct {
	i, l uint64 // index, layer
}

func (p Pos) Next(dst Pos) bool {
	return p.l != 0
}

func (p Pos) Left() Pos {
	return Pos{p.i, p.l - 1}
}

func (p Pos) Right() Pos {
	return Pos{p.i + pow(2, p.l-1), p.l - 1}
}

func pow(x, y uint64) uint64 {
	return uint64(math.Pow(float64(x), float64(y)))
}
