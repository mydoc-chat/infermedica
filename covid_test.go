package infermedica

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp_CovidDiagnosis(t *testing.T) {
	a := GetTestAppInstance()
	r := CovidDiagnosisReq{
		Sex:       SexMale,
		Age:       30,
		Evidences: []EvidenceCovid{},
	}
	res, err := a.CovidDiagnosis(r)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Question)
}

func TestApp_CovidTriage(t *testing.T) {
	a := GetTestAppInstance()
	r := CovidDiagnosisReq{
		Sex: SexMale,
		Age: 30,
		Evidences: []EvidenceCovid{
			EvidenceCovid{
				ID:       "s_22",
				ChoiceID: EvidenceChoiceIDPresent,
			},
			EvidenceCovid{
				ID:       "s_14",
				ChoiceID: EvidenceChoiceIDPresent,
			},
			EvidenceCovid{
				ID:       "s_2",
				ChoiceID: EvidenceChoiceIDPresent,
			},
			EvidenceCovid{
				ID:       "p_27",
				ChoiceID: EvidenceChoiceIDAbsent,
			},
		},
	}
	res, err := a.CovidTriage(r)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Serious)
	assert.NotEmpty(t, res.TriageLevel)
}
