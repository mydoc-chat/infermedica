package infermedica

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// DiagnosisReq is a struct to request diagnosis
type DiagnosisReq struct {
	Sex       Sex               `json:"sex"`
	Age       int               `json:"age"`
	Evidences []Evidence        `json:"evidence"`
	Extras    DiagnosisReqExras `json:"extras"`
}

type DiagnosisReqCovid19 struct {
	Sex       Sex               `json:"sex"`
	Age       int               `json:"age"`
	Evidences []EvidenceCovid19 `json:"evidence"`
	Extras    DiagnosisReqExras `json:"extras"`
}

// DiagnosisReqExras contains extra params for DiagnosisReq
type DiagnosisReqExras struct {
	DisableGroups bool `json:"disable_groups"`
}

// DiagnosisRes is a response struct for diagnosis
type DiagnosisRes struct {
	Question   Question                `json:"question"`
	Conditions []DiagnosisConditionRes `json:"conditions"`
	ShouldStop bool                    `json:"should_stop"`
	Extras     interface{}             `json:"extras"`
}

// QuestionType is a list of question types
type QuestionType string

const (
	// QuestionTypeSingle single question
	QuestionTypeSingle QuestionType = "single"
	// QuestionTypeGroupSingle question group
	QuestionTypeGroupSingle QuestionType = "group_single"
	// QuestionTypeGroupMultiple multiple question groups
	QuestionTypeGroupMultiple QuestionType = "group_multiple"
)

func (qt QuestionType) Ptr() *QuestionType { return &qt }
func (qt QuestionType) String() string     { return string(qt) }

func (qt *QuestionType) IsValid() bool {
	_, err := QuestionTypeFromString(qt.String())
	if err != nil {
		return false
	}
	return true
}

func QuestionTypeFromString(x string) (QuestionType, error) {
	switch strings.ToLower(x) {
	case "single":
		return QuestionTypeSingle, nil
	case "group_single":
		return QuestionTypeGroupSingle, nil
	case "group_multiple":
		return QuestionTypeGroupMultiple, nil
	default:
		return "", fmt.Errorf("Unexpected value for Question Type: %q", x)
	}
}

// Question struct
type Question struct {
	Type  QuestionType   `json:"type"`
	Text  string         `json:"text"`
	Items []QuestionItem `json:"items"`
}

// QuestionItem question item struct
type QuestionItem struct {
	ID      string               `json:"id"`
	Name    string               `json:"name"`
	Choices []QuestionItemChoice `json:"choices"`
}

// QuestionItemChoice question item choice struct
type QuestionItemChoice struct {
	ID    EvidenceChoiceID `json:"id"`
	Label string           `json:"label"`
}

// DiagnosisConditionRes is a response struct for condition + probability
type DiagnosisConditionRes struct {
	Condition
	Probability float64 `json:"probability"`
}

// Diagnosis is a func to request diagnosis for given data
func (a *App) Diagnosis(dr DiagnosisReq) (*DiagnosisRes, error) {
	if !dr.Sex.IsValid() {
		return nil, errors.New("Unexpected value for Sex")
	}
	req, err := a.prepareRequest("POST", "diagnosis", dr)
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

// Covid19 is a func to request diagnosis for covid-19 data
func (a *App) Covid19(dr DiagnosisReqCovid19) (*DiagnosisRes, error) {
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
	// bodyBytes, _ := ioutil.ReadAll(res.Body)
	// fmt.Printf("url: %#v\n", req.URL)
	// fmt.Printf("request: %#v\n", req)
	// fmt.Printf("response: %v", string(bodyBytes))
	// fmt.Printf("response: %#v", res.Body)
	r := DiagnosisRes{}
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("response: %#v\n", r)
	return &r, nil
}
