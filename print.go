package main

import "fmt"

type Print struct {
}

func (c *Print) Exec(s State) Digest {
	fmt.Println(s.cur)
	return nil
}

func (c *Print) Cache(s State, d Digest) {
	fmt.Println(s.cur)
}

func PrintTree(p Path, s, h uint64) {
	var i, l uint64
	cs := s

	fmt.Printf("Tree size: %d height: %d\n", s, h)
	cs = s - 1
	for i = 0; i <= h; i++ {
		fmt.Printf("%[2]*.[1]s", "", i)
		for l = 0; l <= cs; l++ {
			k := Pos{i, l}
			fmt.Printf(" %d ", p[k])
		}
		fmt.Printf("\n")
		cs = cs / 2
	}
}
