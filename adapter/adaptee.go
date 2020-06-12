package adapter

import "fmt"

type Banner struct {
	text string
}

func (b *Banner) ShowWithParen() {
	fmt.Printf("(%s)\n", b.text)
}

func (b *Banner) ShowWithAster() {
	fmt.Printf("*%s*\n", b.text)
}

func NewBanner(text string) *Banner {
	return &Banner{text: text}
}
