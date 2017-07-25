package infermedica

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp_Info(t *testing.T) {
	a := GetTestAppInstance()
	res, err := a.Info()
	assert.NoError(t, err)
	assert.NotEmpty(t, res.UpdatedAt)
	assert.NotEmpty(t, res.ConditionsCount)
	assert.NotEmpty(t, res.LabTestsCount)
	assert.NotEmpty(t, res.RiskFactorsCount)
	assert.NotEmpty(t, res.SymptomsCount)
}
