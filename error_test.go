package validate

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError_GetRootError(t *testing.T) {
	e := FieldError("abc", "A", FieldError("B", "B", ArrayError(10, ErrNotArray)))
	path, err := e.(*Error).GetRootError()
	assert.Equal(t, "A.B[10]", path)
	assert.Equal(t, ErrNotArray, err)
	fmt.Println(err)
}

func TestError_GetLastErrorWithKey(t *testing.T) {
	e := FieldError("abc", "A", FieldError("B", "B", ArrayError(10, ErrNotArray)))
	err := e.(*Error).GetLastErrorWithKey().(*Error)
	assert.Equal(t, "B", err.Key)
	fmt.Println(err)
}
