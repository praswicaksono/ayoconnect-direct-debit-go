package directdebit

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"log/slog"
)

// #nosec
const (
	GetBusinessAccessTokenEndpoint = "/api/v1.0/access-token/b2b"
	GetAuthCodeEndpoint            = "/api/v1.0/get-auth-code"
	AccountBindingEndpoint         = "/api/v1.0/registration-account-binding"
	GetCustomerAccessTokenEndpoint = "/api/v1.0/access-token/b2b2c"
	DebitEndpoint                  = "/api/v1.0/debit/payment-host-to-host"
	UnbindEndpoint                 = "/api/v1.0/registration-account-unbinding"
	DebitStatusEndpoint            = "/api/v1.0/debit/status"
)

func New(c *Config) (*Client, error) {
	// @TODO: clone and override http client timeout here
	return &Client{Config: c}, nil
}

func (c Client) SetHeaders(req *http.Request, headers RequestHeader) *http.Request {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-TIMESTAMP", headers.Timestamp)
	req.Header.Set("X-SIGNATURE", headers.Signature)
	req.Header.Set("X-CLIENT-KEY", headers.ClientKey)
	req.Header.Set("X-PARTNER-ID", headers.PartnerID)

	if headers.ChannelID != "" {
		req.Header.Set("CHANNEL-ID", headers.ChannelID)
	}

	if headers.Authorization != "" {
		req.Header.Set("Authorization", headers.Authorization)
	}

	if headers.AuthorizationCustomer != "" {
		req.Header.Set("Authorization-Customer", headers.AuthorizationCustomer)
	}

	if headers.ExternalID != "" {
		req.Header.Set("X-EXTERNAL-ID", headers.ExternalID)
	}

	return req
}

// TODO: Execute should receive context and use the context id for logging,
//
//	also use http.NewRequestWithContext instead of NewRequest
func (c Client) Execute(ctx context.Context, method string, path string, headers RequestHeader, jsonBytes []byte) ([]byte, error) {
	req, err := http.NewRequest(method, c.Config.EndpointBaseURL+path, bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req = c.SetHeaders(req, headers)

	res, err := c.Config.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if !(res.StatusCode == http.StatusOK || res.StatusCode == http.StatusAccepted) {
		c.Config.Logger.ErrorContext(ctx, "failed to execute request",
			slog.String("method", method),
			slog.String("path", path),
			slog.String("request_body", string(jsonBytes)),
			slog.String("response_status", res.Status),
			slog.String("response_body", string(resBody)),
		)
		errResp := ResponseError{}
		err := json.Unmarshal(resBody, &errResp)
		if err != nil {
			return nil, err
		}
		return nil, &errResp
	}

	return resBody, nil
}

func (c Client) BuildHeader(timestamp string, signature string, b2bToken string, b2b2cToken string, externalID string) RequestHeader {
	headers := RequestHeader{
		Timestamp: timestamp,
		ClientKey: c.Config.ClientID,
		Signature: signature,
		PartnerID: c.Config.MerchantID,
		ChannelID: c.Config.ChannelID,
	}

	if externalID != "" {
		headers.ExternalID = externalID
	}

	if b2bToken != "" {
		// do request to get b2b auth token
		headers.Authorization = "Bearer " + b2bToken
	}

	if b2b2cToken != "" {
		// # do request to get b2b2c auth using existing headers
		headers.AuthorizationCustomer = "Bearer " + b2b2cToken
	}

	return headers
}
