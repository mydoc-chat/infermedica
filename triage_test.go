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
				ChoiceID: "present",
			},
			Evidence{
				ID:       "s_488",
				ChoiceID: "present",
			},
			Evidence{
				ID:       "s_418",
				ChoiceID: "present",
			},
			Evidence{
				ID:       "s_320",
				ChoiceID: "absent",
			},
		},
	}
	res, err := a.Triage(r)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Serious)
	assert.NotEmpty(t, res.TriageLevel)
}
