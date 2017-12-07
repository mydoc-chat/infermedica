package infermedica

import (
	"encoding/json"
	"net/http"
	"time"
)

type InfoRes struct {
	UpdatedAt        time.Time `json:"updated_at"`
	ConditionsCount  int       `json:"conditions_count"`
	SymptomsCount    int       `json:"symptoms_count"`
	RiskFactorsCount int       `json:"risk_factors_count"`
	LabTestsCount    int       `json:"lab_tests_count"`
}

func (a *App) Info() (*InfoRes, error) {
	req, err := a.prepareRequest("GET", "info", nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	r := InfoRes{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
