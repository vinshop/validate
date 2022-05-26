package validate

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestError_GetRootError(t *testing.T) {

	e := FieldError("abc", reflect.StructField{Tag: `json:"a"`}, FieldError("B", reflect.StructField{Tag: `json:"b"`}, ArrayError(10, ErrNotArray)))
	path, err := e.(*Error).GetRootError()
	assert.Equal(t, "a.b[10]", path)
	assert.Equal(t, ErrNotArray, err)
	fmt.Println(err)
}

func TestError_GetLastErrorWithKey(t *testing.T) {
	e := FieldError("abc", reflect.StructField{Tag: `json:"a"`}, FieldError("B", reflect.StructField{Tag: `json:"b"`}, ArrayError(10, ErrNotArray)))
	err := e.(*Error).GetLastErrorWithKey().(*Error)
	assert.Equal(t, "B", err.Key)
	fmt.Println(err)
}
