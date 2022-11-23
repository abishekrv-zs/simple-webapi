package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEmployeeHandler(t *testing.T) {
	tests := []struct {
		description    string
		method         string
		reqBody        []byte
		reqQueryParams map[string]string
		expCode        int
		expResp        string
	}{
		{
			description:    "Case get all employee `/emp`",
			method:         "GET",
			reqQueryParams: nil,
			reqBody:        nil,
			expCode:        200,
			expResp:        `[{"id":"1","name":"Abishek","phoneNumber":"1234567890","department":{"id":"1","name":"Software"}},{"id":"2","name":"Kavin","phoneNumber":"1234567891","department":{"id":"1","name":"Software"}},{"id":"3","name":"Kiren","phoneNumber":"1234567892","department":{"id":"2","name":"Finance"}},{"id":"4","name":"Sujith","phoneNumber":"1234567893","department":{"id":"3","name":"Admin"}}]`,
		},
		{
			description:    "Case get employee by id `/emp?id=1`",
			method:         "GET",
			reqQueryParams: map[string]string{"id": "1"},
			reqBody:        nil,
			expCode:        200,
			expResp:        `{"id":"1","name":"Abishek","phoneNumber":"1234567890","department":{"id":"1","name":"Software"}}`,
		},
		{
			description:    "case get emp with invalid id",
			method:         "GET",
			reqQueryParams: map[string]string{"id": "0"},
			reqBody:        nil,
			expCode:        404,
			expResp:        `{"error":"id not found"}`,
		},
		{
			description:    "case to add an emp",
			method:         "POST",
			reqQueryParams: nil,
			reqBody:        []byte(`{"id":"5","name":"newGuy","phoneNumber":"987643210","department":{"id":"4","name":"HR"}}`),
			expCode:        201,
			expResp:        `{"id":"5","name":"newGuy","phoneNumber":"987643210","department":{"id":"4","name":"HR"}}`,
		},
		{
			description:    "case with mismatching fields",
			method:         "POST",
			reqQueryParams: nil,
			reqBody:        []byte(`{"id":"5","name":2,"phoneNumber":0,"department":{"name":"HR"}}`),
			expCode:        400,
			expResp:        `{"error":"invalid request body"}`,
		},
		{
			description:    "case with PUT method",
			method:         "PUT",
			reqQueryParams: nil,
			reqBody:        nil,
			expCode:        405,
			expResp:        `{"error":"invalid request method"}`,
		},
	}

	for _, tc := range tests {
		mockReq, _ := http.NewRequest(tc.method, "/emp", bytes.NewReader(tc.reqBody))
		mockResp := httptest.NewRecorder()

		// Mocking request with query params
		if tc.reqQueryParams != nil {
			queryParams := mockReq.URL.Query()
			for key, value := range tc.reqQueryParams {
				queryParams.Set(key, value)
			}
			mockReq.URL.RawQuery = queryParams.Encode()
		}

		employeeHandler(mockResp, mockReq)

		assert.Equal(t, tc.expCode, mockResp.Code, tc.description)
		assert.Equal(t, tc.expResp, mockResp.Body.String())
	}
}
