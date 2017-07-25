package infermedica

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp_Conditions(t *testing.T) {
	a := GetTestAppInstance()
	res, err := a.Conditions()
	assert.NoError(t, err)
	assert.NotEqual(t, len(*res), 0)
	assert.NotEmpty(t, (*res)[0].ID)
	assert.NotEmpty(t, (*res)[0].Name)
	assert.NotEmpty(t, (*res)[0].CommonName)

	testConditionByID(a, (*res)[0], t)
}

func testConditionByID(a App, r ConditionRes, t *testing.T) {
	res, err := a.ConditionByID(r.ID)
	assert.NoError(t, err)
	assert.Equal(t, res.Name, r.Name)
	assert.Equal(t, res.CommonName, r.CommonName)
}
