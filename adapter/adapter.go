package adapter

type PrintBanner struct {
	*Banner
}

func NewPrintBanner(text string) *PrintBanner {
	return &PrintBanner{NewBanner(text)}
}

func (p PrintBanner) PrintWeak() {
	p.ShowWithParen()
}

func (p PrintBanner) PrintStrong() {
	p.ShowWithAster()
}
