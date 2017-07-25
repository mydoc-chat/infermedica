package infermedica

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp_Lookup(t *testing.T) {
	a := GetTestAppInstance()
	res, err := a.Lookup("headache", SexFemale)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.ID)
	assert.NotEmpty(t, res.Label)
}
