Ayoconnect Direct Debit Go SDK
===
This library contains ayoconnect direct debit v2 sdk for Go.

Overview
===

API that have implemented in this sdk:

| Feature                        | Implemented        |
|--------------------------------|--------------------|
| Generate B2B Access Token      | :white_check_mark: |
| Generate B2B2C Access Token    | :white_check_mark: |
| Get Auth Code                  | :white_check_mark: |
| OTP Verification               | :x:                |
| Registration Account Binding   | :white_check_mark: |
| Registration Account Unbinding | :white_check_mark: |
| Debit Charge Host To Host      | :white_check_mark: |
| Debit Host To Host             | :white_check_mark: |
| Get Card List                  | :x:                |

Note: OTP Based Account Binding / Unbinding / Debit currently not implemented yet. any PR are welcome.

Installation
===

```
go get github.com/praswicaksono/ayoconnect-direct-debit-go
```

Then simply use import

```
import github.com/praswicaksono/ayoconnect-direct-debit-go/directdebit
```

If you find any error please run `go mod tidy` to clear your go.mod file

Example
===

```go
// private key must be in PKCS8 format
privkey := `-----BEGIN RSA PRIVATE KEY-----
MIIBOgIBAAJBAJ9wKHG/JDCr+6eZ6as5hHspDhZEUkveTF9GnCcDqRVmS/E4rhYw
niRfSSf2hfp516cYpC7zJ2AFB4id63T3lxECAwEAAQJALk7qQFdvEH/zYPOwTd4v
34HGKKuBZ63Sat3cXuyOQLt3OSj/lgRx/XELRjdLMADQrxTI0ijUvr9jqLB0I6O1
DQIhAMo3voVyYA0dmZvNWPRf6JzkFeexsXKmsrpNij9pwvPHAiEAyde1FPNTAeQn
e0DW+gtE5dHg/omiZDYBKLBj4jb4bmcCIDhvSjqP6wJ+CkqTCopY4eA3P23EB5PJ
tgOMdFKyP3gtAiEAgC0db2p94guTDvBEFJGndRJtAPdCSsUIw2AQbg1egi0CIH+a
Z4Mar+3f+SFIElF+M+aD4AzzibGUJghJ03cbR6Pt
-----END RSA PRIVATE KEY-----`

// you can pass custom http client and logger here
cfg = &directdebit.Config{
    RsaPrivateKey: privkey,
    ClientID:      "123",
    MerchantID:    "123",
    HTTPClient:    &http.Client{},
    Logger:        slog.Default(),
}

request = &directdebit.DebitRequest{}
b2bToken = "sampleB2bToken"
externalID = "sampleExternalID"
b2b2cToken = "sampleB2b2cToken"

client, _ = directdebit.New(cfg)
resp, err := client.Debit(context.Background(), request, b2bToken, b2b2cToken, externalID)

if err != nil {
    panic(err)
}
```

Contributing
===

If you would like to contribue please read our contributing guidelines. Any form of contribution is welcome.
