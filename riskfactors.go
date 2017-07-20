package infermedica

import (
	"encoding/json"
	"net/http"
)

type RiskFactorRes struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CommonName  string `json:"common_name"`
	SexFilter   string `json:"sex_filter"`
	Category    string `json:"category"`
	Seriousness string `json:"seriousness"`
	ImageUrl    string `json:"image_url"`
	ImageSource string `json:"image_source"`
}

func (a *app) RiskFactors() (*[]RiskFactorRes, error) {
	req, err := a.prepareRequest("GET", "risk_factors", nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	riskFactors := []RiskFactorRes{}
	err = json.NewDecoder(res.Body).Decode(riskFactors)
	if err != nil {
		return nil, err
	}
	return &riskFactors, nil
}

func (a *app) RiskFactorByID(id string) (*RiskFactorRes, error) {
	req, err := a.prepareRequest("GET", "risk_factors/"+id, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	r := RiskFactorRes{}
	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
