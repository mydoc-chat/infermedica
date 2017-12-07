package infermedica

import (
	"encoding/json"
	"net/http"
	"time"
)

type RiskFactorRes struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	CommonName  string    `json:"common_name"`
	SexFilter   SexFilter `json:"sex_filter"`
	Category    string    `json:"category"`
	Seriousness string    `json:"seriousness"`
	ImageURL    string    `json:"image_url"`
	ImageSource string    `json:"image_source"`
}

func (a *App) RiskFactors() (*[]RiskFactorRes, error) {
	req, err := a.prepareRequest("GET", "risk_factors", nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	r := []RiskFactorRes{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *App) RiskFactorByID(id string) (*RiskFactorRes, error) {
	req, err := a.prepareRequest("GET", "risk_factors/"+id, nil)
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
	r := RiskFactorRes{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
