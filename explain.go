package infermedica

import (
	"encoding/json"
	"net/http"
)

type ExplainReq struct {
	Sex      Sex        `json:"sex"`
	Age      int        `json:"age"`
	Target   string     `json:"target"`
	Evidence []Evidence `json:"evidence"`
}

type ExplainRes struct {
	SupportingEvidence  []EvidenceItem `json:"supporting_evidence"`
	ConflictingEvidence []EvidenceItem `json:"conflicting_evidence"`
}

type EvidenceItem struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CommonName string `json:"common_name"`
}

func (a *app) Explain(er ExplainReq) (*ExplainRes, error) {
	req, err := a.prepareRequest("POST", "explain", er)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	r := ExplainRes{}
	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
