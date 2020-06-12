package builder

import (
	"fmt"
	"os"
	"strings"
)

type Builder interface {
	MakeTitle(title string)
	MakeString(str string)
	MakeItems(items []string)
	Close()
}

type TextBuilder struct {
	buffer []string
}

func NewTextBuilder() *TextBuilder {
	return &TextBuilder{}
}

func (t *TextBuilder) MakeTitle(title string) {
	t.buffer = append(t.buffer, "======================\n")
	t.buffer = append(t.buffer, fmt.Sprintf("[%s]", title))
}

func (t *TextBuilder) MakeString(str string) {
	t.buffer = append(t.buffer, fmt.Sprintf("■%s", str))
}

func (t *TextBuilder) MakeItems(items []string) {
	for _, item := range items {
		t.buffer = append(t.buffer, fmt.Sprintf("  ・%s", item))
	}
}

func (t *TextBuilder) Close() {
	t.buffer = append(t.buffer, "======================\n")
}

func (t *TextBuilder) Result() string {
	return strings.Join(t.buffer, "")
}

type HTMLBuilder struct {
	filename    string
	initialized bool
}

func NewHTMLBuilder() *HTMLBuilder {
	return &HTMLBuilder{}
}

func (b *HTMLBuilder) MakeTitle(title string) {
	if b.initialized {
		return
	}
	b.filename = fmt.Sprintf("%s.html", title)
	f, _ := os.Create(b.filename)
	defer f.Close()
	f.Write([]byte(fmt.Sprintf("<html><head><title>%s</title></head><body>", title)))
	f.Write([]byte(fmt.Sprintf("<h1>%s</h1>", title)))
	b.initialized = true
}

func (b *HTMLBuilder) MakeString(str string) {
	if !b.initialized {
		return
	}
	f, _ := os.OpenFile(b.filename, os.O_WRONLY|os.O_APPEND, 0666)
	defer f.Close()
	f.Write([]byte(fmt.Sprintf("<p>%s</p>", str)))
}

func (b *HTMLBuilder) MakeItems(items []string) {
	if !b.initialized {
		return
	}
	f, _ := os.OpenFile(b.filename, os.O_WRONLY|os.O_APPEND, 0666)
	defer f.Close()
	f.Write([]byte("<ul>"))
	for _, item := range items {
		f.Write([]byte(fmt.Sprintf("<li>%s</li>", item)))
	}
	f.Write([]byte("</ul>"))
}

func (b *HTMLBuilder) Close() {
	if !b.initialized {
		return
	}
	f, _ := os.OpenFile(b.filename, os.O_WRONLY|os.O_APPEND, 0666)
	defer f.Close()
	f.Write([]byte("</body></html>"))
}

func (b *HTMLBuilder) Result() string {
	return b.filename
}
