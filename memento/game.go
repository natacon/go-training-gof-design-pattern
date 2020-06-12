package memento

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type memento struct {
	money  int
	fruits []string
}

func newMemento(money int) *memento {
	return &memento{money: money}
}

func (m *memento) Money() int {
	return m.money
}

func (m *memento) addFruits(fruit string) {
	m.fruits = append(m.fruits, fruit)
}

type Gamer struct {
	money      int
	fruits     []string
	fruitsName []string
}

func (g *Gamer) Money() int {
	return g.money
}

func NewGamer(money int) *Gamer {
	return &Gamer{
		money: money,
		fruitsName: []string{
			"りんご",
			"ぶどう",
			"ばなな",
			"みかん",
		},
	}
}

func (g *Gamer) Bet() {
	rand.Seed(time.Now().UnixNano())
	dice := rand.Intn(6) + 1
	if dice == 1 {
		g.money += 100
		fmt.Println("所持金が増えました。")
	} else if dice == 2 {
		g.money /= 2
		fmt.Println("所持金が半分になりました。")
	} else if dice == 6 {
		fruit := g.getFruit()
		fmt.Printf("フルーツ(%s)をもらいました。", fruit)
		g.fruits = append(g.fruits, fruit)
	} else {
		fmt.Println("何も起こりませんでした。")
	}
}

func (g *Gamer) CreateMemento() *memento {
	m := newMemento(g.money)
	for _, fruit := range g.fruits {
		if strings.HasPrefix(fruit, "おいしい") {
			m.addFruits(fruit)
		}
	}
	return m
}

func (g *Gamer) RestoreMemento(m *memento) {
	g.money = m.money
	g.fruits = m.fruits
}

func (g *Gamer) getFruit() string {
	prefix := ""
	if rand.Int()%2 == 0 {
		prefix = "おいしい"
	}
	return fmt.Sprintf("%s%s", prefix, g.fruitsName[rand.Intn(len(g.fruitsName))])
}

func (g *Gamer) String() string {
	return fmt.Sprintf("[money = %d, fruits = %v]", g.money, g.fruits)
}
