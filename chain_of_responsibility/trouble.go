package chain_of_responsibility

import "fmt"

type Trouble struct {
	number int
}

func (t *Trouble) String() string {
	return fmt.Sprintf("[Trouble %d]", t.number)
}

func (t *Trouble) Number() int {
	return t.number
}

func NewTrouble(number int) *Trouble {
	return &Trouble{number: number}
}
