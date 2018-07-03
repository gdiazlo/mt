package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.RWMutex
	var d Digest
	var v Visit
	event := []byte("test event 0")

	t := &Tree{make(Path), 0, m}

	for i := 0; i < 300; i++ {
		d, v = t.Add(event)
	}
	c := v.(*ComputeHash)

	fmt.Println(d, len(c.path))

}
