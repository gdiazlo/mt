package main

type Path map[Pos]Digest

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
	c.path[s.cur] = d
}

func (c *ComputeVisitor) halt(s State) Digest {
	c.path[s.cur] = hash(c.value)
	return c.path[s.cur]
}

func (c *ComputeVisitor) next(s State) Digest {
	var l, r Digest
	l = c.path[s.cur.Left()]
	r = c.path[s.cur.Right()]
	c.path[s.cur] = hash(l, r)
	return c.path[s.cur]
}

func NewComputeVisitor(value []byte) *ComputeVisitor {
	var c ComputeVisitor
	c.value = value
	c.path = make(Path)

	return &c
}
