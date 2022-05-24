package validate

type Validatable interface {
	Validate() error
}

type Validate func(data interface{}) error
