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

func (p Pos) Dir(dst Pos) Dir {

	if dst.l == 0 || p.i > dst.i {
		return Halt
	}
	return Next
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
