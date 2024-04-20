package directdebit

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func (c Client) AccountBinding(ctx context.Context, req *AccountBindingRequest, b2bToken, externalID string) (*AccountBindingResponse, error) {
	endpoint := AccountBindingEndpoint
	timestamp := time.Now().Format(time.RFC3339)

	if req.MerchantID == "" {
		req.MerchantID = c.Config.MerchantID
	}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	signature := generateHmacSignature("POST", endpoint, b2bToken, string(body), timestamp, c.Config.ClientSecret)
	headers := c.BuildHeader(timestamp, signature, b2bToken, "", externalID)

	resp, err := c.Execute(ctx, http.MethodPost, endpoint, headers, body)
	if err != nil {
		return nil, err
	}

	respEntity := AccountBindingResponse{}
	err = json.Unmarshal(resp, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
