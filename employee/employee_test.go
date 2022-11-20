package employee

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetEmployee(t *testing.T) {

	tests := []struct {
		description string
		req         string
		expCode     int
		expResp     string
	}{
		{
			"Normal get request",
			"",
			200,
			`[{"id":1,"name":"Abishek","age":22,"address":"Chennai"},{"id":2,"name":"Kavin","age":22,"address":"Mumbai"}]`,
		},
		{
			"Request with some random body",
			"sadsahdksadk",
			200,
			`[{"id":1,"name":"Abishek","age":22,"address":"Chennai"},{"id":2,"name":"Kavin","age":22,"address":"Mumbai"}]`,
		},
	}

	for _, tc := range tests {
		mockReq := httptest.NewRequest("GET", "/employee", strings.NewReader(tc.req))
		mockResp := httptest.NewRecorder()

		GetEmployee(mockResp, mockReq)

		assert.Equal(t, tc.expCode, mockResp.Code)
		assert.Equal(t, tc.expResp, strings.TrimSpace(mockResp.Body.String()))
	}
}

func TestPostEmployee(t *testing.T) {
	tests := []struct {
		description string
		req         string
		expResp     string
		expCode     int
	}{
		{
			"Normal case to add an emp",
			`{"id":3,"name":"Sujith","age":22,"address":"Bangalore"}`,
			`{"id":3,"name":"Sujith","age":22,"address":"Bangalore"}`,
			201,
		},
		{
			"Missing fields in request body",
			`{"id":198,"name":"Anonymous"}`,
			`{"id":198,"name":"Anonymous","age":0,"address":""}`,
			201,
		},
		{
			"Blank request body",
			"",
			"Conflict in parsing request body",
			409,
		},
	}

	for _, tc := range tests {
		mockReq := httptest.NewRequest("POST", "/employee", strings.NewReader(tc.req))
		mockResp := httptest.NewRecorder()

		PostEmployee(mockResp, mockReq)

		assert.Equal(t, tc.expCode, mockResp.Code)
		assert.Equal(t, tc.expResp, strings.TrimSpace(mockResp.Body.String()))
	}
}
