package main

import "fmt"

type PrintVisitor struct {
	m    Path
	next Action
	halt Action
}

func (p *PrintVisitor) Exec(s State) Digest {
	if s.Next() {
		return p.next(s)
	}
	return p.halt(s)
}

func (p *PrintVisitor) Cache(s State, d Digest) {
	fmt.Printf("cached")
}

func NewPrintVisitor(m Path) *PrintVisitor {
	var p PrintVisitor
	p.m = m
	p.halt = func(s State) Digest {
		fmt.Printf("%d", s.cur)
		return nil
	}

	p.next = func(s State) Digest {
		fmt.Printf("%d", s.cur.Left())
		fmt.Printf("%d", s.cur.Right())
		if s.cur.i >= pow(2, s.cur.l)-1 {
			fmt.Printf("\n")
		}
		return nil
	}

	return &p
}
