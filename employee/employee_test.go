package employee

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestGetEmployee(t *testing.T) {

	tests := []struct {
		description string
		expCode     int
		expResp     []employee
	}{
		{
			"Normal get request",
			200,
			[]employee{
				{1, "Abishek", 22, "Chennai"},
				{2, "Kavin", 22, "Mumbai"},
			},
		},
	}

	for _, tc := range tests {
		mockReq := httptest.NewRequest("GET", "/employee", nil)
		mockResp := httptest.NewRecorder()

		GetEmployee(mockResp, mockReq)

		var actResp []employee

		_ = json.Unmarshal(mockResp.Body.Bytes(), &actResp)

		assert.Equal(t, tc.expCode, mockResp.Code, tc.description)
		assert.Equal(t, tc.expResp, actResp, tc.description)

	}
}

func TestPostEmployee(t *testing.T) {
	tests := []struct {
		description string
		req         employee
		expCode     int
		expResp     employee
	}{
		{
			"Normal case to add an emp",
			employee{3, "Sujith", 22, "Bangalore"},
			201,
			employee{3, "Sujith", 22, "Bangalore"},
		},
	}

	for _, tc := range tests {

		reqBody, _ := json.Marshal(tc.req)
		mockReq := httptest.NewRequest("POST", "/employee", bytes.NewReader(reqBody))
		mockResp := httptest.NewRecorder()

		PostEmployee(mockResp, mockReq)

		var actResp employee
		_ = json.Unmarshal(mockResp.Body.Bytes(), &actResp)

		assert.Equal(t, tc.expCode, mockResp.Code, tc.description)
		assert.Equal(t, tc.expResp, actResp, tc.description)
	}
}
