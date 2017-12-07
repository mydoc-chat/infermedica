package infermedica

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type LookupRes struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

func (a *App) Lookup(phrase string, sex Sex) (*LookupRes, error) {
	if !sex.IsValid() {
		return nil, errors.New("Unexpected value for Sex")
	}
	url := "lookup?phrase=" + phrase + "&sex=" + sex.String()
	req, err := a.prepareRequest("GET", url, nil)
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
	r := LookupRes{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
