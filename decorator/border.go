package decorator

import "strings"

type Border struct {
	*Display
	borderDisplay IDisplay
}

func NewBorder(display IDisplay) *Border {
	return &Border{
		borderDisplay: display,
		Display:       NewDisplay(),
	}
}

type SideBorder struct {
	*Border
	borderChar string
}

func (b *SideBorder) getColumns() int {
	return 1 + b.borderDisplay.getColumns() + 1
}

func (b *SideBorder) getRows() int {
	return b.borderDisplay.getRows()
}

func (b *SideBorder) getRowText(row int) string {
	return strings.Join([]string{b.borderChar, b.borderDisplay.getRowText(row), b.borderChar}, "")
}

func NewSideBorder(display IDisplay, borderChar string) *SideBorder {
	sideBorder := &SideBorder{
		Border:     NewBorder(display),
		borderChar: borderChar,
	}
	sideBorder.display = sideBorder
	return sideBorder
}

type FullBorder struct {
	*Border
}

func NewFullBorder(display IDisplay) *FullBorder {
	fullBorder := &FullBorder{Border: NewBorder(display)}
	fullBorder.display = fullBorder
	return fullBorder
}

func (b *FullBorder) getColumns() int {
	return 1 + b.borderDisplay.getColumns() + 1
}

func (b *FullBorder) getRows() int {
	return 1 + b.borderDisplay.getRows() + 1
}

func (b *FullBorder) getRowText(row int) string {
	if row == 0 {
		return "+" + b.makeLine("-", b.borderDisplay.getColumns()) + "+"
	} else if row == b.borderDisplay.getRows()+1 {
		return "+" + b.makeLine("-", b.borderDisplay.getColumns()) + "+"
	}
	return "|" + b.borderDisplay.getRowText(row-1) + "|"
}

func (b *FullBorder) makeLine(char string, count int) string {
	buf := ""
	for i := 0; i < count; i++ {
		buf += char
	}
	return buf
}

type UpDownBorder struct {
	*Border
	borderChar string
}

func NewUpDownBorder(display IDisplay, borderChar string) *UpDownBorder {
	upDownBorder := &UpDownBorder{
		Border:     NewBorder(display),
		borderChar: borderChar,
	}
	upDownBorder.display = upDownBorder
	return upDownBorder
}

func (d *UpDownBorder) getColumns() int {
	return d.borderDisplay.getColumns()
}

func (d *UpDownBorder) getRows() int {
	return 1 + d.borderDisplay.getRows() + 1
}

func (d *UpDownBorder) getRowText(row int) string {
	if row == 0 || row == (d.getRows()-1) {
		return d.makeLine(d.borderChar, d.getColumns())
	}
	return d.borderDisplay.getRowText(row - 1)
}

func (d *UpDownBorder) makeLine(ch string, count int) string {
	buf := ""
	for i := 0; i < count; i++ {
		buf += ch
	}
	return buf
}
