package bridge

import (
	"math/rand"
	"time"
)

type Display struct {
	impl DisplayImpl
}

func NewDisplay(impl DisplayImpl) *Display {
	return &Display{impl: impl}
}

func (d *Display) open() {
	d.impl.rawOpen()
}

func (d *Display) print() {
	d.impl.rawPrint()
}

func (d *Display) close() {
	d.impl.rawClose()
}

func (d *Display) DisplayFunc() {
	d.open()
	d.print()
	d.close()
}

type CountDisplay struct {
	*Display
}

func NewCountDisplay(impl DisplayImpl) *CountDisplay {
	return &CountDisplay{Display: NewDisplay(impl)}
}

func (d *CountDisplay) MultiDisplay(times int) {
	d.open()
	for i := 0; i < times; i++ {
		d.print()
	}
	d.close()
}

type RandomDisplay struct {
	*CountDisplay
}

func NewRandomDisplay(impl DisplayImpl) *RandomDisplay {
	return &RandomDisplay{CountDisplay: NewCountDisplay(impl)}
}

func (d *RandomDisplay) RandomDisplay(times int) {
	d.open()
	rand.Seed(time.Now().UnixNano())
	randomTimes := rand.Intn(times)
	for i := 0; i < randomTimes; i++ {
		d.print()
	}
	d.close()
}

type IncreaseDisplay struct {
	*CountDisplay
	step int
}

func NewIncreaseDisplay(impl DisplayImpl, step int) *IncreaseDisplay {
	return &IncreaseDisplay{
		CountDisplay: NewCountDisplay(impl),
		step:         step,
	}
}

func (d *IncreaseDisplay) IncreaseDisplay(level int) {
	count := 0
	for i := 0; i < level; i++ {
		d.MultiDisplay(count)
		count += d.step
	}
}
