package controller

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type mockRequestSender struct {
	method   string
	url      string
	response *http.Response
	wantErr  bool
}

func (mrs *mockRequestSender) SendRequest(method, url, values map[string]interface{}) (*http.Response, error) {
	if mrs.wantErr {
		return nil, fmt.Errorf("send request error")
	}
	return mrs.response, nil
}
func Test_GetCard(t *testing.T) {

	JsonCardResponse := `{
			"cards": [
				{
					"value": "9",
					"suit": "HEARTS"
				},
			],
		}`

	responseValue := 12

	validURL := "https://testApi.com/api/deck/new/draw/?count=1"

	type globals struct {
		requestSender func(method, url, values map[string]interface{}) (*http.Response, error)
	}

	tests := []struct {
		name        string
		globals     globals
		wantedValue int
		wantErr     bool
	}{
		{
			name: "Valid test",
			globals: globals{
				requestSender: (&mockRequestSender{
					method: http.MethodGet,
					url:    validURL,
					response: &http.Response{
						StatusCode:    http.StatusOK,
						Body:          ioutil.NopCloser(bytes.NewBufferString(JsonCardResponse)),
						ContentLength: int64(len(JsonCardResponse)),
					},
				}).SendRequest,
			},
			wantedValue: responseValue,
			wantErr:     false,
		},
		{
			name: "InValid test",
			globals: globals{
				requestSender: (&mockRequestSender{
					wantErr: true,
				}).SendRequest,
			},
			wantedValue: 0,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, _, err := GetCard()
			if (err != nil) != tt.wantErr {
				t.Errorf("PlayCards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantedValue {
				t.Errorf("PlayCards() Got Value = %v, wanted Value %v", gotValue, tt.wantedValue)
			}
		})
	}
}
