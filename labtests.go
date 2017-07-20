package infermedica

import (
	"encoding/json"
	"net/http"
)

type LabTestsRes struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	CommonName string      `json:"common_name"`
	Category   string      `json:"category"`
	Results    []LabResult `json:"results"`
}

type LabResult struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func (a *app) LabTests() (*[]LabTestsRes, error) {
	req, err := a.prepareRequest("GET", "lab_tests", nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	r := []LabTestsRes{}
	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *app) LabTestByID(id string) (*LabTestsRes, error) {
	req, err := a.prepareRequest("GET", "lab_tests/"+id, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	r := LabTestsRes{}
	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
