package validate

type Valid struct {
	data interface{}
	fns  Rules
}

func (v *Valid) Do(_ interface{}) error {
	return v.Validate()
}

func Use(data interface{}, fns ...Rule) *Valid {
	return &Valid{
		data: data,
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
