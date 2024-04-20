package directdebit_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"log/slog"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/praswicaksono/ayoconnect-direct-debit-go/directdebit"
)

var _ = Describe("GetBusinessAccessToken", func() {
	var (
		client *directdebit.Client
		cfg    *directdebit.Config
		server *httptest.Server
	)

	BeforeEach(func() {
		privkey := `-----BEGIN PRIVATE KEY-----
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
5U0gfFZB2kJ9FaBV+zZotA==
-----END PRIVATE KEY-----`

		cfg = &directdebit.Config{
			RsaPrivateKey:   privkey,
			ClientID:        "123",
			MerchantID:      "123",
			EndpointBaseURL: "", // this will be set after the server starts
			HTTPClient:      &http.Client{},
			Logger:          slog.Default(),
		}
	})

	AfterEach(func() {
		server.Close()
	})

	When("the server's response is invalid json", func() {
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`invalid json`))
			}))

			cfg.EndpointBaseURL = server.URL
			client, _ = directdebit.New(cfg)
		})

		It("returns an error", func() {
			resp, err := client.GetBusinessAccessToken(context.Background())
			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())
		})
	})

	When("successful retrieval of business access token", func() {
		expectedResp := &directdebit.GetAccessTokenResponse{
			ResponseCode:    "2001000",
			ResponseMessage: "success",
			TokenType:       "BearerToken",
			ResponseTime:    "20220122010203",
			AccessToken:     "afp5fwm796mpjhfn7nrj39ntmzgfknrn",
			ExpiredIn:       3599,
		}
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				respJSON, _ := json.Marshal(expectedResp)
				w.WriteHeader(http.StatusOK)
				w.Write(respJSON)
			}))

			cfg.EndpointBaseURL = server.URL
			client, _ = directdebit.New(cfg)
		})

		It("returns the business access token without error", func() {
			resp, err := client.GetBusinessAccessToken(context.Background())
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp).Should(Equal(expectedResp))
		})
	})

	When("the server returns a 500 error", func() {
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			}))

			cfg.EndpointBaseURL = server.URL
			client, _ = directdebit.New(cfg)
		})

		It("returns an error", func() {
			resp, err := client.GetBusinessAccessToken(context.Background())
			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())
		})
	})

	When("the server returns a 400 error", func() {
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusBadRequest)
			}))

			cfg.EndpointBaseURL = server.URL
			client, _ = directdebit.New(cfg)
		})

		It("returns an error", func() {
			resp, err := client.GetBusinessAccessToken(context.Background())
			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())
		})
	})

	When("server response is valid JSON but with unexpected structure", func() {
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"unexpected": "value"}`))
			}))

			cfg.EndpointBaseURL = server.URL
			client, _ = directdebit.New(cfg)
		})

		It("returns empty response", func() {
			resp, err := client.GetBusinessAccessToken(context.Background())
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.ResponseCode).Should(BeEmpty())
			Expect(resp.ResponseMessage).Should(BeEmpty())
			Expect(resp.ResponseTime).Should(BeEmpty())
			Expect(resp.TokenType).Should(BeEmpty())
			Expect(resp.AccessToken).Should(BeEmpty())
			Expect(resp.ExpiredIn).Should(BeZero())
		})
	})

	When("an invalid RSA private key is used", func() {
		BeforeEach(func() {
			cfg.RsaPrivateKey = "InvalidPrivateKey"
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"test": "response"}`))
			}))

			cfg.EndpointBaseURL = server.URL
			client, _ = directdebit.New(cfg)
		})

		It("returns an error", func() {
			_, err := client.GetBusinessAccessToken(context.Background())
			Expect(err).Should(HaveOccurred())
		})
	})
})

