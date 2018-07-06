package main

type MetaVisitor struct {
	vs []Visitor
}

func (m MetaVisitor) Visit(s State) Digest {
	var d Digest
	for _, v := range m.vs {
		dv := v.Visit(s)
		if dv != nil {
			d = dv
		}
	}
	return d
}

func (m MetaVisitor) VisitCached(s State, d Digest) {
	for _, v := range m.vs {
		v.VisitCached(s, d)
	}
}

func NewMetaVisitor(vs ...Visitor) Visitor {
	var m MetaVisitor
	m.vs = make([]Visitor, len(vs))
	m.vs = append(m.vs, vs...)

	return &m
}