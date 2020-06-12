package strategy

import (
	"math/rand"
	"time"
)

type Strategy interface {
	nextHand() *Hand
	study(win bool)
}

type WinningStrategy struct {
	won      bool
	prevHand *Hand
}

func NewWinningStrategy() *WinningStrategy {
	return &WinningStrategy{}
}

func (s *WinningStrategy) nextHand() *Hand {
	if !s.won {
		rand.Seed(time.Now().UnixNano())
		s.prevHand = hand(rand.Intn(3))
	}
	return s.prevHand
}

func (s *WinningStrategy) study(win bool) {
	s.won = win
}

type RandomStrategy struct {
}

func NewRandomStrategy() *RandomStrategy {
	return &RandomStrategy{}
}

func (s *RandomStrategy) nextHand() *Hand {
	rand.Seed(time.Now().UnixNano())
	return hand(rand.Intn(3))
}

func (s *RandomStrategy) study(win bool) {
}
