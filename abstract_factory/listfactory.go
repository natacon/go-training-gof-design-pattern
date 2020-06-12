package abstract_factory

import (
	"fmt"
	"strings"
)

type ListFactory struct {
}

func NewListFactory() *ListFactory {
	return &ListFactory{}
}

func (f *ListFactory) CreateLink(caption, url string) ILink {
	return NewListLink(caption, url)
}

func (f *ListFactory) CreateTray(caption string) ITray {
	return NewListTray(caption)
}

func (f *ListFactory) CreatePage(title, author string) IPage {
	return NewListPage(title, author)
}

type ListLink struct {
	*Link
}

func NewListLink(caption, url string) *ListLink {
	return &ListLink{
		Link: NewLink(caption, url),
	}
}

func (l *ListLink) makeHTML() string {
	return fmt.Sprintf(`<li><a href="%s">%s</a></li>`, l.url, l.Caption)
}

type ListTray struct {
	*Tray
}

func NewListTray(caption string) *ListTray {
	return &ListTray{
		Tray: NewTray(caption),
	}
}

func (l *ListTray) makeHTML() string {
	var buf []string
	buf = append(buf, "<li>")
	buf = append(buf, fmt.Sprintf("%s", l.Caption))
	buf = append(buf, "<ul>")
	for _, item := range l.tray {
		buf = append(buf, item.makeHTML())
	}
	buf = append(buf, "</ul>")
	buf = append(buf, "</li>")
	return strings.Join(buf, "")
}

type ListPage struct {
	*Page
}

func NewListPage(title, author string) *ListPage {
	return &ListPage{
		Page: NewPage(title, author),
	}
}

func (l *ListPage) makeHTML() string {
	buf := "<html>"
	buf += fmt.Sprintf("<head><title>%s</title></head>", l.title)
	buf += "<body>"
	buf += fmt.Sprintf("<h1>%s</h1>", l.title)
	buf += "<ul>"
	for _, item := range l.content {
		buf += item.makeHTML()
	}
	buf += "</ul>"
	buf += fmt.Sprintf("<hr><adress>%s</adress>", l.author)
	buf += "</body></html>"
	return buf
}
