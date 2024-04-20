# Ayoconnect Direct Debit Go SDK

This library contains ayoconnect direct debit v2 sdk for Go.

# Overview

API that have implemented in this sdk:

| Feature                        | Implemented        |
| ------------------------------ | ------------------ |
| Generate B2B Access Token      | :white_check_mark: |
| Generate B2B2C Access Token    | :white_check_mark: |
| Get Auth Code                  | :white_check_mark: |
| OTP Verification               | :x:                |
| Registration Account Binding   | :white_check_mark: |
| Registration Account Unbinding | :white_check_mark: |
| Debit Charge Host To Host      | :white_check_mark: |
| Get Card List                  | :x:                |

List Of Public API

```go
type ClientInterface interface {
	GetBusinessAccessToken(ctx context.Context) (*GetAccessTokenResponse, error)
	AccountBinding(ctx context.Context, req *AccountBindingRequest, b2bToken, externalID string) (*AccountBindingResponse, error)
	GetAuthCode(ctx context.Context, req *GetAuthCodeRequest, b2bToken, externalID string) (*GetAuthCodeResponse, error)
	GetCustomerAccessToken(ctx context.Context, authCode, accessTokenB2B string) (*GetAccessTokenResponse, error)
	Debit(ctx context.Context, req *DebitRequest, b2bToken, b2b2cToken, externalID string) (*DebitResponse, error)
	Unbind(ctx context.Context, req *AccountUnbindRequest, b2bToken, b2b2cToken, externalID string) (*AccountUnbindResponse, error)
	DebitStatus(ctx context.Context, b2bToken, debitExternalID, externalID string) (*DebitResponse, error)
}
```

Note: OTP Based Account Binding / Unbinding / Debit currently not implemented yet. any PR are welcome.

# Flow

## Card Binding Flow

List of API Used:

- Generate B2B Access Token
- Get Auth Code
- Registration Account Binding

![Card Binding](https://storage.googleapis.com/dd-ui-static-dev/api-flows/cardBindingV2Flow.jpg)

## Debit Flow (Non OTP)

List of API Used:

- Generate B2B Access Token
- Generate B2B2C Access Token
- Debit Charge Host To Host

![Debit](https://storage.googleapis.com/dd-ui-static-dev/api-flows/chargePaymentV2Flow.jpg)

# Installation

```
go get github.com/praswicaksono/ayoconnect-direct-debit-go
```

Then simply use import

```
import github.com/praswicaksono/ayoconnect-direct-debit-go/directdebit
```

If you find any error please run `go mod tidy` to clear your go.mod file

# Configuration

You may configure the sdk by passing a `Config` struct to the `New` function. Here list configuration options:

```go
type Config struct {
	ClientID        string
	ClientSecret    string
	MerchantID      string
	RsaPrivateKey   string
	EndpointBaseURL string
	ChannelID       string
	Logger          *slog.Logger
	HTTPClient      *http.Client
}
```

To instantiante new client you can pass configuration to `New` function

```go
package example

import "github.com/praswicaksono/ayoconnect-direct-debit-go/directdebit"

func main() {
    cfg := &directdebit.Config{
		RsaPrivateKey: privkey,
		ClientID:      "123",
		MerchantID:    "123",
		HTTPClient:    &http.Client{},
		Logger:        slog.Default(),
	}

	client, _ := directdebit.New(cfg)
}
```

# Example

Get B2B Access Token

```go
package example

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/praswicaksono/ayoconnect-direct-debit-go/directdebit"
)

func genearetB2BToken() string {
    // private key must be in PKCS8 Format
	privkey := `-----BEGIN RSA PRIVATE KEY-----
MIIBOgIBAAJBAJ9wKHG/JDCr+6eZ6as5hHspDhZEUkveTF9GnCcDqRVmS/E4rhYw
niRfSSf2hfp516cYpC7zJ2AFB4id63T3lxECAwEAAQJALk7qQFdvEH/zYPOwTd4v
34HGKKuBZ63Sat3cXuyOQLt3OSj/lgRx/XELRjdLMADQrxTI0ijUvr9jqLB0I6O1
DQIhAMo3voVyYA0dmZvNWPRf6JzkFeexsXKmsrpNij9pwvPHAiEAyde1FPNTAeQn
e0DW+gtE5dHg/omiZDYBKLBj4jb4bmcCIDhvSjqP6wJ+CkqTCopY4eA3P23EB5PJ
tgOMdFKyP3gtAiEAgC0db2p94guTDvBEFJGndRJtAPdCSsUIw2AQbg1egi0CIH+a
Z4Mar+3f+SFIElF+M+aD4AzzibGUJghJ03cbR6Pt
-----END RSA PRIVATE KEY-----`

	cfg := &directdebit.Config{
		RsaPrivateKey: privkey,
		ClientID:      "123",
		MerchantID:    "123",
		HTTPClient:    &http.Client{},
		Logger:        slog.Default(),
	}

	client, _ := directdebit.New(cfg)
	resp, err := client.GetBusinessAccessToken(context.Background())
	if err != nil {
		slog.Error(err.Error())
		return ""
	}

	return resp.AccessToken
}
```

# Contributing

If you would like to contribue please read our contributing guidelines. Any form of contribution is welcome.
