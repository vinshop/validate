package validate

// Conditional condition validator
type Conditional struct {
	custom BoolFunc
	true   Rules
	false  Rules
}

func (c *Conditional) Do(data interface{}) error {
	if c.custom != nil {
		if c.custom() {
			return c.true.Do(data)
		} else {
			return c.false.Do(data)
		}
	}
	if !IsZero(data) {
		return c.true.Do(data)
	}
	return c.false.Do(data)
}

// If create new conditional for array, if data is non-zero then return Then, else return Else
func If(ok ...BoolFunc) *Conditional {
	c := &Conditional{
		true:  make(Rules, 0),
		false: make(Rules, 0),
	}
	if len(ok) > 0 {
		c.custom = ok[0]
	}

	return c
}

// Then validator if condition is non-zero
func (c *Conditional) Then(fns ...Rule) *Conditional {
	c.true = append(c.true, fns...)
	return c
}

// Else validator if condition is zero
func (c *Conditional) Else(fns ...Rule) *Conditional {
	c.false = append(c.false, fns...)
	return c
}

// SwitchCase switch case validator
type SwitchCase struct {
	cases  map[interface{}]Rules
	def    Rules
	custom []*caseCustom
}

type caseCustom struct {
	ok   BoolFunc
	rule Rule
}

func (s *SwitchCase) Do(data interface{}) error {
	r, ok := s.cases[Wrap(data).Data]
	if ok {
		return r.Do(data)
	}
	for _, custom := range s.custom {
		if custom.ok() {
			return custom.rule.Do(data)
		}
	}
	return s.def.Do(data)
}

// Switch create SwitchCase validator
func Switch() *SwitchCase {
	return &SwitchCase{
		cases: make(map[interface{}]Rules),
		def:   make(Rules, 0),
	}
}

// Case validator if value match case
func (s *SwitchCase) Case(value interface{}, rule ...Rule) *SwitchCase {
	s.cases[value] = append(s.cases[value], rule...)
	return s
}

// CaseCustom custom validate, has higher priority than Case
func (s *SwitchCase) CaseCustom(ok BoolFunc, rule ...Rule) *SwitchCase {
	s.custom = append(s.custom, &caseCustom{
		ok:   ok,
		rule: Rules(rule),
	})
	return s
}

// CaseMany same as Case but take multiple case
func (s *SwitchCase) CaseMany(value []interface{}, rule ...Rule) *SwitchCase {
	for _, v := range value {
		s.Case(v, rule...)
	}
	return s
}

// Default if not match any case, return default
func (s *SwitchCase) Default(rule ...Rule) *SwitchCase {
	s.def = append(s.def, rule...)
	return s
}
