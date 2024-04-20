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

var _ = Describe("AccountBinding", func() {
	var (
		client     *directdebit.Client
		cfg        *directdebit.Config
		server     *httptest.Server
		request    *directdebit.AccountBindingRequest
		b2bToken   string
		externalID string
	)

	BeforeEach(func() {
		privkey := `-----BEGIN RSA PRIVATE KEY-----
MIIBOgIBAAJBAJ9wKHG/JDCr+6eZ6as5hHspDhZEUkveTF9GnCcDqRVmS/E4rhYw
niRfSSf2hfp516cYpC7zJ2AFB4id63T3lxECAwEAAQJALk7qQFdvEH/zYPOwTd4v
34HGKKuBZ63Sat3cXuyOQLt3OSj/lgRx/XELRjdLMADQrxTI0ijUvr9jqLB0I6O1
DQIhAMo3voVyYA0dmZvNWPRf6JzkFeexsXKmsrpNij9pwvPHAiEAyde1FPNTAeQn
e0DW+gtE5dHg/omiZDYBKLBj4jb4bmcCIDhvSjqP6wJ+CkqTCopY4eA3P23EB5PJ
tgOMdFKyP3gtAiEAgC0db2p94guTDvBEFJGndRJtAPdCSsUIw2AQbg1egi0CIH+a
Z4Mar+3f+SFIElF+M+aD4AzzibGUJghJ03cbR6Pt
-----END RSA PRIVATE KEY-----`

		cfg = &directdebit.Config{
			RsaPrivateKey: privkey,
			ClientID:      "123",
			MerchantID:    "123",
			HTTPClient:    &http.Client{},
			Logger:        slog.Default(),
		}

		request = &directdebit.AccountBindingRequest{}
		b2bToken = "sampleB2bToken"
		externalID = "sampleExternalID"
	})

	AfterEach(func() {
		if server != nil {
			server.Close()
		}
	})

	When("error occurs in generateRSASignature", func() {
		It("returns an error", func() {
			client, _ = directdebit.New(cfg)
			resp, err := client.AccountBinding(context.Background(), request, b2bToken, externalID)
			Expect(err).To(HaveOccurred())
			Expect(resp).To(BeNil())
		})
	})

	When("error occurs in json.Marshal", func() {
		It("returns an error", func() {
			client, _ = directdebit.New(cfg)
			resp, err := client.AccountBinding(context.Background(), request, b2bToken, externalID)
			Expect(err).To(HaveOccurred())
			Expect(resp).To(BeNil())
		})
	})

	When("error occurs in execute", func() {
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			}))

			cfg.EndpointBaseURL = server.URL
			client, _ = directdebit.New(cfg)
		})

		It("returns an error", func() {
			resp, err := client.AccountBinding(context.Background(), request, b2bToken, externalID)
			Expect(err).To(HaveOccurred())
			Expect(resp).To(BeNil())
		})
	})

	When("error occurs in json.Unmarshal", func() {
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`invalid json`))
			}))

			cfg.EndpointBaseURL = server.URL
			client, _ = directdebit.New(cfg)
		})

		It("returns an error", func() {
			resp, err := client.AccountBinding(context.Background(), request, b2bToken, externalID)
			Expect(err).To(HaveOccurred())
			Expect(resp).To(BeNil())
		})
	})

	When("AccountBinding is successful", func() {
		expectedResp := &directdebit.AccountBindingResponse{
			ResponseCode:       "2003100",
			ResponseMessage:    "success",
			PartnerReferenceNo: "2020102900000000000001",
			AccountToken:       "gagjzxjp9tjnduyv2zrd8qc98txt3ynb",
			TokenStatus:        "ACTIVE",
			AuthCode:           "a4sd5a4fsaf5d5f4df66ad85f4",
			UserInfo: directdebit.UserInfo{
				PublicUserID: "AYOPOP-XU56ZX",
			},
			AdditionalInfo: directdebit.AccountBindingAdditionalInfo{
				MaskedCard: "************3489",
				BankCode:   "002",
			},
		}
		BeforeEach(func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				respJSON, _ := json.Marshal(expectedResp)
				w.WriteHeader(http.StatusOK)
				w.Write(respJSON)
			}))

			cfg.EndpointBaseURL = server.URL
			cfg.MerchantID = "MEKARI"
			client, _ = directdebit.New(cfg)
		})

		It("returns the response without error", func() {
			resp, err := client.AccountBinding(context.Background(), request, b2bToken, externalID)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp).To(Equal(expectedResp))
			Expect(request.MerchantID).To(Equal("MEKARI"))
		})

		It("use merchantID from request", func() {
			request.MerchantID = "test-merchant"
			resp, err := client.AccountBinding(context.Background(), request, b2bToken, externalID)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp).To(Equal(expectedResp))
			Expect(request.MerchantID).NotTo(Equal("MEKARI"))
		})
	})
})
