package infermedica

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type app struct {
	baseURL string
	appID   string
	appKey  string
}

func NewApp(id, key string) app {
	return app{
		baseURL: "https://api.infermedica.com/v2/",
		appID:   id,
		appKey:  key,
	}
}

type Sex string

const (
	SexMale   Sex = "male"
	SexFemale Sex = "female"
)

func (s Sex) Ptr() *Sex      { return &s }
func (s Sex) String() string { return string(s) }

func (s *Sex) IsValid() bool {
	_, err := SexFromString(s.String())
	if err != nil {
		return false
	}
	return true
}

func SexFromString(x string) (Sex, error) {
	switch strings.ToLower(x) {
	case "male":
		return SexMale, nil
	case "female":
		return SexFemale, nil
	default:
		return "", fmt.Errorf("Unexpected value for Sex: %q", x)
	}
}

type Evidence struct {
	ID       string `json:"id"`
	ChoiceID string `json:"choice_id"`
}

func (a app) prepareRequest(method, url string, body interface{}) (*http.Request, error) {
	var b *bytes.Buffer
	if body != nil {
		json.NewEncoder(b).Encode(body)
	}
	req, err := http.NewRequest(method, a.baseURL+url, b)
	if err != nil {
		return nil, err
	}
	req.Header.Add("App-Id", a.appID)
	req.Header.Add("App-Key", a.appKey)
	req.Header.Add("Content-Type", "application/json")
	return req, nil

}
