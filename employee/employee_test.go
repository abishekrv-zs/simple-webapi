package employee

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGetAllEmployee(t *testing.T) {

	tests := []struct {
		description string
		expCode     int
		expResp     []employee
	}{
		{
			"Normal get all request",
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

		GetAllEmployee(mockResp, mockReq)

		var actResp []employee

		if err := json.Unmarshal(mockResp.Body.Bytes(), &actResp); err != nil {
			log.Println(err)
		}

		assert.Equal(t, tc.expCode, mockResp.Code, tc.description)
		assert.Equal(t, tc.expResp, actResp, tc.description)
	}
}

func TestGetEmployee(t *testing.T) {
	tests := []struct {
		description string
		empId       int
		expCode     int
		expResp     string
	}{
		{
			description: "case for valid emp id",
			empId:       1,
			expCode:     200,
			expResp:     `{"id":1,"name":"Abishek","age":22,"address":"Chennai"}`,
		},
		{
			description: "case for invalid emp id",
			empId:       -1,
			expCode:     400,
			expResp:     `{"error": "invalid employee id"}`,
		},
		{
			description: "case for non existent emp id",
			empId:       999,
			expCode:     404,
			expResp:     `{"error": "employee id not found"}`,
		},
	}

	for _, tc := range tests {

		mockReq := httptest.NewRequest("GET", "/employee", nil)
		mockReq = mux.SetURLVars(mockReq, map[string]string{"id": strconv.Itoa(tc.empId)})
		mockResp := httptest.NewRecorder()

		GetEmployee(mockResp, mockReq)

		assert.Equal(t, tc.expCode, mockResp.Code, tc.description)
		assert.Equal(t, tc.expResp, mockResp.Body.String(), tc.description)
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
			description: "Case with invalid id",
			req:         map[string]any{"id": -5, "name": "Anonymous", "age": 22},
			expCode:     400,
			expResp:     `{"error": "invalid id"}`,
		},
		{
			description: "Request body with wrong types",
			req:         map[string]any{"id": "abishek", "name": 26},
			expCode:     400,
			expResp:     `{"error": "invalid request body"}`,
		},
	}

	for _, tc := range tests {

		reqBody, err := json.Marshal(tc.req)
		if err != nil {
			log.Println(err)
		}
		mockReq := httptest.NewRequest("POST", "/employee", bytes.NewReader(reqBody))
		mockResp := httptest.NewRecorder()

		PostEmployee(mockResp, mockReq)

		assert.Equal(t, tc.expCode, mockResp.Code, tc.description)
		assert.Equal(t, tc.expResp, mockResp.Body.String(), tc.description)
	}
}
