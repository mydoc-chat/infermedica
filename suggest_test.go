package infermedica

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp_Suggest(t *testing.T) {
	a := GetTestAppInstance()
	r := SuggestReq{
		Sex: SexMale,
		Age: 30,
		Evidences: []Evidence{
			Evidence{
				ID:       "s_1193",
				ChoiceID: EvidenceChoiceIDPresent,
				Initial:  true,
			},
			Evidence{
				ID:       "s_488",
				ChoiceID: EvidenceChoiceIDPresent,
				Initial:  true,
			},
		},
	}
	res, err := a.Suggest(r)
	assert.NoError(t, err)
	assert.NotEqual(t, len(*res), 0)
	assert.NotEmpty(t, (*res)[0].ID)
	assert.NotEmpty(t, (*res)[0].Name)
	assert.NotEmpty(t, (*res)[0].CommonName)
}
