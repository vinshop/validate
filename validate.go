package validate

// Valid main validator
type Valid struct {
	data *Wrapper
	fns  Rules
}

func (v *Valid) Do(_ interface{}) error {
	return v.Validate()
}

// Use add Rule for data
func Use(data interface{}, fns ...Rule) *Valid {
	return &Valid{
		data: Wrap(data),
		fns:  fns,
	}
}

func (v *Valid) Validate() error {
	for _, fn := range v.fns {
		if err := fn.Do(v.data); err != nil {
			return err
		}
	}
	return nil
}
