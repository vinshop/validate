package validate

// Valid main validator
type Valid struct {
	data        *Wrapper
	fns         Rules
	includePath bool
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

func (v *Valid) IncludePath() *Valid {
	v.includePath = true
	return v
}

func (v *Valid) Validate() error {
	for _, fn := range v.fns {
		if err := fn.Do(v.data); err != nil {
			if cErr, ok := err.(*Error); ok {
				if v.includePath {
					return cErr.IncludePath()
				}
				return cErr
			}

			return err
		}
	}
	return nil
}
