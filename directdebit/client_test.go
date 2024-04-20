package directdebit_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/praswicaksono/ayoconnect-direct-debit-go/directdebit"
)

func TestSetHeaders(t *testing.T) {
	client := directdebit.Client{}
	req := &http.Request{
		Header: make(http.Header),
	}
	headers := directdebit.RequestHeader{
		Timestamp:             "timestamp",
		Signature:             "signature",
		ClientKey:             "clientKey",
		Authorization:         "auth",
		AuthorizationCustomer: "clientAuth",
		ExternalID:            "externalID",
	}

	resultReq := client.SetHeaders(req, headers)

	tests := map[string]struct {
		headerName  string
		expectedVal string
	}{
		"Content-Type":           {"Content-Type", "application/json"},
		"Accept":                 {"Accept", "application/json"},
		"X-TIMESTAMP":            {"X-TIMESTAMP", "timestamp"},
		"X-SIGNATURE":            {"X-SIGNATURE", "signature"},
		"X-CLIENT-KEY":           {"X-CLIENT-KEY", "clientKey"},
		"Authorization":          {"Authorization", "auth"},
		"Authorization-Customer": {"Authorization-Customer", "clientAuth"},
		"X-EXTERNAL-ID":          {"X-EXTERNAL-ID", "externalID"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := resultReq.Header.Get(tc.headerName); got != tc.expectedVal {
				t.Errorf("Expected header %s to have value %s, but got %s", tc.headerName, tc.expectedVal, got)
			}
		})
	}
}

func TestExecute(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("response body"))
	}))
	defer ts.Close()

	client := directdebit.Client{
		Config: &directdebit.Config{
			EndpointBaseURL: ts.URL,
			HTTPClient:      &http.Client{},
		},
	}

	method := "GET"
	path := "/testpath"
	headers := directdebit.RequestHeader{}
	jsonBytes := []byte("test payload")

	res, err := client.Execute(context.Background(), method, path, headers, jsonBytes)
	if err != nil {
		t.Fatalf("Did not expect an error, but got: %v", err)
	}

	expectedResponse := []byte("response body")
	if !bytes.Equal(res, expectedResponse) {
		t.Errorf("Expected response to be %s, but got %s", expectedResponse, res)
	}
}

func TestBuildHeader(t *testing.T) {
	client := directdebit.Client{
		Config: &directdebit.Config{
			ClientID:   "testClientID",
			MerchantID: "testMerchantID",
			ChannelID:  "testChannelID",
		},
	}

	timestamp := "testTimestamp"
	signature := "testSignature"
	b2bToken := "testB2BToken"
	b2b2cToken := "testB2B2CToken"
	externalID := "testExternalID"

	header := client.BuildHeader(timestamp, signature, b2bToken, b2b2cToken, externalID)

	if header.Timestamp != timestamp {
		t.Errorf("Expected Timestamp to be %s, but got %s", timestamp, header.Timestamp)
	}

	if header.Signature != signature {
		t.Errorf("Expected Signature to be %s, but got %s", signature, header.Signature)
	}

	if header.Authorization != "Bearer "+b2bToken {
		t.Errorf("Expected Authorization to be %s, but got %s", b2bToken, header.Authorization)
	}

	if header.AuthorizationCustomer != "Bearer "+b2b2cToken {
		t.Errorf("Expected ClientAuthorization to be %s, but got %s", b2b2cToken, header.AuthorizationCustomer)
	}

	if header.ExternalID != externalID {
		t.Errorf("Expected ExternalID to be %s, but got %s", externalID, header.ExternalID)
	}
}

func TestExecuteErrorScenarios(t *testing.T) {
	// Error in http.NewRequest
	client := directdebit.Client{
		Config: &directdebit.Config{
			EndpointBaseURL: "http://example.com",
			HTTPClient:      &http.Client{},
		},
	}

	_, err := client.Execute(context.Background(), "\n", "testpath", directdebit.RequestHeader{}, []byte("test payload"))
	if err == nil {
		t.Errorf("Expected an error due to invalid method, but got none")
	}

	// Error in c.config.HTTPClient.Do(req)
	client.Config.EndpointBaseURL = "http://invalid-url:1234" // Using an invalid URL to simulate failure

	_, err = client.Execute(context.Background(), "GET", "testpath", directdebit.RequestHeader{}, []byte("test payload"))
	if err == nil {
		t.Errorf("Expected an error due to connection problem, but got none")
	}

	// Error in io.ReadAll(res.Body)
	ts2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "1")
	}))
	defer ts2.Close()

	client.Config.EndpointBaseURL = ts2.URL

	_, err = client.Execute(context.Background(), "GET", "testpath", directdebit.RequestHeader{}, []byte("test payload"))
	if err == nil {
		t.Errorf("Expected an error while reading the response body, but got none")
	}
}
