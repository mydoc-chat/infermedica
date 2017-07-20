package infermedica

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type DiagnosisReq struct {
	Sex      Sex        `json:"sex"`
	Age      int        `json:"age"`
	Evidence []Evidence `json:"evidence"`
}

type DiagnosisRes struct {
	Question   Question                `json:"question"`
	Conditions []DiagnosisConditionRes `json:"conditions"`
}

type Question struct {
	Type  string         `json:"type"`
	Text  string         `json:"text"`
	Items []QuestionItem `json:"items"`
}

type QuestionItem struct {
	ID      string               `json:"id"`
	Name    string               `json:"name"`
	Choices []QuestionItemChoice `json:"choices"`
}

type QuestionItemChoice struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

type DiagnosisConditionRes struct {
	Condition
	Probability float64 `json:"probability"`
}

func (a *app) Diagnosis(dr DiagnosisReq) (*DiagnosisRes, error) {
	if !dr.Sex.IsValid() {
		return nil, errors.New("Unexpected value for Sex")
	}
	req, err := a.prepareRequest("POST", "diagnosis", dr)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	r := DiagnosisRes{}
	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
