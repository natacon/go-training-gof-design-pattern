package state

type state interface {
	DoClock(context context, hour int)
	DoUse(context context)
	DoAlarm(context context)
	DoPhone(context context)
	stateString() string
}

type dayState struct {
}

var dayStateInstance = &dayState{}

func (s *dayState) DoClock(context context, hour int) {
	if hour < 9 || 17 <= hour {
		context.ChangeState(NightState())
	}
}

func (s *dayState) DoUse(context context) {
	context.RecordLog("金庫使用（昼間）")
}

func (s *dayState) DoAlarm(context context) {
	context.CallSecurityCenter("非常ベル（昼間）")
}

func (s *dayState) DoPhone(context context) {
	context.CallSecurityCenter("通常の通話（昼間）")
}

func (s *dayState) stateString() string {
	return "[昼間]"
}

func DayState() *dayState {
	return dayStateInstance
}

type nightState struct {
}

func (s *nightState) DoClock(context context, hour int) {
	if 9 <= hour && hour < 17 {
		context.ChangeState(DayState())
	}
}

func (s *nightState) DoUse(context context) {
	context.CallSecurityCenter("非常：夜間の金庫使用！")
}

func (s *nightState) DoAlarm(context context) {
	context.CallSecurityCenter("非常ベル（夜間）")
}

func (s *nightState) DoPhone(context context) {
	context.CallSecurityCenter("夜間の通話録音")
}

func (s *nightState) stateString() string {
	return "[夜間]"
}

var nightStateInstance = &nightState{}

func NightState() *nightState {
	return nightStateInstance
}
