package infermedica

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp_RiskFactors(t *testing.T) {
	a := GetTestAppInstance()
	res, err := a.RiskFactors()
	assert.NoError(t, err)
	assert.NotEqual(t, len(*res), 0)
	assert.NotEmpty(t, (*res)[0].ID)
	assert.NotEmpty(t, (*res)[0].Name)
	assert.NotEmpty(t, (*res)[0].CommonName)

	testRiskFactorByID(a, (*res)[0], t)
}

func testRiskFactorByID(a App, r RiskFactorRes, t *testing.T) {
	res, err := a.RiskFactorByID(r.ID)
	assert.NoError(t, err)
	assert.Equal(t, res.Name, r.Name)
	assert.Equal(t, res.CommonName, r.CommonName)
}
