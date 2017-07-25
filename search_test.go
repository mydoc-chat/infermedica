package infermedica

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp_Search(t *testing.T) {
	a := GetTestAppInstance()
	res, err := a.Search("Smoking", SexFemale, 8, SearchTypeRiskFactor)
	assert.NoError(t, err)
	assert.NotEqual(t, len(*res), 0)
	assert.NotEmpty(t, (*res)[0].ID)
	assert.NotEmpty(t, (*res)[0].Label)
}
