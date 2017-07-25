package infermedica

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp_Parse(t *testing.T) {
	a := GetTestAppInstance()
	req := ParseReq{
		Text: "I feel smoach pain but no couoghing today",
	}
	res, err := a.Parse(req)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Mentions)
	assert.NotEqual(t, len(res.Mentions), 0)
}
