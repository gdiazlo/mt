package main

type State struct {
	src, dst Pos
}

type Visit interface {
	Exec(State) Digest
	Cache(State, Digest)
}

func Traverse(t *Tree, src, dst Pos, v Visit) {
	var s State
	var d Digest
	var ok bool

	s.src = src
	s.dst = dst

	if d, ok = t.Cached(src, dst); ok {
		v.Cache(s, d)
		return
	}

	switch src.Dir(dst) {
	case Left:
		Traverse(t, src.Left(), dst, v)
		d = v.Exec(s)
	case Right:
		Traverse(t, src.Left(), dst, v)
		Traverse(t, src.Right(), dst, v)
		d = v.Exec(s)
	case Halt:
		d = v.Exec(s)
	}

	t.Cache(src, dst, d)

}
