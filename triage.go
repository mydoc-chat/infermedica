package infermedica

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type TriageReq struct {
	Sex       Sex        `json:"sex"`
	Age       int        `json:"age"`
	Evidences []Evidence `json:"evidence"`
}

type TriageRes struct {
	TriageLevel TriageLevel `json:"triage_level"`
	Serious     []Serious   `json:"serious"`
}

type Serious struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CommonName  string `json:"common_name"`
	IsEmergency bool   `json:"is_emergency"`
}

type TriageLevel string

const (
	TriageLevelEmergency    TriageLevel = "emergency"
	TriageLevelConsultation TriageLevel = "consultation"
	TriageLevelSelfCare     TriageLevel = "self_care"
)

func (s TriageLevel) Ptr() *TriageLevel { return &s }
func (s TriageLevel) String() string    { return string(s) }

func (s *TriageLevel) IsValid() bool {
	_, err := TriageLevelFromString(s.String())
	if err != nil {
		return false
	}
	return true
}

func TriageLevelFromString(x string) (TriageLevel, error) {
	switch strings.ToLower(x) {
	case "emergency":
		return TriageLevelEmergency, nil
	case "consultation":
		return TriageLevelConsultation, nil
	case "self_care":
		return TriageLevelSelfCare, nil
	default:
		return "", fmt.Errorf("Unexpected value for triage level: %q", x)
	}
}

func (a *App) Triage(tr TriageReq) (*TriageRes, error) {
	if !tr.Sex.IsValid() {
		return nil, errors.New("Unexpected value for Sex")
	}
	req, err := a.prepareRequest("POST", "triage", tr)
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
	r := TriageRes{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
