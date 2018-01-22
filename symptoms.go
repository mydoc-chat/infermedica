package infermedica

import (
	"encoding/json"
	"net/http"
	"time"
)

type SymptomRes struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	CommonName     string         `json:"common_name"`
	Category       string         `json:"category"`
	Seriousness    string         `json:"seriousness"`
	Children       []SymptomChild `json:"children"`
	ImageURL       string         `json:"image_url"`
	ImageSource    string         `json:"image_source"`
	ParentID       string         `json:"parent_id"`
	ParentRelation string         `json:"parent_relation"`
}

type SymptomChild struct {
	ID             string `json:"id"`
	ParentRelation string `json:"parent_relation"`
}

func (a *App) Symptoms() (*[]SymptomRes, error) {
	req, err := a.prepareRequest("GET", "symptoms", nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	r := []SymptomRes{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *App) SymptomsIDMap() (*map[string]SymptomRes, error) {
	r, err := a.Symptoms()
	if err != nil {
		return nil, err
	}
	rmap := make(map[string]SymptomRes)
	for _, sr := range *r {
		rmap[sr.ID] = sr
	}
	return &rmap, nil
}

func (a *App) SymptomByID(id string) (*SymptomRes, error) {
	req, err := a.prepareRequest("GET", "symptoms/"+id, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	r := SymptomRes{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
