package main

import (
	"math"
	"sync"
)

type Tree struct {
	cache Path
	size  uint64
	sync.RWMutex
}

func (t Tree) Root() Pos {
	return Pos{0, uint64(math.Ceil(math.Log2(float64(t.size))))}
}

func (t Tree) Last() Pos {
	return Pos{t.size, 0}
}

func (t *Tree) Add(event []byte) (Digest, Visit) {
	t.Lock()
	defer t.Unlock()
	c := NewComputeHash(event)
	t.size++
	Traverse(t, t.Root(), t.Last(), c)

	return c.path[t.Root()], c
}

func (t Tree) Cached(src, dst Pos) (d Digest, ok bool) {
	if dst.i >= src.i+pow(2, src.l)-1 {
		d, ok = t.cache[src]
	}
	return
}

func (t *Tree) Cache(src, dst Pos, d Digest) {
	if dst.i >= src.i+pow(2, src.l)-1 {
		t.cache[src] = d
	}
}
