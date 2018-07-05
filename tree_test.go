package main

import (
	"fmt"
	"sync"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func newTree() *Tree {
	var m sync.RWMutex
	t := Tree{make(Path), 0, m}
	return &t
}

func TestAdd(t *testing.T) {

	testCases := []struct {
		eventDigest      Digest
		expectedRootHash Digest
	}{
		{[]byte{0x0}, []byte{0x0}},
		{[]byte{0x1}, []byte{0x1}},
		{[]byte{0x2}, []byte{0x1}},
		{[]byte{0x3}, []byte{0x1}},
		{[]byte{0x4}, []byte{0x1}},
		{[]byte{0x5}, []byte{0x1}},
		{[]byte{0x6}, []byte{0x2}},
		{[]byte{0x7}, []byte{0x2}},
		{[]byte{0x8}, []byte{0x2}},
		{[]byte{0x9}, []byte{0x2}},
	}

	tree := newTree()

	for i, c := range testCases {
		rh, v := tree.Add(c.eventDigest)
<<<<<<< HEAD
		ch := v.(*ComputeVisitor)
		// r := tree.Root()
		// fmt.Println(ch.path)
		// PrintTree(ch.path, tree.size, r.l)
		Traverse(newTree(), State{tree.Root(), tree.Last(), tree.size}, NewPrintVisitor(ch.path))
		fmt.Printf("\n-----------------\n")
=======
		cv := v.(*ComputeVisitor)
		fmt.Println(cv.path)

>>>>>>> simplify visit interface, printer will not work probably
		assert.Equalf(t, c.expectedRootHash, rh, "Incorrect root hash for index %d", i)
	}
}
