package example

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/praswicaksono/ayoconnect-direct-debit-go/directdebit"
)

func genearetB2BToken() string {
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
