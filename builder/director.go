package builder

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() {
	d.builder.MakeTitle("Greeting")
	d.builder.MakeString("朝から昼にかけて")
	d.builder.MakeItems([]string{
		"おはよう", "こんにちは",
	})
	d.builder.MakeString("夜に")
	d.builder.MakeItems([]string{
		"こんばんは", "おやすみ", "さよなら",
	})
	d.builder.Close()
}
