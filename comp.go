package main

import (
	"fmt"
	"sort"
)

type Path map[Pos]Digest

func (p Path) String() string {
	var keys []Pos

	for k := range p {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		if keys[i].l == keys[j].l {
			return keys[i].i < keys[j].i
		}
		return keys[i].l < keys[j].l
	})

	var t string
	var last = uint64(0)
	for _, k := range keys {
		if k.l != last {
			t = fmt.Sprintf("%v\n", t)
			last = k.l
		}
		t = fmt.Sprintf("%v %v", t, k)
	}

	return t
}

type Action func(State) Digest

type ComputeVisitor struct {
	value []byte
	path  Path
}

func (c *ComputeVisitor) Visit(s State) Digest {
	if s.Next() {
		return c.next(s)
	}
	return c.halt(s)
}

func (c *ComputeVisitor) VisitCached(s State, d Digest) {
	c.path[s.p] = d
}

func (c *ComputeVisitor) halt(s State) Digest {
	c.path[s.p] = hash(c.value)
	return c.path[s.p]
}

func (c *ComputeVisitor) next(s State) Digest {
	var l, r Digest
	l = c.path[s.p.Left()]
	r = c.path[s.p.Right()]
	c.path[s.p] = hash(l, r)
	return c.path[s.p]
}

func NewComputeVisitor(value []byte) *ComputeVisitor {
	var c ComputeVisitor
	c.value = value
	c.path = make(Path)

	return &c
}
