package main

type Path map[Pos]Digest

type Action func(State) Digest

type ComputeHash struct {
	value   []byte
	path    Path
	actions []Action
}

func (c *ComputeHash) Exec(s State) Digest {
	return c.actions[s.src.Dir(s.dst)](s)
}

func (c *ComputeHash) Cache(s State, d Digest) {
	c.path[s.src] = d
}

func NewComputeHash(value []byte) *ComputeHash {
	var c ComputeHash
	c.value = value
	c.path = make(Path)
	c.actions = make([]Action, 3)

	c.actions[Halt] = func(s State) Digest {
		// c.path[s.src] = hash(s.src.Bytes(), c.value)
		c.path[s.src] = hash(c.value)
		return c.path[s.src]
	}

	c.actions[Left] = func(s State) Digest {
		// c.path[s.src] = hash(s.src.Bytes(), c.path[s.src])
		c.path[s.src] = hash(c.path[s.src])
		return c.path[s.src]
	}

	c.actions[Right] = func(s State) Digest {
		both := append(c.path[s.src], c.path[s.src.LeftSibbling()]...)
		// c.path[s.src] = hash(s.src.Bytes(), both)
		c.path[s.src] = hash(both)
		return c.path[s.src]
	}

	return &c
}
