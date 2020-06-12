package bridge

import (
	"fmt"
	"os"
	"strings"
)

type DisplayImpl interface {
	rawOpen()
	rawPrint()
	rawClose()
}

type StringDisplay struct {
	text  string
	width int
}

func NewStringDisplay(text string) *StringDisplay {
	return &StringDisplay{
		text:  text,
		width: len(text),
	}
}

func (d *StringDisplay) rawOpen() {
	d.printLine()
}

func (d *StringDisplay) rawPrint() {
	fmt.Printf("|%s|\n", d.text)
}

func (d *StringDisplay) rawClose() {
	d.printLine()
}

func (d *StringDisplay) printLine() {
	fmt.Print("+")
	fmt.Print(strings.Repeat("-", d.width))
	fmt.Println("+")
}

type TextFileDisplay struct {
	filename string
	file     *os.File
}

func (d *TextFileDisplay) rawOpen() {
	f, _ := os.Open(d.filename)
	d.file = f
}

func (d *TextFileDisplay) rawPrint() {
	buf := make([]byte, 64)
	d.file.Read(buf)
	fmt.Printf("%s\n", buf)
}

func (d *TextFileDisplay) rawClose() {
	d.file.Close()
}

func NewTextFileDisplay(filename string) *TextFileDisplay {
	return &TextFileDisplay{filename: filename}
}

type CharDisplayImpl struct {
	head byte
	body byte
	foot byte
}

func NewCharDisplayImpl(head byte, body byte, foot byte) *CharDisplayImpl {
	return &CharDisplayImpl{head: head, body: body, foot: foot}
}

func (d *CharDisplayImpl) rawOpen() {
	fmt.Print(string(d.head))
}

func (d *CharDisplayImpl) rawPrint() {
	fmt.Print(string(d.body))
}

func (d *CharDisplayImpl) rawClose() {
	fmt.Println(string(d.foot))
}
