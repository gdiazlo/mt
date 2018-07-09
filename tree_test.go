package main

import (
	"fmt"
	"sync"
	"testing"

	assert "github.com/stretchr/testify/require"
)

type Case struct {
	in  Digest
	out Digest
}

var testCases []Case
var hashers []Hasher

func init() {
	testCases = []Case{
		{[]byte{0x0}, []byte{0x0}},
		{[]byte{0x1}, []byte{0x1}},
		{[]byte{0x2}, []byte{0x3}},
		{[]byte{0x3}, []byte{0x0}},
		{[]byte{0x4}, []byte{0x4}},
		{[]byte{0x5}, []byte{0x1}},
		{[]byte{0x6}, []byte{0x7}},
		{[]byte{0x7}, []byte{0x0}},
		{[]byte{0x8}, []byte{0x8}},
		{[]byte{0x9}, []byte{0x1}},
	}

	hashers = []Hasher{
		xorhasher,
		pearsonhasher,
		sha256hasher,
	}

}
func duppath(p Path) Path {
	np := make(Path)
	for k, v := range p {
		np[k] = v
	}
	return np
}
func newTree(h Hasher) *Tree {
	var m sync.RWMutex
	t := Tree{make(Path), 0, h, m}
	return &t
}

func TestAddXor(t *testing.T) {

	tree := newTree(xorhasher)

	for i, c := range testCases {
		rh, _ := tree.Add(c.in)
		assert.Equalf(t, c.out, rh, "Incorrect root hash for index %d", i)
	}
}

func TestAddAndVerify(t *testing.T) {

	for i, h := range hashers {
		tr := newTree(h)

		event := h([]byte("a test event"))

		_, v := tr.Add(event)
		ch := *v.(*ComputeVisitor)

		nt := newTree(h)
		nt.cache = duppath(ch.path)
		nt.size = 1

		nc := NewComputeVisitor(h, event)
		Traverse(nt, State{nt.Root(), nt.size - 1}, nc)

		assert.Equal(t, nc.path[nt.Root()], ch.path[tr.Root()], "root digests !equal when using hasher %d", i)
	}

}

func TestIncremental(t *testing.T) {
	tree := newTree(xorhasher)
	d := make([]Digest, len(testCases))

	for i, c := range testCases {
		d[i], _ = tree.Add(c.in)
		assert.Equalf(t, c.out, d[i], "Incorrect root hash for index %d", i)
	}

	j := Pos{4, 0}
	rj := Pos{0, 3}
	k := Pos{9, 0}
	rk := Pos{0, 4}

	di, v := tree.Incremental(j, k)
	ch := v.(*ComputeVisitor)

	fmt.Println("---")
	fmt.Println(len(ch.path), ch.path)

	e := testCases[j.i].in
	cv := NewComputeVisitor(xorhasher, e)

	/*
		pv := NewPrintVisitor()
		jv := NewMetaVisitor(cv, pv)
		kv := NewMetaVisitor(cv, pv)
	*/
	fmt.Println("digest ", di)
	Traverse(tree, State{rj, 4}, cv)
	fmt.Println("From ", rj, " -->  ", cv.path)
	Traverse(tree, State{rk, 9}, cv)
	fmt.Println("From ", rk, " -->  ", cv.path)

}
