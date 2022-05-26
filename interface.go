package validate

type Validatable interface {
	Validate() error
}

type RuleFn func(data interface{}) error

func (v RuleFn) Do(data interface{}) error {
	return v(data)
}

type Rule interface {
	Do(data interface{}) error
}

type Rules []Rule

func (r Rules) Do(data interface{}) error {
	for _, fn := range r {
		if err := fn.Do(data); err != nil {
			return err
		}
	}
	return nil
}
