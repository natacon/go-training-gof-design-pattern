package abstract_factory

import (
	"fmt"
	"os"
)

type IItem interface {
	makeHTML() string
}

type Item struct {
	item    IItem
	Caption string
}

func NewItem(caption string) *Item {
	return &Item{Caption: caption}
}

type ILink interface {
	IItem
}
type Link struct {
	*Item
	url string
}

func NewLink(caption, url string) *Link {
	return &Link{
		Item: NewItem(caption),
		url:  url,
	}
}

type ITray interface {
	IItem
	Add(item IItem)
}

type Tray struct {
	*Item
	tray []IItem
}

func (t *Tray) Add(item IItem) {
	t.tray = append(t.tray, item)
}

func NewTray(caption string) *Tray {
	return &Tray{
		Item: NewItem(caption),
	}
}

type IPage interface {
	Add(item IItem)
	Output(page IPage)
	makeHTML() string
}

type Page struct {
	title   string
	author  string
	content []IItem
}

func NewPage(title string, author string) *Page {
	return &Page{title: title, author: author}
}

func (p *Page) Add(item IItem) {
	p.content = append(p.content, item)
}

func (p *Page) Output(page IPage) {
	filename := fmt.Sprintf("%s.html", p.title)
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write([]byte(page.makeHTML()))
}

type Factory interface {
	CreateLink(caption, url string) ILink
	CreateTray(caption string) ITray
	CreatePage(title, author string) IPage
}
