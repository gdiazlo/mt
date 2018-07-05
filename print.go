package main

import "fmt"

type PrintVisitor struct {
	m    Path
<<<<<<< HEAD
=======
	c, s uint64
>>>>>>> simplify visit interface, printer will not work probably
	next Action
	halt Action
}

func (p *PrintVisitor) Exec(s State) Digest {
<<<<<<< HEAD
=======

	defer func() {
		lc := p.s / (2 * (s.cur.l + 1))
		if p.c > lc {
			fmt.Printf("%d\n", lc)
			p.c = 0
		}
	}()

>>>>>>> simplify visit interface, printer will not work probably
	if s.Next() {
		return p.next(s)
	}
	return p.halt(s)
}

func (p *PrintVisitor) Cache(s State, d Digest) {
	fmt.Printf("cached")
}

<<<<<<< HEAD
func NewPrintVisitor(m Path) *PrintVisitor {
	var p PrintVisitor
	p.m = m
	p.halt = func(s State) Digest {
		fmt.Printf("%d", s.cur)
=======
func NewPrintVisitor(m Path, s uint64) *PrintVisitor {
	var p PrintVisitor
	p.m = m
	p.s = s

	p.halt = func(s State) Digest {
		fmt.Printf("%d", s.cur)
		p.c++
>>>>>>> simplify visit interface, printer will not work probably
		return nil
	}

	p.next = func(s State) Digest {
<<<<<<< HEAD
		fmt.Printf("%d", s.cur.Left())
		fmt.Printf("%d", s.cur.Right())
		if s.cur.i >= pow(2, s.cur.l)-1 {
			fmt.Printf("\n")
		}
=======
		fmt.Printf("%d", s.cur)
		p.c++
>>>>>>> simplify visit interface, printer will not work probably
		return nil
	}

	return &p
}
