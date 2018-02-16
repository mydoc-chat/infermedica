package infermedica

import (
	"encoding/json"
	"net/http"
	"time"
)

type ParseReq struct {
	Text string `json:"text"`
}

type ParseRes struct {
	Mentions []Mention `json:"mentions"`
}

type Mention struct {
	Orth       string `json:"orth"`
	Name       string `json:"name"`
	ID         string `json:"id"`
	ChoiceID   string `json:"choice_id"`
	Type       string `json:"type"`
	CommonName string `json:"common_name"`
}

func (a *App) Parse(pr ParseReq) (*ParseRes, error) {
	model := a.model
	a.model = ""
	req, err := a.preparePOSTRequest("parse", pr)
	if err != nil {
		return nil, err
	}
	a.model = model
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	r := ParseRes{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
