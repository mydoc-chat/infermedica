package infermedica

import (
	"encoding/json"
	"net/http"
)

type Condition struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CommonName string `json:"common_name"`
}

type ConditionRes struct {
	Condition
	SexFilter   string          `json:"sex_filter"`
	Categories  []string        `json:"categories"`
	Prevalence  string          `json:"prevalence"`
	Acuteness   string          `json:"acuteness"`
	Severity    string          `json:"severity"`
	Extras      ConditionExtras `json:"extras"`
	TriageLevel string          `json:"triage_level"`
}

type ConditionExtras struct {
	Hint      string `json:"hint"`
	ICD10Code string `json:"icd10_code"`
}

func (a *app) Conditions() (*[]ConditionRes, error) {
	req, err := a.prepareRequest("GET", "conditions", nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	r := []ConditionRes{}
	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *app) ConditionByID(id string) (*ConditionRes, error) {
	req, err := a.prepareRequest("GET", "conditions/"+id, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	r := ConditionRes{}
	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
