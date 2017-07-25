package infermedica

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

type App struct {
	baseURL string
	appID   string
	appKey  string
}

func NewApp(id, key string) App {
	return App{
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

func (a App) prepareRequest(method, url string, body interface{}) (*http.Request, error) {

	switch method {
	case "GET":
		return a.prepareGETRequest(url)
	case "POST":
		return a.preparePOSTRequest(url, body)
	}
	return nil, errors.New("Method not allowed")
}

func (a App) addHeaders(req *http.Request) {
	req.Header.Add("App-Id", a.appID)
	req.Header.Add("App-Key", a.appKey)
	req.Header.Add("Content-Type", "application/json")
}

func (a App) prepareGETRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", a.baseURL+url, nil)
	if err != nil {
		return nil, err
	}
	a.addHeaders(req)
	return req, nil
}

func (a App) preparePOSTRequest(url string, body interface{}) (*http.Request, error) {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", a.baseURL+url, b)
	if err != nil {
		return nil, err
	}
	a.addHeaders(req)
	return req, nil
}
