package infermedica

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// SuggestReq is a struct to request suggestions
type SuggestReq struct {
	Sex       Sex        `json:"sex"`
	Age       int        `json:"age"`
	Evidences []Evidence `json:"evidence"`
}

// SuggestRes is a response struct for suggest
type SuggestRes struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CommonName string `json:"common_name"`
}

// Suggest is a func to request suggestions
func (a *App) Suggest(sr SuggestReq) (*[]SuggestRes, error) {
	if !sr.Sex.IsValid() {
		return nil, errors.New("Unexpected value for Sex")
	}
	req, err := a.prepareRequest("POST", "suggest", sr)
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
	r := []SuggestRes{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
