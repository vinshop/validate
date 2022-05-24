package validate

type Valid struct {
	data interface{}
	fns  []Validate
}

func Use(data interface{}, fns ...Validate) *Valid {
	return &Valid{
		data: data,
		fns:  fns,
	}
}

func (v *Valid) Validate() error {
	for _, fn := range v.fns {
		if err := fn(v.data); err != nil {
			return err
		}
	}
	return nil
}
