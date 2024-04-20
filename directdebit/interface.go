package directdebit

import "context"

// TODO: pass context into all function.
//
//go:generate mockgen -destination=./mock/client.go -package=mock bitbucket.org/mid-kelola-indonesia/corepayment-service/pkg/ayoconnect ClientInterface
type ClientInterface interface {
	GetBusinessAccessToken(ctx context.Context) (*GetAccessTokenResponse, error)
	AccountBinding(ctx context.Context, req *AccountBindingRequest, b2bToken, externalID string) (*AccountBindingResponse, error)
	GetAuthCode(ctx context.Context, req *GetAuthCodeRequest, b2bToken, externalID string) (*GetAuthCodeResponse, error)
	GetCustomerAccessToken(ctx context.Context, authCode, accessTokenB2B string) (*GetAccessTokenResponse, error)
	Debit(ctx context.Context, req *DebitRequest, b2bToken, b2b2cToken, externalID string) (*DebitResponse, error)
	Unbind(ctx context.Context, req *AccountUnbindRequest, b2bToken, b2b2cToken, externalID string) (*AccountUnbindResponse, error)
	DebitStatus(ctx context.Context, b2bToken, debitExternalID, externalID string) (*DebitResponse, error)
}

// to check if the Client already satisfies the interface.
var _ ClientInterface = (*Client)(nil)
