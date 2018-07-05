package main

type Path map[Pos]Digest

type Action func(State) Digest

type ComputeVisitor struct {
<<<<<<< HEAD
	value   []byte
	path    Path
	actions []Action
}

func (c *ComputeVisitor) Exec(s State) Digest {
	if s.Next() {
		return c.actions[Next](s)
	}
	return c.actions[Halt](s)
}

func (c *ComputeVisitor) Cache(s State, d Digest) {
	c.path[s.cur] = d
}

func NewComputeVisitor(value []byte) *ComputeVisitor {
	var c ComputeVisitor
	c.value = value
	c.path = make(Path)
	c.actions = make([]Action, 2)
=======
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

func (c *ComputeVisitor) Path() Path {
	return c.path
}
>>>>>>> simplify visit interface, printer will not work probably

func (c *ComputeVisitor) halt(s State) Digest {
	c.path[s.cur] = hash(c.value)
	return c.path[s.cur]
}

<<<<<<< HEAD
	c.actions[Next] = func(s State) Digest {
		var l, r Digest
		l = c.path[s.cur.Left()]
		r = c.path[s.cur.Right()]
		c.path[s.cur] = hash(l, r)
		return c.path[s.cur]
	}
=======
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
>>>>>>> simplify visit interface, printer will not work probably

	return &c
}
