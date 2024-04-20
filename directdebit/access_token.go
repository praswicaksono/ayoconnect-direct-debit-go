package directdebit

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func (c Client) GetBusinessAccessToken(ctx context.Context) (*GetAccessTokenResponse, error) {
	timestamp := time.Now().Format(time.RFC3339)
	signature, err := generateRSASignature(timestamp, c.Config.RsaPrivateKey, c.Config.ClientID)
	if err != nil {
		return nil, err
	}

	headers := c.BuildHeader(timestamp, signature, "", "", "")

	body, err := json.Marshal(
		GetBusinessAccessTokenRequest{
			GrantType:      "client_credentials",
			AdditionalInfo: GetBusinessAccessTokenAdditionalInfo{MerchantID: c.Config.MerchantID},
		},
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.Execute(ctx, http.MethodPost, GetBusinessAccessTokenEndpoint, headers, body)
	if err != nil {
		return nil, err
	}

	respEntity := GetAccessTokenResponse{}
	err = json.Unmarshal(resp, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}

func (c Client) GetAuthCode(ctx context.Context, req *GetAuthCodeRequest, b2bToken, externalID string) (*GetAuthCodeResponse, error) {
	timestamp := time.Now().Format(time.RFC3339)

	if req.MerchantID == "" {
		req.MerchantID = c.Config.MerchantID
	}
	urlValues, err := StructToURLValues(req)
	if err != nil {
		return nil, err
	}

	endpoint := GetAuthCodeEndpoint + "?" + urlValues.Encode()

	signature := generateHmacSignature("GET", GetAuthCodeEndpoint, b2bToken, "", timestamp, c.Config.ClientSecret)
	headers := c.BuildHeader(timestamp, signature, b2bToken, "", externalID)

	resp, err := c.Execute(ctx, http.MethodGet, endpoint, headers, nil)
	if err != nil {
		return nil, err
	}

	respEntity := GetAuthCodeResponse{}
	err = json.Unmarshal(resp, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}

func (c Client) GetCustomerAccessToken(ctx context.Context, authCode, accessTokenB2B string) (*GetAccessTokenResponse, error) {
	timestamp := time.Now().Format(time.RFC3339)
	signature, err := generateRSASignature(timestamp, c.Config.RsaPrivateKey, c.Config.ClientID)
	if err != nil {
		return nil, err
	}

	headers := c.BuildHeader(timestamp, signature, accessTokenB2B, "", "")

	body, err := json.Marshal(
		GetCustomerAccessTokenRequest{
			GrantType: "authorization_code",
			AuthCode:  authCode,
			AdditionalInfo: GetCustomerAccessTokenAdditionalInfo{
				MerchantID: c.Config.MerchantID,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	endpoint := GetCustomerAccessTokenEndpoint

	resp, err := c.Execute(ctx, http.MethodPost, endpoint, headers, body)
	if err != nil {
		return nil, err
	}

	respEntity := GetAccessTokenResponse{}
	err = json.Unmarshal(resp, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
