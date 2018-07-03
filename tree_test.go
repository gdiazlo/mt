package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
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
		{[]byte{0x2}, []byte{0x3}},
		{[]byte{0x3}, []byte{0x0}},
		{[]byte{0x4}, []byte{0x4}},
		{[]byte{0x5}, []byte{0x1}},
		{[]byte{0x6}, []byte{0x7}},
		{[]byte{0x7}, []byte{0x0}},
		{[]byte{0x8}, []byte{0x8}},
		{[]byte{0x9}, []byte{0x1}},
	}

	tree := newTree()

	// Note that we are using fake hashing functions and the index
	// as the value of the event's digest to make predictable hashes

	for i, c := range testCases {
		rh, v := tree.Add(c.eventDigest)
		ch := v.(*ComputeHash)
		fmt.Println(ch.path)
		assert.Equalf(t, c.expectedRootHash, rh, "Incorrect root hash for index %d", i)
	}

}
