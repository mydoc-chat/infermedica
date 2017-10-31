package infermedica

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp_LabTests(t *testing.T) {
	a := GetTestAppInstance()
	res, err := a.LabTests()
	assert.NoError(t, err)
	assert.NotEqual(t, len(*res), 0)
	assert.NotEmpty(t, (*res)[0].ID)
	assert.NotEmpty(t, (*res)[0].Name)
	assert.NotEmpty(t, (*res)[0].CommonName)

	testLabTestByID(a, (*res)[0], t)
}

func testLabTestByID(a App, r LabTestsRes, t *testing.T) {
	res, err := a.LabTestByID(r.ID)
	assert.NoError(t, err)
	assert.Equal(t, res.Name, r.Name)
	assert.Equal(t, res.CommonName, r.CommonName)
}

func TestApp_LabTestsRecommend(t *testing.T) {
	a := GetTestAppInstance()
	r := DiagnosisReq{
		Sex: SexMale,
		Age: 30,
		Evidences: []Evidence{
			Evidence{
				ID:       "s_1193",
				ChoiceID: EvidenceChoiceIDPresent,
			},
			Evidence{
				ID:       "s_488",
				ChoiceID: EvidenceChoiceIDPresent,
			},
			Evidence{
				ID:       "s_418",
				ChoiceID: EvidenceChoiceIDPresent,
			},
		},
	}
	res, err := a.LabTestsRecommend(r)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Recommended)
	assert.NotEmpty(t, res.Obligatory)
}
