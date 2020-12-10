package infermedica

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// CovidDiagnosisReq describes covid interview diagnosis request
type CovidDiagnosisReq struct {
	Sex       Sex               `json:"sex"`
	Age       int               `json:"age"`
	Evidences []EvidenceCovid   `json:"evidence"`
	Extras    DiagnosisReqExras `json:"extras"`
}

// CovidTriageRes describes covid interview triage response
type CovidTriageRes struct {
	Description string           `json:"description"`
	Label       string           `json:"label"`
	Serious     []Serious        `json:"serious"`
	TriageLevel CovidTriageLevel `json:"triage_level"`
}

// CovidTriageLevel describes covid interview triage level
type CovidTriageLevel string

const (
	// CovidTriageLevelNoRisk stands for NoRisk
	CovidTriageLevelNoRisk CovidTriageLevel = "no_risk"
	// CovidTriageLevelSelfMonitoring stands for SelfMonitoring
	CovidTriageLevelSelfMonitoring CovidTriageLevel = "self_monitoring"
	// CovidTriageLevelQuarantine stands for Quarantine
	CovidTriageLevelQuarantine CovidTriageLevel = "quarantine"
	// CovidTriageLevelIsolationCall stands for IsolationCall
	CovidTriageLevelIsolationCall CovidTriageLevel = "isolation_call"
	// CovidTriageLevelCallDoctor stands for CallDoctor
	CovidTriageLevelCallDoctor CovidTriageLevel = "call_doctor"
	// CovidTriageLevelIsolationAmbulance stands for IsolationAmbulance
	CovidTriageLevelIsolationAmbulance CovidTriageLevel = "isolation_ambulance"
)

// CovidDiagnosis is a func to request diagnosis for covid-19 data
func (a *App) CovidDiagnosis(dr CovidDiagnosisReq) (*DiagnosisRes, error) {
	if !dr.Sex.IsValid() {
		return nil, errors.New("Unexpected value for Sex")
	}
	b, _ := json.Marshal(dr)
	fmt.Printf("body: %v\n", string(b))
	req, err := a.prepareRequest("POST", "covid19/diagnosis", dr)
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
	r := DiagnosisRes{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// CovidTriage is a func to request triage for given data
func (a *App) CovidTriage(dr CovidDiagnosisReq) (*CovidTriageRes, error) {
	if !dr.Sex.IsValid() {
		return nil, errors.New("Unexpected value for Sex")
	}
	req, err := a.prepareRequest("POST", "covid19/triage", dr)
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
	r := CovidTriageRes{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *App) CovidRiskFactors() (*[]RiskFactorRes, error) {
	req, err := a.prepareRequest("GET", "covid19/risk_factors", nil)
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

func (a *App) CovidSymptoms() (*[]SymptomRes, error) {
	req, err := a.prepareRequest("GET", "covid19/symptoms", nil)
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
