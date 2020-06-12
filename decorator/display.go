package decorator

import "fmt"

type IDisplay interface {
	getColumns() int
	getRows() int
	getRowText(row int) string
}
type Display struct {
	display IDisplay
}

func NewDisplay() *Display {
	return &Display{}
}

func (d *Display) Show() {
	text := ""
	for i := 0; i < d.display.getRows(); i++ {
		text += fmt.Sprintf("%s\n", d.display.getRowText(i))
	}
	fmt.Print(text)
}

type StringDisplay struct {
	*Display
	text string
}

func NewStringDisplay(text string) *StringDisplay {
	stringDisplay := &StringDisplay{
		Display: NewDisplay(),
		text:    text,
	}
	stringDisplay.display = stringDisplay
	return stringDisplay
}

func (d *StringDisplay) getColumns() int {
	return len(d.text)
}

func (d *StringDisplay) getRows() int {
	return 1
}

func (d *StringDisplay) getRowText(row int) string {
	if row == 0 {
		return d.text
	}
	return ""
}

type MultiStringDisplay struct {
	*Display
	body    []string
	columns int
}

func (d *MultiStringDisplay) getColumns() int {
	return d.columns
}

func (d *MultiStringDisplay) getRows() int {
	return len(d.body)
}

func (d *MultiStringDisplay) getRowText(row int) string {
	return d.body[row]
}

func (d *MultiStringDisplay) Add(msg string) {
	d.body = append(d.body, msg)
	d.updateColumn(msg)
}

func (d *MultiStringDisplay) updateColumn(msg string) {
	if len(msg) > d.columns {
		d.columns = len(msg)
	}
	for row := 0; row < len(d.body); row++ {
		fills := d.columns - len(d.body[row])
		if fills > 0 {
			d.body[row] = d.body[row] + d.spaces(fills)
		}
	}
}

func (d *MultiStringDisplay) spaces(count int) string {
	buf := ""
	for i := 0; i < count; i++ {
		buf += " "
	}
	return buf
}

func NewMultiStringDisplay() *MultiStringDisplay {
	multiStringDisplay := &MultiStringDisplay{
		Display: NewDisplay(),
	}
	multiStringDisplay.display = multiStringDisplay
	return multiStringDisplay
}
