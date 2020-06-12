package mediator

import "fmt"

type Mediator interface {
	createColleagues()
	colleagueChanged()
}

type LoginFrame struct {
	CheckGuest   *ColleagueCheckbox
	CheckLogin   *ColleagueCheckbox
	TextUser     *ColleagueTextField
	TextPass     *ColleagueTextField
	ButtonOk     *ColleagueButton
	ButtonCancel *ColleagueButton
}

func (f *LoginFrame) Print() {
	fmt.Printf("%+v%+v\n%+v%+v\n%+v%+v\n%+v%+v\n%+v%+v\n%+v%+v\n\n",
		f.CheckGuest.Component, f.CheckGuest,
		f.CheckLogin.Component, f.CheckLogin,
		f.TextUser.Component, f.TextUser,
		f.TextPass.Component, f.TextPass,
		f.ButtonOk.Component, f.ButtonOk,
		f.ButtonCancel.Component, f.ButtonCancel,
	)
}

func NewLoginFrame(title string) *LoginFrame {
	lf := &LoginFrame{}
	lf.createColleagues()
	lf.colleagueChanged()
	return lf
}

func (f *LoginFrame) createColleagues() {
	f.CheckGuest = NewColleagueCheckbox("Guest", true)
	f.CheckLogin = NewColleagueCheckbox("Login", false)
	f.TextUser = NewColleagueTextField("", 10)
	f.TextPass = NewColleagueTextField("", 10)
	f.ButtonOk = NewColleagueButton("OK")
	f.ButtonCancel = NewColleagueButton("Cancel")

	f.CheckGuest.SetMediator(f)
	f.CheckLogin.SetMediator(f)
	f.TextUser.SetMediator(f)
	f.TextPass.SetMediator(f)
	f.ButtonOk.SetMediator(f)
	f.ButtonCancel.SetMediator(f)
}

func (f *LoginFrame) colleagueChanged() {
	f.ButtonCancel.setColleagueEnabled(true)
	if f.CheckGuest.State() {
		f.TextUser.setColleagueEnabled(false)
		f.TextPass.setColleagueEnabled(false)
		f.ButtonOk.setColleagueEnabled(true)
	} else {
		f.TextUser.setColleagueEnabled(true)
		f.userpassChanged()
	}
}

func (f *LoginFrame) userpassChanged() {
	if f.TextUser.Text() != "" {
		f.TextPass.setColleagueEnabled(true)
		if len(f.TextUser.Text()) >= 4 &&
			len(f.TextPass.Text()) >= 4 {
			f.ButtonOk.setColleagueEnabled(true)
		} else {
			f.ButtonOk.setColleagueEnabled(false)
		}
	} else {
		f.TextPass.setColleagueEnabled(false)
		f.ButtonOk.setColleagueEnabled(false)
	}
}
