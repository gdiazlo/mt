package main

type State struct {
	cur, dst Pos
	n        uint64
}

func (s State) Left() State {
	return State{s.cur.Left(), s.dst, s.n}
}

func (s State) Right() State {
	return State{s.cur.Right(), s.dst, s.n}
}

func (s State) Next() bool {
	return s.cur.Next(s.dst) && s.n >= 1
}

type Visit interface {
	Exec(State) Digest
	Cache(State, Digest)
}

func Traverse(t *Tree, s State, v Visit) {
	var d Digest
	var ok bool

	// defer func() { fmt.Println("Visit -> ", s.cur) }()

	if d, ok = t.Cached(s); ok {
		v.Cache(s, d)
		return
	}

	if s.Next() {
		Traverse(t, s.Left(), v)
		Traverse(t, s.Right(), v)
	}

	d = v.Exec(s)

	t.Cache(s, d)

}
