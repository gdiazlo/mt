package main

import (
	"encoding/binary"
	"math"
)

type Dir int

const (
	Left Dir = iota
	Right
	Halt
)

type Pos struct {
	i, l uint64 // index, layer
}

func (p Pos) Dir(dst Pos) Dir {

	switch {
	case p.l == 0 || p.i > dst.i:
		return Halt
	case dst.i < p.i+pow(2, p.l-1):
		return Left
	case dst.i >= p.i+pow(2, p.l-1):
		return Right
	default:
		panic("Position out of tree")
	}

}

func (p Pos) Left() Pos {
	return Pos{p.i, p.l - 1}
}

func (p Pos) Right() Pos {
	return Pos{p.i + pow(2, p.l-1), p.l - 1}
}

func (p Pos) LeftSibbling() Pos {
	return Pos{p.i - 1, p.l}
}

func (p Pos) Bytes() []byte {
	b := make([]byte, 16)
	binary.LittleEndian.PutUint64(b, p.i)
	binary.LittleEndian.PutUint64(b[8:], p.l)
	return b
}

func (p Pos) Cached() bool {
	return false
}

func pow(x, y uint64) uint64 {
	return uint64(math.Pow(float64(x), float64(y)))
}