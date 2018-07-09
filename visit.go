package main

type State struct {
	p Pos
	v uint64
}

func (s State) Left() State {
	return State{s.p.Left(), s.v}
}

func (s State) Right() State {
	return State{s.p.Right(), s.v}
}

func (s State) Next() bool {
	if s.v < s.p.i {
		return false
	}
	return s.p.Next()
}

type Visitor interface {
	Visit(State) Digest
	VisitCached(State, Digest)
}

func Traverse(t *Tree, s State, v Visitor) {
	var d Digest
	var ok bool

	if d, ok = t.Cached(s); ok {
		v.VisitCached(s, d)
		// return
	}
	if s.Next() {
		Traverse(t, s.Left(), v)
		Traverse(t, s.Right(), v)
	}

	if d == nil {
		d = v.Visit(s)
		t.Cache(s, d)
	}

}
