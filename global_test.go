package validate

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSetNumberEpsilon(t *testing.T) {
	ep := rand.Float64()
	SetNumberEpsilon(ep)
	assert.Equal(t, ep, NumberEpsilon)
}
