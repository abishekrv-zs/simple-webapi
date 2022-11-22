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
		req         map[string]any
		expCode     int
		expResp     string
	}{
		{
			description: "Normal case to add an emp",
			req:         map[string]any{"id": 3, "name": "Sujith", "age": 22, "address": "Bangalore"},
			expCode:     201,
			expResp:     `{"id":3,"name":"Sujith","age":22,"address":"Bangalore"}`,
		},
		{
			description: "Case with missing fields(expect id)",
			req:         map[string]any{"id": 999, "address": "Kochi"},
			expCode:     201,
			expResp:     `{"id":999,"address":"Kochi"}`,
		},
		{
			description: "Case with conflicting id (id already exists)",
			req:         map[string]any{"id": 1, "name": "hulk", "age": 100},
			expCode:     409,
			expResp:     `{"error": "id already exists"}`,
		},
		{
			description: "Case with empty id",
			req:         map[string]any{"name": "Anonymous", "age": 22},
			expCode:     400,
			expResp:     `{"error": "id cannot be 0"}`,
		},
		{
			description: "Request body with wrong types",
			req:         map[string]any{"id": "abishek", "name": 26},
			expCode:     400,
			expResp:     `{"error": "invalid request body"}`,
		},
	}

	for _, tc := range tests {

		reqBody, _ := json.Marshal(tc.req)
		mockReq := httptest.NewRequest("POST", "/employee", bytes.NewReader(reqBody))
		mockResp := httptest.NewRecorder()

		PostEmployee(mockResp, mockReq)

		assert.Equal(t, tc.expCode, mockResp.Code, tc.description)
		assert.Equal(t, tc.expResp, mockResp.Body.String(), tc.description)
	}
}
