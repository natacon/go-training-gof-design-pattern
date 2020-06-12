package observer

import (
	"fmt"
	"strings"
	"time"
)

type Observer interface {
	Update(generator *NumberGenerator)
}

type DigitObserver struct {
}

func NewDigitObserver() *DigitObserver {
	return &DigitObserver{}
}

func (o *DigitObserver) Update(generator *NumberGenerator) {
	fmt.Printf("DigitObserver:%d\n", generator.Number())
	time.Sleep(time.Millisecond * 100)
}

type GraphObserver struct {
}

func (o *GraphObserver) Update(generator *NumberGenerator) {
	fmt.Print("GraphObserver:")
	fmt.Print(strings.Repeat("*", generator.Number()))
	fmt.Println("")
	time.Sleep(time.Millisecond * 100)
}

func NewGraphObserver() *GraphObserver {
	return &GraphObserver{}
}
