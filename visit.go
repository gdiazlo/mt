package main

type State struct {
	cur, dst Pos
}

type Visit interface {
	Exec(State) Digest
	Cache(State, Digest)
}

func Traverse(t *Tree, cur, dst Pos, v Visit) {
	var s State
	var d Digest
	var ok bool

	s.cur = cur
	s.dst = dst

	if d, ok = t.Cached(cur, dst); ok {
		v.Cache(s, d)
		return
	}

	switch cur.Dir(dst) {
	case Next:
		Traverse(t, cur.Left(), dst, v)
		Traverse(t, cur.Right(), dst, v)
		d = v.Exec(s)
	case Halt:
		d = v.Exec(s)
	}

	t.Cache(cur, dst, d)

}
