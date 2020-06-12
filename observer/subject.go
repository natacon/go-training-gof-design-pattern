package observer

import (
	"math/rand"
	"time"
)

type NumberGenerator struct {
	observers []Observer
	number    int
}

func NewNumberGenerator(number int) *NumberGenerator {
	return &NumberGenerator{number: number}
}

func (g *NumberGenerator) Number() int {
	return g.number
}

func (g *NumberGenerator) NotifyObservers() {
	for _, observer := range g.observers {
		observer.Update(g)
	}
}

func (g *NumberGenerator) AddObserver(observer Observer) {
	g.observers = append(g.observers, observer)
}

type RandomNumberGenerator struct {
	*NumberGenerator
}

func NewRandomNumberGenerator() *RandomNumberGenerator {
	return &RandomNumberGenerator{
		NumberGenerator: NewNumberGenerator(0),
	}
}

func (g *RandomNumberGenerator) Execute() {
	for i := 0; i < 20; i++ {
		rand.Seed(time.Now().UnixNano())
		g.number = rand.Intn(50)
		g.NotifyObservers()
	}
}

type IncrementalNumberGenerator struct {
	*NumberGenerator
	end int
	inc int
}

func NewIncrementalNumberGenerator(number int, end int, inc int) *IncrementalNumberGenerator {
	return &IncrementalNumberGenerator{
		NumberGenerator: NewNumberGenerator(number),
		end:             end,
		inc:             inc,
	}
}

func (g *IncrementalNumberGenerator) Execute() {
	for g.number < g.end {
		g.NotifyObservers()
		g.number += g.inc
	}
}
