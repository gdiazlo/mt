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
		{[]byte{0x3}, []byte{0x0}},
		{[]byte{0x4}, []byte{0x4}},
		{[]byte{0x5}, []byte{0x4}},
		{[]byte{0x6}, []byte{0x1}},
		{[]byte{0x7}, []byte{0x0}},
		{[]byte{0x8}, []byte{0x0}},
		{[]byte{0x9}, []byte{0x1}},
	}

	tree := newTree()

	for i, c := range testCases {
		rh, v := tree.Add(c.eventDigest)
		ch := v.(*ComputeVisitor)
		fmt.Println(tree.Root(), len(ch.path), ch.path)
		assert.Equalf(t, c.expectedRootHash, rh, "Incorrect root hash for index %d", i)
	}
}

func TestIncremental(t *testing.T) {
	tree := newTree()

	for i := 0; i < 13; i++ {
		tree.Add([]byte{0x0})
	}

	j := Pos{4, 0}
	k := Pos{12, 0}

	d, v := tree.Incremental(j, k)
	ch := v.(*ComputeVisitor)
	fmt.Println(d, len(ch.path), ch.path)
}
