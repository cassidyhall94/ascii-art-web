package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAsciiArt(t *testing.T) {
	testCases := []struct {
		banner           string
		input            string
		expectedResponse int
	}{
		{
			// when valid banner and input is passed
			banner:           "shadow.txt",
			input:            "this is history",
			expectedResponse: http.StatusOK,
		},
		{
			// when an invalid input file is passed
			banner:           "shadow.txt",
			input:            "",
			expectedResponse: http.StatusBadRequest,
		},
		{
			// when an invalid banner file
			banner:           "shadwo.txt",
			input:            "this is history",
			expectedResponse: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		form := url.Values{}
		form.Add("Banner", tc.banner)
		form.Add("input", tc.input)

		request := httptest.NewRequest("POST", "/ascii-art", strings.NewReader(form.Encode()))
		request.PostForm = form
		responseRecorder := httptest.NewRecorder()

		process(responseRecorder, request)
		if responseRecorder.Code != tc.expectedResponse {
			t.Errorf("Want status '%d', got '%d'", tc.expectedResponse, responseRecorder.Code)
		}
		// assert.Equal(t, responseRecorder.Code, tc.expectedResponse)
	}
}
