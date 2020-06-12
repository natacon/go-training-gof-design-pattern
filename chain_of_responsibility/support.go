package chain_of_responsibility

import "fmt"

type supportInterface interface {
	resolve(trouble *Trouble) bool
	Support(trouble *Trouble)
	SetNext(next supportInterface) supportInterface
}

type support struct {
	name string
	own  supportInterface
	next supportInterface
}

func (s *support) String() string {
	return fmt.Sprintf("[%s]", s.name)
}

func NewSupport(name string) *support {
	return &support{name: name}
}

func (s *support) SetNext(next supportInterface) supportInterface {
	s.next = next
	return next
}

func (s *support) Support(trouble *Trouble) {
	if s.own.resolve(trouble) {
		s.done(trouble)
	} else if s.next != nil {
		s.next.Support(trouble)
	} else {
		s.fail(trouble)
	}
}

func (s *support) done(trouble *Trouble) {
	fmt.Printf("%v is resolved by %s.\n", trouble, s)
}

func (s *support) fail(trouble *Trouble) {
	fmt.Printf("%v is cannot be resolved.\n", trouble)
}

type NoSupport struct {
	*support
}

func NewNoSupport(name string) *NoSupport {
	support := &NoSupport{
		support: NewSupport(name),
	}
	support.own = support
	return support
}

func (s *NoSupport) resolve(trouble *Trouble) bool {
	return false
}

type LimitSupport struct {
	*support
	limit int
}

func NewLimitSupport(name string, limit int) *LimitSupport {
	support := &LimitSupport{
		support: NewSupport(name),
		limit:   limit,
	}
	support.own = support
	return support
}

func (s *LimitSupport) resolve(trouble *Trouble) bool {
	if trouble.Number() < s.limit {
		return true
	}
	return false
}

type OddSupport struct {
	*support
}

func NewOddSupport(name string) *OddSupport {
	support := &OddSupport{
		support: NewSupport(name),
	}
	support.own = support
	return support
}

func (s *OddSupport) resolve(trouble *Trouble) bool {
	if trouble.Number()%2 == 1 {
		return true
	}
	return false
}

type SpecialSupport struct {
	*support
	number int
}

func NewSpecialSupport(name string, number int) *SpecialSupport {
	support := &SpecialSupport{
		support: NewSupport(name),
		number:  number,
	}
	support.own = support
	return support
}

func (s *SpecialSupport) resolve(trouble *Trouble) bool {
	if trouble.number == s.number {
		return true
	}
	return false
}
