package mediator

// Component もともとの例がjava.awtパッケージを利用したものなので、
// 画面部品を擬似的に表現する。
type Component struct {
	enabled         bool
	backgroundColor string
}

//func (c *Component) String() string {
//	return fmt.Sprintf("%+v", c)
//}

func NewComponent() *Component {
	return &Component{}
}

func (c *Component) SetBackgroundColor(backgroundColor string) {
	c.backgroundColor = backgroundColor
}

func (c *Component) SetEnabled(enabled bool) {
	c.enabled = enabled
}

type Colleague interface {
	SetMediator(mediator Mediator)
	setColleagueEnabled(enabled bool)
}

type ColleagueButton struct {
	*Component
	mediator Mediator
	caption  string
}

func NewColleagueButton(caption string) *ColleagueButton {
	return &ColleagueButton{
		Component: NewComponent(),
		caption:   caption,
	}
}

func (c *ColleagueButton) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *ColleagueButton) setColleagueEnabled(enabled bool) {
	c.SetEnabled(enabled)
}

type ColleagueTextField struct {
	*Component
	mediator Mediator
	text     string
	columns  int
}

func (c *ColleagueTextField) SetText(text string) {
	c.text = text
}

func (c *ColleagueTextField) Text() string {
	return c.text
}

func NewColleagueTextField(text string, columns int) *ColleagueTextField {
	return &ColleagueTextField{
		Component: NewComponent(),
		text:      text,
		columns:   columns,
	}
}

func (c *ColleagueTextField) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *ColleagueTextField) setColleagueEnabled(enabled bool) {
	c.SetEnabled(enabled)
	if enabled {
		c.SetBackgroundColor("White")
	} else {
		c.SetBackgroundColor("LightGray")
	}
}

func (c *ColleagueTextField) TextValueChanged() {
	c.mediator.colleagueChanged()
}

type ColleagueCheckbox struct {
	*Component
	mediator Mediator
	caption  string
	state    bool
}

func (c *ColleagueCheckbox) SetState(state bool) {
	c.state = state
}

func (c *ColleagueCheckbox) Caption() string {
	return c.caption
}

func (c *ColleagueCheckbox) State() bool {
	return c.state
}

func NewColleagueCheckbox(caption string, state bool) *ColleagueCheckbox {
	return &ColleagueCheckbox{
		Component: NewComponent(),
		caption:   caption,
		state:     state,
	}
}

func (c *ColleagueCheckbox) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *ColleagueCheckbox) setColleagueEnabled(enabled bool) {
	c.SetEnabled(enabled)
}

func (c *ColleagueCheckbox) ItemStateChanged() {
	c.mediator.colleagueChanged()
}
