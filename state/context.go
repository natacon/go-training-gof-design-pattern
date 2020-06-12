package state

import "fmt"

type context interface {
	SetClock(hour int)
	ChangeState(state state)
	CallSecurityCenter(msg string)
	RecordLog(msg string)
}

type safeFrame struct {
	// State 本当はprivateであるべき。main関数から呼びたいのでpublicにした
	State       state
	title       string
	textClock   *TextField
	TextScreen  *TextArea
	buttonUse   *Button
	buttonAlarm *Button
	buttonPhone *Button
	buttonExit  *Button
}

func NewSafeFrame(title string) *safeFrame {
	f := &safeFrame{
		State:       DayState(),
		title:       title,
		textClock:   NewTextField(60),
		TextScreen:  NewTextArea(10, 60),
		buttonUse:   NewButton("金庫使用"),
		buttonAlarm: NewButton("非常ベル"),
		buttonPhone: NewButton("通常通話"),
		buttonExit:  NewButton("終了"),
	}
	return f
}

//// actionPerformed ボタンが押されたら呼ばれる
//func (f *safeFrame) actionPerformed() {
//	// ボタンごとの分岐を適当に表現する → mainから各メソッドを呼ぶことにした
//}

func (f *safeFrame) SetClock(hour int) {
	clockString := "現在時刻は"
	if hour < 10 {
		clockString += fmt.Sprintf("0%d:00", hour)
	} else {
		clockString += fmt.Sprintf("%d:00", hour)
	}
	fmt.Println(clockString)
	f.textClock.SetText(clockString)
	f.State.DoClock(f, hour)
}

func (f *safeFrame) ChangeState(state state) {
	fmt.Printf("%sから%sへ状態が変化しました。", state.stateString(), f.State.stateString())
	f.State = state
}

func (f *safeFrame) CallSecurityCenter(msg string) {
	f.TextScreen.addText(fmt.Sprintf("call! %s\n", msg))
}

func (f *safeFrame) RecordLog(msg string) {
	f.TextScreen.addText(fmt.Sprintf("record ... %s\n", msg))
}

type TextField struct {
	width int
	text  string
}

func (t *TextField) SetText(text string) {
	t.text = text
}

func NewTextField(width int) *TextField {
	return &TextField{width: width}
}

type TextArea struct {
	height, width int
	text          []string
}

func NewTextArea(height int, width int) *TextArea {
	return &TextArea{height: height, width: width}
}

func (t *TextArea) addText(text string) {
	t.text = append(t.text, text)
}

type Button struct {
	caption string
}

func NewButton(caption string) *Button {
	return &Button{caption: caption}
}
