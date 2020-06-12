package prototype

import (
	"fmt"
	"strings"
)

// Product is interface for prototype
type Product interface {
	Use(s string)
	CreateClone() Product
}

type MessageBox struct {
	decochar byte
}

func NewMessageBox(decochar byte) *MessageBox {
	return &MessageBox{decochar: decochar}
}

func (m *MessageBox) Use(s string) {
	length := len(s)
	line := strings.Repeat(string(m.decochar), length+4)
	fmt.Println(line)
	fmt.Printf("%s %s %s\n", string(m.decochar), s, string(m.decochar))
	fmt.Println(line)
}

func (m *MessageBox) CreateClone() Product {
	return NewMessageBox(m.decochar)
}

type UnderlinePen struct {
	ulchar byte
}

func NewUnderlinePen(ulchar byte) *UnderlinePen {
	return &UnderlinePen{ulchar: ulchar}
}

func (u *UnderlinePen) Use(s string) {
	length := len(s)
	fmt.Printf("\"%s\"\n ", s)
	line := strings.Repeat(string(u.ulchar), length)
	fmt.Println(line)
}

func (u *UnderlinePen) CreateClone() Product {
	return NewUnderlinePen(u.ulchar)
}
