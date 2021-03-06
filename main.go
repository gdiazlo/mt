package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.RWMutex
	var d Digest
	var v Visitor
	event := []byte("test event 0")

	t := &Tree{make(Path), 0, xorhasher, m}

	for i := 0; i < 300; i++ {
		d, v = t.Add(event)
	}
	c := v.(*ComputeVisitor)

	fmt.Println(d, len(c.path))

}
