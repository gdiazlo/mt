package main

type Path map[Pos]Digest

type Action func(State) Digest

type ComputeVisitor struct {
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

	c.actions[Halt] = func(s State) Digest {
		c.path[s.cur] = hash(c.value)
		return c.path[s.cur]
	}

	c.actions[Next] = func(s State) Digest {
		var l, r Digest
		l = c.path[s.cur.Left()]
		r = c.path[s.cur.Right()]
		c.path[s.cur] = hash(l, r)
		return c.path[s.cur]
	}

	return &c
}
