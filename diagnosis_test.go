package infermedica

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp_Diagnosis(t *testing.T) {
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
		},
	}
	res, err := a.Diagnosis(r)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Question)
	assert.NotEmpty(t, res.Conditions)
}
