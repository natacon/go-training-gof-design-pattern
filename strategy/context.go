package strategy

import "fmt"

type Player struct {
	name      string
	strategy  Strategy
	winCount  int
	loseCount int
	gameCount int
}

func NewPlayer(name string, strategy Strategy) *Player {
	return &Player{name: name, strategy: strategy}
}

func (p *Player) NextHand() *Hand {
	return p.strategy.nextHand()
}

func (p *Player) Win() {
	p.strategy.study(true)
	p.winCount++
	p.gameCount++
}

func (p *Player) Lose() {
	p.strategy.study(false)
	p.loseCount++
	p.gameCount++
}

func (p *Player) Even() {
	p.gameCount++
}

func (p *Player) String() string {
	return fmt.Sprintf("[%s:%d games, %d win, %d lose]", p.name, p.gameCount, p.winCount, p.loseCount)
}
