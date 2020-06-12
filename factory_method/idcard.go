package factory_method

import "fmt"

// TODO: 問題4-2をやる

type IDCard struct {
	owner string
}

func NewIDCard(owner string) *IDCard {
	return &IDCard{owner: owner}
}

func (c *IDCard) Owner() string {
	return c.owner
}

func (c *IDCard) Use() {
	fmt.Printf("%sのカードを作ります。\n", c.owner)
}

type IDCardFactory struct {
	*Factory
	owners []string
}

func (f *IDCardFactory) Owners() []string {
	return f.owners
}

func NewIDCardFactory() *IDCardFactory {
	f := &IDCardFactory{
		Factory: &Factory{},
	}
	f.factory = f
	return f
}
func (f *IDCardFactory) CreateProduct(owner string) Product {
	return NewIDCard(owner)
}

func (f *IDCardFactory) RegisterProduct(product Product) {
	// TODO: 型アサーションを復習
	// とりあえずこれを参考にした→ https://stackoverflow.com/questions/42773848/impossible-type-assertions-with-casting-from-interface-type-to-the-actual-type
	card := product.(*IDCard)
	f.owners = append(f.owners, card.Owner())
}
