package singleton

var instance = &singleton{}

type singleton struct {
	i int
}

func (s *singleton) SetI(i int) {
	s.i = i
}

func GetInstance() *singleton {
	return instance
}
