package infermedica

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp_Explain(t *testing.T) {
	a := GetTestAppInstance()
	r := ExplainReq{
		Sex:    SexMale,
		Age:    30,
		Target: "c_49",
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
		},
	}
	resp, err := a.Explain(r)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.SupportingEvidence)
}