var _ = Describe("GetAuthCode", func() {
	var (
		client *directdebit.Client
		cfg    *directdebit.Config
		server *httptest.Server
	)

	BeforeEach(func() {
		privkey := `-----BEGIN PRIVATE KEY-----
MIIG/AIBADANBgkqhkiG9w0BAQEFAASCBuYwggbiAgEAAoIBgQCrb37xqR+7N0aJ
JY7juytGVC8mlQYFI3hcXZJVhnMlLsL0y2F/OBxduMXNlM30dBA9dYr4ggKdSYMu
toOMsGCgYavYemmIW0YlpTz2up+O50khPhdIHsA1R0B1v9QCDkYEYpE9RL9bDTP+
zJb/uq4QA6f/EQeuSwNUaV1s/6uBmrtk5vXSJi/AsZ5trAknVb3Onv3GXIbhaIAB
Y/e/f+19Z9y48UOJL0kSqAGhSM7Pd6AakuJ+aV2hCwhB3zbsI55bPOORRbqosBNU
wkW+xid837u+vc2SCm5FMvVrvkIwMicqhDKsQlvHpG1+VSv5S92HUwGpC5U1+XYO
Y7yHaX/KLcq1lYAttjrwRWs5EVHao8hYjjtqrmQicRBSfQHDmY2nklvylHveTzKp
j83oQxZYRyy+lqD9IjSHBRVsDcaMDfDMqzrUCWqavMxQVqnSDLgHt/CW8osuJ0UM
2o01hZB/srEE0BXCW5HoZlaW7udd8Ti7ZAs36+PP/LMs+rus0bcCAwEAAQKCAYAR
ByaPP1KxCEj/x1S9hvJB7ouuY9/ws7i5R+wIha27ND1WDjt1ZO/gWUGAbXbVgI+6
Ywn2LAexcsNOaP+BAmXemET24BXKXvKFO7fl89x0V8G6RQ4P8kn6IMUkzPR0bdGD
jvzJHqJ5G0MeXFjlNrgiTBKsMZdXNwkyIbMPaAezfFh/qbch8/wLQjkvwIY6O3h6
ZO1k/fzBt9z7BmBty3md2qqgTgp8vk8eRMTArdgo4ENtUEih8LpFjDB6Rn8Qjmr7
GkPKeAENCfLoZbIkzFguROIdnBHDWZghBxxo0MwNStOsi/DX+kJfrTgz4bzS+UEG
6Ylf64z7E68mmVfkC7M3VzvFgRqPK1oG+iu/4U7etz+oErKX4cqJnlKJgItXK46D
pmsHcJowcsgziugaity4DOj+gFie5POtN1ux9pI7Do0ZAD6HHpei53l8RpQGgBtl
+iAlPx2eQpW/uV1bapreWbYbNajT3jxkf7LtYsIsGsr17ti99pNMh6HMO4tSjckC
gcEA1WxDMo1KlydmEz7smlUFCH3rWUPdm8F8mCRzo0Q7F4QFoXe0rDLZZlt6puFn
Bez8B8bEtDmB5bYjEZhdPyhbl+HQ/4XlUo88GJUfWwpIiwWecxalFwJE5K3bLP1C
Yzgx8BarGsr4/miY1JD/Kjw8wVEgPGG6oZeBdS4Z9uqlUuiLMsIo8oSXPdkfI9Gg
Sxn2GcBGLsM0IPcdfEgXJw64VF0kGAu//Vkpv4xc5LSqKNeR28zNEUF2xyTLgYIM
SHepAoHBAM2i5KcDBVpE+iqKvXdTB85XJCnZXqG8uVXHwonC7P/UiQeZWL2xJlFV
ibxguhaLsbCG+gXzVJdwSSoKX+p0pk8G+Lf5JYhQLtQQEZUc2mSOC2e5bxhK/Lcw
yVuSFIftg0+fc6fNiYTXBX839r1wLWs+cshGMvAO0POimhw6TjVcDjbaAaTpN77Z
UW7644i2eRenPKFbNvww16jVHhIVdY7oSCAKAGivuyFvSKRghZMh/JP456OJXHQ6
lWIS019aXwKBwEX6vrnnrEqNx6GN42TjdcgICdB2OUbmFaWJZkVljP6z8mi0aJCC
B9jRLBFmHTLLNwSRv1Pc+2PH6g3N6N1ZrVbK2429aKk+gBULaIGgiJLVH9Ra230E
6HQXMaO50zfXaEByHl6lqSk6QMqKVLCTmdRFdo11+g0cMX2rxSW6YMUjrOjS0zxa
D4FfHR/Qj3+wnoppClow9XnNrWRf+v96iyRWegxMZgJ7Zv4A10DCoHzN2my45ZC/
52N7BCON8dsdKQKBwDPZyAfothfN3rqNYzrMP+KijGbU/YyQtrbPeNkdwn67i5XT
79Fc8sl9ZQ6P4TxAGxzk2/RWJ9VLpdco6IiIw0qX+m0BMJqPhU9JgfV0YgkK3Ata
cY3RkqlqbstdKTohBIQ2M4ZzSCKrySIL7XZU687n3y9qq/tl8QAN1wgZF5FS1e60
x8daWwkPaP4v2uGlCSGStLIG+vVaJ3bVzhBHQu422cDiZLoA3ZGPquRvxh6Uakix
cU8GGr7f6rzg/FVFxwKBwEkrok6ccZf3MVLnlNmOCAb7nP+ehquOVv75fRdPV+q5
B+PKt975LyM4bbmKw0AHz/jKoFTCPyDFgmA43i75VDjJ33ADALcLoPy333unC5I1
3SVZtXzV5nLeynEq9ZsvvcEvABokiqoe1Vo3evNrfRxb4meIg5spLEGsBRRoOaPO
6UtnlQQpNKvFyLx3mCwtYkcp/5KJs0WSPJakpoQRlsXjEzzmJesBxVn80ble2qDu
5U0gfFZB2kJ9FaBV+zZotA==
-----END PRIVATE KEY-----`

		cfg = &directdebit.Config{
			RsaPrivateKey:   privkey,
			ClientID:        "123",
			MerchantID:      "123",
			EndpointBaseURL: "", // this will be set after the server starts
			HTTPClient:      &http.Client{},
			Logger:          slog.Default(),
		}
	})

	AfterEach(func() {
		if server != nil {
			server.Close()
		}
	})

	When("MerchantID is not provided in the request", func() {
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				Expect(r.URL.Query().Get("merchantId")).To(Equal(cfg.MerchantID))
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"some_key": "some_value"}`))
			}))

			cfg.EndpointBaseURL = server.URL
			client, _ = directdebit.New(cfg)
		})

		It("uses the MerchantID from the config", func() {
			req := &directdebit.GetAuthCodeRequest{}
			resp, err := client.GetAuthCode(context.Background(), req, "testToken", "123")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp).ShouldNot(BeNil())
		})
	})

	When("Server returns valid response", func() {
		expectedResp := &directdebit.GetAuthCodeResponse{
			ResponseCode:    "2003000",
			ResponseMessage: "Request has been processed successfully",
			AuthCode:        "a6975f82d00a4ddc9633087fefb6275e",
			State:           "8gcz7rgqtcg8nnjkixhx46pw6wmhjprd",
		}
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				respJSON, _ := json.Marshal(expectedResp)
				w.WriteHeader(http.StatusOK)
				w.Write(respJSON)
			}))

			cfg.EndpointBaseURL = server.URL
			client, _ = directdebit.New(cfg)
		})

		It("parses the response correctly", func() {
			req := &directdebit.GetAuthCodeRequest{MerchantID: "providedMerchantID"}
			resp, err := client.GetAuthCode(context.Background(), req, "testToken", "123")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp).Should(Equal(expectedResp))
		})
	})

	When("Server returns invalid json response", func() {
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`invalid json`))
			}))

			cfg.EndpointBaseURL = server.URL
			client, _ = directdebit.New(cfg)
		})

		It("returns an error", func() {
			req := &directdebit.GetAuthCodeRequest{MerchantID: "providedMerchantID"}
			resp, err := client.GetAuthCode(context.Background(), req, "testToken", "123")
			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())
		})
	})

	When("StructToURLValues receives valid struct with proper tags", func() {
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				Expect(r.URL.RawQuery).To(ContainSubstring("lang=ID&merchantId=providedMerchantID&redirectUrl=&scopes=CARD_REGISTRATION&seamlessData=%7B%22mobileNumber%22%3A%2262822999999%22%2C%22bankCode%22%3A%22002%22%7D&state=2pzfbuurmij73ymh84efew8rtf3bt4q8"))
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"some_key": "some_value"}`))
			}))

			cfg.EndpointBaseURL = server.URL
			client, _ = directdebit.New(cfg)
		})

		It("converts struct to URL values correctly", func() {
			req := &directdebit.GetAuthCodeRequest{
				Lang:       "ID",
				MerchantID: "providedMerchantID",
				Scopes:     "CARD_REGISTRATION",
				State:      "2pzfbuurmij73ymh84efew8rtf3bt4q8",
				SeamlessData: directdebit.SeamlessData{
					MobileNumber: "62822999999",
					BankCode:     "002",
				},
			}
			resp, err := client.GetAuthCode(context.Background(), req, "testToken", "123")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp).ShouldNot(BeNil())
		})
	})

	When("Struct contains non-string field with 'url' tag", func() {
		type InvalidStruct struct {
			Field int `url:"field"`
		}

		It("returns an error", func() {
			_, err := directdebit.StructToURLValues(InvalidStruct{Field: 42})
			Expect(err).Should(HaveOccurred())
			Expect(err.Error()).To(Equal("non-string field with 'url' tag encountered"))
		})
	})

	When("Struct contains some fields without 'url' tags", func() {
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				Expect(r.URL.Query().Get("MerchantID")).To(Equal("providedMerchantID")) // Ensure the tagged field is present
				Expect(r.URL.Query().Get("NoTagField")).To(BeEmpty())                   // Ensure the non-tagged field is absent
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"some_key": "some_value"}`))
			}))

			cfg.EndpointBaseURL = server.URL
			client, _ = directdebit.New(cfg)
		})

		It("only converts fields with 'url' tags to URL values", func() {
			req := &struct {
				MerchantID string `url:"MerchantID"`
				NoTagField string
			}{
				MerchantID: "providedMerchantID",
				NoTagField: "shouldNotAppearInQuery",
			}

			values, err := directdebit.StructToURLValues(*req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(values.Get("MerchantID")).To(Equal("providedMerchantID"))
			Expect(values.Get("NoTagField")).To(BeEmpty())
		})
	})

	When("execute returns an error", func() {
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			}))

			cfg.EndpointBaseURL = server.URL
			client, _ = directdebit.New(cfg)
		})

		It("returns an error", func() {
			req := &directdebit.GetAuthCodeRequest{MerchantID: "providedMerchantID"}
			resp, err := client.GetAuthCode(context.Background(), req, "testToken", "123")
			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())
		})
	})
})

var _ = Describe("GetCustomerAccessToken", func() {
	var (
		client         *directdebit.Client
		cfg            *directdebit.Config
		server         *httptest.Server
		authCode       = "123"
		accessTokenB2B = "accessTokenB2Bsample"
	)

	BeforeEach(func() {
		privkey := `-----BEGIN PRIVATE KEY-----
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
5U0gfFZB2kJ9FaBV+zZotA==
-----END PRIVATE KEY-----`

		cfg = &directdebit.Config{
			RsaPrivateKey:   privkey,
			ClientID:        "123",
			MerchantID:      "123",
			EndpointBaseURL: "", // this will be set after the server starts
			HTTPClient:      &http.Client{},
			Logger:          slog.Default(),
		}
	})

	AfterEach(func() {
		server.Close()
	})

	When("the server's response is invalid json", func() {
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`invalid json`))
			}))

			cfg.EndpointBaseURL = server.URL
			client, _ = directdebit.New(cfg)
		})

		It("returns an error", func() {
			resp, err := client.GetCustomerAccessToken(context.Background(), authCode, accessTokenB2B)
			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())
		})
	})

	When("successful retrieval of customer access token", func() {
		expectedResp := &directdebit.GetAccessTokenResponse{
			ResponseCode:    "2002000",
			ResponseMessage: "success",
			TokenType:       "BearerToken",
			ResponseTime:    "20220122010203",
			AccessToken:     "a1b2c3d4e5f6g7h8i9j0",
			ExpiredIn:       3599,
		}
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				respJSON, _ := json.Marshal(expectedResp)
				w.WriteHeader(http.StatusOK)
				w.Write(respJSON)
			}))

			cfg.EndpointBaseURL = server.URL
			client, _ = directdebit.New(cfg)
		})

		It("returns the business access token without error", func() {
			resp, err := client.GetCustomerAccessToken(context.Background(), authCode, accessTokenB2B)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp).Should(Equal(expectedResp))
		})
	})

	When("the server returns a 500 error", func() {
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			}))

			cfg.EndpointBaseURL = server.URL
			client, _ = directdebit.New(cfg)
		})

		It("returns an error", func() {
			resp, err := client.GetCustomerAccessToken(context.Background(), authCode, accessTokenB2B)
			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())
		})
	})

	When("the server returns a 400 error", func() {
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusBadRequest)
			}))

			cfg.EndpointBaseURL = server.URL
			client, _ = directdebit.New(cfg)
		})

		It("returns an error", func() {
			resp, err := client.GetCustomerAccessToken(context.Background(), authCode, accessTokenB2B)
			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())
		})
	})

	When("server response is valid JSON but with unexpected structure", func() {
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"unexpected": "value"}`))
			}))

			cfg.EndpointBaseURL = server.URL
			client, _ = directdebit.New(cfg)
		})

		It("returns empty response", func() {
			resp, err := client.GetCustomerAccessToken(context.Background(), authCode, accessTokenB2B)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.ResponseCode).Should(BeEmpty())
			Expect(resp.ResponseMessage).Should(BeEmpty())
			Expect(resp.ResponseTime).Should(BeEmpty())
			Expect(resp.TokenType).Should(BeEmpty())
			Expect(resp.AccessToken).Should(BeEmpty())
			Expect(resp.ExpiredIn).Should(BeZero())
		})
	})

	When("an invalid RSA private key is used", func() {
		BeforeEach(func() {
			cfg.RsaPrivateKey = "InvalidPrivateKey"
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"test": "response"}`))
			}))

			cfg.EndpointBaseURL = server.URL
			client, _ = directdebit.New(cfg)
		})

		It("returns an error", func() {
			_, err := client.GetCustomerAccessToken(context.Background(), authCode, accessTokenB2B)
			Expect(err).Should(HaveOccurred())
		})
	})
})
