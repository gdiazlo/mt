package main

type Action func(State) Digest

type ComputeVisitor struct {
	value []byte
	path  Path
	hash  Hasher
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
	if s.v < s.p.i {
		return nil
	}
	c.path[s.p] = c.hash(c.value)
	return c.path[s.p]
}

func (c *ComputeVisitor) next(s State) Digest {
	var l, r Digest
	l = c.path[s.p.Left()]
	r = c.path[s.p.Right()]
	c.path[s.p] = c.hash(l, r)
	return c.path[s.p]
}

func NewComputeVisitor(h Hasher, value []byte) *ComputeVisitor {
	var c ComputeVisitor
	c.value = value
	c.path = make(Path)
	c.hash = h

	return &c
}
