package directdebit_test

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/praswicaksono/ayoconnect-direct-debit-go/directdebit"
)

var _ = Describe("DebitStatus", func() {
	var (
		client     *directdebit.Client
		cfg        *directdebit.Config
		server     *httptest.Server
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

		b2bToken = "sampleB2bToken"
		externalID = "sampleExternalID"
	})

	AfterEach(func() {
		if server != nil {
			server.Close()
		}
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
			resp, err := client.DebitStatus(context.Background(), b2bToken, externalID, externalID)
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
			resp, err := client.DebitStatus(context.Background(), b2bToken, externalID, externalID)
			Expect(err).To(HaveOccurred())
			Expect(resp).To(BeNil())
		})
	})

	When("debit status successful", func() {
		It("return correct response", func() {
			expectedResp := &directdebit.DebitResponse{
				ResponseCode:       "2003330",
				ResponseMessage:    "success",
				ReferenceNo:        "t4tn57kibeunbam9dtr89urv8h2jbem9",
				PartnerReferenceNo: "t4tn57kibeunbam9dtr89urv8h2jbem9",
				Amount: directdebit.Amount{
					Value:    "1000000",
					Currency: "IDR",
				},
				AdditionalInfo: directdebit.DebitAdditionalInfo{
					PublicUserID:  "TEST",
					Remarks:       "Snap Payment",
					PaymentResult: "success",
				},
			}
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				respJSON, _ := json.Marshal(expectedResp)
				w.WriteHeader(http.StatusOK)
				w.Write(respJSON)
			}))

			cfg.EndpointBaseURL = server.URL
			cfg.MerchantID = "MEKARI"
			client, _ = directdebit.New(cfg)

			resp, err := client.DebitStatus(context.Background(), b2bToken, externalID, externalID)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.AdditionalInfo.PublicUserID).Should(Equal("TEST"))
		})
	})
})
