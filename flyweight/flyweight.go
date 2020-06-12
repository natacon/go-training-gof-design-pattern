package flyweight

import (
	"fmt"
	"os"
)

type bigChar struct {
	fontData string
}

func newBigChar(charName string) *bigChar {
	char := &bigChar{}
	data := make([]byte, 256)
	f, err := os.Open(fmt.Sprintf("flyweight/big%s.txt", charName))
	if err != nil {
		char.fontData = fmt.Sprintf("%s?", charName)
	} else {
		_, err := f.Read(data)
		if err != nil {
			fmt.Println(err)
		}
		char.fontData = string(data)
	}
	return char
}

func (b *bigChar) print() {
	fmt.Println(b.fontData)
}

type bigCharFactory struct {
	pool map[string]*bigChar
}

func newBigCharFactory() *bigCharFactory {
	return &bigCharFactory{
		pool: make(map[string]*bigChar),
	}
}

var bigCharFactoryInstance = newBigCharFactory()

func BigCharFactory() *bigCharFactory {
	return bigCharFactoryInstance
}

func (f *bigCharFactory) BigChar(charName string) *bigChar {
	bc, ok := f.pool[charName]
	if !ok {
		bc = newBigChar(charName)
		f.pool[charName] = bc
	}
	return bc
}

type BigString struct {
	bigChars []*bigChar
}

func NewBigString(str string) *BigString {
	bs := &BigString{}
	factory := BigCharFactory()
	for _, s := range str {
		bs.bigChars = append(bs.bigChars, factory.BigChar(string(s)))
	}
	return bs
}

func (b *BigString) Print() {
	for _, bc := range b.bigChars {
		bc.print()
	}
}
