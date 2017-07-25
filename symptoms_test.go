package infermedica

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp_Symptoms(t *testing.T) {
	a := GetTestAppInstance()
	res, err := a.Symptoms()
	assert.NoError(t, err)
	assert.NotEqual(t, len(*res), 0)
	assert.NotEmpty(t, (*res)[0].ID)
	assert.NotEmpty(t, (*res)[0].Name)
	assert.NotEmpty(t, (*res)[0].CommonName)

	testSymptomByID(a, (*res)[0], t)
}

func testSymptomByID(a App, r SymptomRes, t *testing.T) {
	res, err := a.SymptomByID(r.ID)
	assert.NoError(t, err)
	assert.Equal(t, res.Name, r.Name)
	assert.Equal(t, res.CommonName, r.CommonName)
}
