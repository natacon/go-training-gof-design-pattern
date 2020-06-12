package proxy

import (
	"fmt"
	"time"
)

type Printable interface {
	SetPrinterName(name string)
	GetPrinterName() string
	Print(str string)
}

type Printer struct {
	name string
}

func NewPrinter(name string) *Printer {
	p := &Printer{name: name}
	p.heavyJob(fmt.Sprintf("Printerのインスタンス(%s)を生成中", name))
	return p
}

func (p *Printer) SetPrinterName(name string) {
	p.name = name
}

func (p *Printer) GetPrinterName() string {
	return p.name
}

func (p *Printer) Print(str string) {
	fmt.Printf("=== %s ===\n%s\n", p.name, str)
}

func (p *Printer) heavyJob(msg string) {
	fmt.Print(msg)
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Print(".")
	}
	fmt.Println("完了")
}

type PrinterProxy struct {
	name string
	real *Printer
}

func NewPrinterProxy(name string) *PrinterProxy {
	return &PrinterProxy{name: name}
}

func (p *PrinterProxy) SetPrinterName(name string) {
	if p.real != nil {
		p.real.SetPrinterName(name)
	}
	p.name = name
}

func (p *PrinterProxy) GetPrinterName() string {
	return p.name
}

func (p *PrinterProxy) Print(str string) {
	p.realize()
	p.real.Print(str)
}

func (p *PrinterProxy) realize() {
	if p.real == nil {
		p.real = NewPrinter(p.name)
	}
}
