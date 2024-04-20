package directdebit

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func (c Client) DebitStatus(
	ctx context.Context,
	b2bToken,
	debitTxExternalID,
	externalID string,
) (*DebitResponse, error) {
	timestamp := time.Now().Format(time.RFC3339)

	signature := generateHmacSignature("GET", DebitStatusEndpoint, b2bToken, "", timestamp, c.Config.ClientSecret)
	headers := c.BuildHeader(timestamp, signature, b2bToken, "", externalID)

	endpoint := DebitStatusEndpoint + "?XExternalId=" + debitTxExternalID + "&merchantId=" + c.Config.MerchantID

	resp, err := c.Execute(ctx, http.MethodGet, endpoint, headers, nil)
	if err != nil {
		return nil, err
	}

	respEntity := DebitResponse{}
	err = json.Unmarshal(resp, &respEntity)
	if err != nil {
		return nil, err
	}

	return &respEntity, nil
}
