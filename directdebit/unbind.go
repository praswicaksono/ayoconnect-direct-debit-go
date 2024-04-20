package directdebit

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func (c Client) Unbind(
	ctx context.Context,
	req *AccountUnbindRequest,
	b2bToken,
	b2b2cToken,
	externalID string,
) (*AccountUnbindResponse, error) {
	endpoint := UnbindEndpoint
	timestamp := time.Now().Format(time.RFC3339)

	if req.MerchantID == "" {
		req.MerchantID = c.Config.MerchantID
	}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	signature := generateHmacSignature("POST", endpoint, b2bToken, string(body), timestamp, c.Config.ClientSecret)
	headers := c.BuildHeader(timestamp, signature, b2bToken, b2b2cToken, externalID)

	resp, err := c.Execute(ctx, http.MethodPost, endpoint, headers, body)
	if err != nil {
		return nil, err
	}

	respEntity := AccountUnbindResponse{}
	err = json.Unmarshal(resp, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
