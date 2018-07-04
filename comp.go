package main

type Path map[Pos]Digest

type Action func(State) Digest

type ComputeHash struct {
	value   []byte
	path    Path
	actions []Action
}

func (c *ComputeHash) Exec(s State) Digest {
	return c.actions[s.cur.Dir(s.dst)](s)
}

func (c *ComputeHash) Cache(s State, d Digest) {
	c.path[s.cur] = d
}

func NewComputeHash(value []byte) *ComputeHash {
	var c ComputeHash
	c.value = value
	c.path = make(Path)
	c.actions = make([]Action, 2)
	var l, r Digest

	c.actions[Halt] = func(s State) Digest {
		c.path[s.cur] = hash(c.value)
		return c.path[s.cur]
	}

	c.actions[Next] = func(s State) Digest {
		l = c.path[s.cur.Left()]
		r = c.path[s.cur.Right()]

		c.path[s.cur] = hash(l, r)
		return c.path[s.cur]
	}

	return &c
}
