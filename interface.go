package validate

type Validatable interface {
	Validate() error
}

// RuleFn rule as func
type RuleFn func(data interface{}) error

// Do execute the logic
func (v RuleFn) Do(data interface{}) error {
	return v(data)
}

// Rule common interface
type Rule interface {
	Do(data interface{}) error
}

// Rules array of Rule
type Rules []Rule

func (r Rules) Do(data interface{}) error {
	for _, fn := range r {
		if err := fn.Do(data); err != nil {
			return err
		}
	}
	return nil
}
