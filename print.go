package main

import "fmt"

type PrintVisitor struct {
	next Action
	halt Action
}

func (p *PrintVisitor) Visit(s State) Digest {
	fmt.Println(s)
	return nil
}

func (p *PrintVisitor) VisitCached(s State, d Digest) {
	fmt.Println(s)
}

func NewPrintVisitor() *PrintVisitor {
	var p PrintVisitor
	return &p
}
