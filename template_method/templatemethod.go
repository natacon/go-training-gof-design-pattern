package template_method

import "fmt"

type IDisplay interface {
	open()
	print()
	close()
}

type AbstractDisplay struct {
	printer IDisplay
}

func (d *AbstractDisplay) Display() {
	d.printer.open()
	for i := 0; i < 5; i++ {
		d.printer.print()
	}
	d.printer.close()
}

type CharDisplay struct {
	*AbstractDisplay
	ch byte
}

func NewCharDisplay(ch byte) *CharDisplay {
	charDisplay := &CharDisplay{ch: ch, AbstractDisplay: &AbstractDisplay{}}
	charDisplay.printer = charDisplay
	return charDisplay
}

func (c *CharDisplay) open() {
	fmt.Print("<<")
}

func (c *CharDisplay) close() {
	fmt.Print(">>\n")
}

func (c *CharDisplay) print() {
	fmt.Print(string(c.ch))
}

type StringDisplay struct {
	*AbstractDisplay
	text  string
	width int
}

func NewStringDisplay(text string) *StringDisplay {
	stringDisplay := &StringDisplay{
		AbstractDisplay: &AbstractDisplay{},
		text:            text,
		width:           len(text),
	}
	stringDisplay.printer = stringDisplay
	return stringDisplay
}

func (d StringDisplay) open() {
	d.printLine()
}

func (d StringDisplay) print() {
	fmt.Printf("|%s|\n", d.text)
}

func (d StringDisplay) close() {
	d.printLine()
}

func (d *StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < d.width; i++ {
		fmt.Print("-")
	}
	fmt.Print("+\n")
}
