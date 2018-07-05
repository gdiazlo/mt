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

	Traverse(t, State{t.Root(), t.Last(), t.size}, c)

	return c.path[t.Root()], c
}

func (t Tree) Cached(s State) (d Digest, ok bool) {
	// if s.dst.i >= s.cur.i+pow(2, s.cur.l)-1 {
	d, ok = t.cache[s.cur]
	// }
	return
}

func (t *Tree) Cache(s State, d Digest) {
	// if s.dst.i >= s.cur.i+pow(2, s.cur.l)-1 {
	if d != nil {
		t.cache[s.cur] = d
	}
}
