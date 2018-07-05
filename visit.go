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
	Visit(State) Digest
	VisitCached(State, Digest)
}

func Traverse(t *Tree, s State, v Visit) {
	var d Digest
	var ok bool

	// defer func() { fmt.Println("Visit -> ", s.cur) }()

	if d, ok = t.Cached(s); ok {
<<<<<<< HEAD
		v.Cache(s, d)
=======
		v.VisitCached(s, d)
>>>>>>> simplify visit interface, printer will not work probably
		return
	}

	if s.Next() {
		Traverse(t, s.Left(), v)
		Traverse(t, s.Right(), v)
	}
	d = v.Visit(s)

<<<<<<< HEAD
	d = v.Exec(s)

=======
>>>>>>> simplify visit interface, printer will not work probably
	t.Cache(s, d)

}
