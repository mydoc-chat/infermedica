package infermedica

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp_Triage(t *testing.T) {
	a := GetTestAppInstance()
	r := TriageReq{
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
			Evidence{
				ID:       "s_320",
				ChoiceID: EvidenceChoiceIDAbsent,
			},
		},
	}
	res, err := a.Triage(r)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Serious)
	assert.NotEmpty(t, res.TriageLevel)
}
