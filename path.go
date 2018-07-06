package main

import (
	"fmt"
	"sort"
)

type Path map[Pos]Digest

func (p Path) Stringt() string {
	var keys []Pos

	for k := range p {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		if keys[i].l == keys[j].l {
			return keys[i].i < keys[j].i
		}
		return keys[i].l < keys[j].l
	})

	var t string
	var last = uint64(0)
	for _, k := range keys {
		if k.l != last {
			t = fmt.Sprintf("%v\n", t)
			last = k.l
		}
		t = fmt.Sprintf("%v %b", t, k)
	}

	return t
}