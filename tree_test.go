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

}

func newTree() *Tree {
	var m sync.RWMutex
	t := Tree{make(Path), 0, m}
	return &t
}

func TestAdd(t *testing.T) {

	tree := newTree()

	for i, c := range testCases {
		rh, _ := tree.Add(c.in)
		assert.Equalf(t, c.out, rh, "Incorrect root hash for index %d", i)
	}
}

func TestIncremental(t *testing.T) {
	tree := newTree()
	d := make([]Digest, len(testCases))

	for i, c := range testCases {
		d[i], _ = tree.Add(c.in)
		assert.Equalf(t, c.out, d[i], "Incorrect root hash for index %d", i)
	}

	j := Pos{4, 0}
	k := Pos{9, 0}

	di, v := tree.Incremental(j, k)
	ch := v.(*ComputeVisitor)

	fmt.Println("---")
	fmt.Println(ch.path)

	e := testCases[j.i].in
	cv := NewComputeVisitor(e)
	pv := NewPrintVisitor()

	jv := NewMetaVisitor(cv, pv)
	kv := NewMetaVisitor(cv, pv)

	Traverse(tree, State{j, 4}, jv)
	Traverse(tree, State{j, 9}, kv)

	fmt.Println(di, " ", d[4], " == ", jv, d[9], " == ", kv)

}
