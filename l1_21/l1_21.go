package l1_21

import "fmt"

type Printer interface {
	PrintModern(s string)
}

type OldPrinter struct {
}

func (p *OldPrinter) PrintLegacy(s string) {
	fmt.Println("LEGACY")
	fmt.Println(s)
}

type ModernPrinter struct {
}

func (p *ModernPrinter) PrintModern(s string) {
	fmt.Println("MODERN")
	fmt.Println(s)
}

type AdapterPrinter struct {
	OP *OldPrinter
}

func (p *AdapterPrinter) PrintModern(s string) {
	p.OP.PrintLegacy(s)
}
