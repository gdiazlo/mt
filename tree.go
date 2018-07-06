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
	c := NewComputeVisitor(event)
	t.size++
	Traverse(t, State{t.Root(), t.size - 1}, c)

	return c.path[t.Root()], c
}

func (t *Tree) Incremental(j, k Pos) (Digest, Visit) {
	t.Lock()
	defer t.Unlock()
	c := NewComputeVisitor(nil)
	t.size++
	Traverse(t, State{t.Root(), j.i}, c)
	Traverse(t, State{t.Root(), k.i}, c)

	return c.path[t.Root()], c
}

func (t Tree) Cached(s State) (d Digest, ok bool) {
	if s.v >= s.p.i+pow(2, s.p.l)-1 {
		d, ok = t.cache[s.p]
	}
	return
}

func (t *Tree) Cache(s State, d Digest) {
	if s.v >= s.p.i+pow(2, s.p.l)-1 {
		t.cache[s.p] = d
	}
}
