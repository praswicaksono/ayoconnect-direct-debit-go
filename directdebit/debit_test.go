package directdebit_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/praswicaksono/ayoconnect-direct-debit-go/directdebit"
)

var _ = Describe("Debit", func() {
	var (
		client     *directdebit.Client
		cfg        *directdebit.Config
		server     *httptest.Server
		request    *directdebit.DebitRequest
		b2bToken   string
		b2b2cToken string
		externalID string
	)

	BeforeEach(func() {
		cfg = &directdebit.Config{
			ClientID:   "123",
			MerchantID: "123",
			HTTPClient: &http.Client{},
		}

		request = &directdebit.DebitRequest{}
		b2bToken = "sampleB2bToken"
		b2b2cToken = "sampleB2b2cToken"
		externalID = "sampleExternalID"
	})

	AfterEach(func() {
		if server != nil {
			server.Close()
		}
	})

	When("debit error", func() {
		It("return error when failed to execute http", func() {
			client, _ = directdebit.New(cfg)
			resp, err := client.Debit(context.Background(), request, b2bToken, b2b2cToken, externalID)
			Expect(err).To(HaveOccurred())
			Expect(resp).To(BeNil())
		})

		It("return error when failed to unmarshal JSON", func() {
			server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`invalid json`))
			}))

			cfg.EndpointBaseURL = server.URL
			client, _ = directdebit.New(cfg)

			resp, err := client.Debit(context.Background(), request, b2bToken, b2b2cToken, externalID)
			Expect(err).To(HaveOccurred())
			Expect(err).Should(MatchError("invalid character 'i' looking for beginning of value"))
			Expect(resp).To(BeNil())
		})
	})

	When("debit successful", func() {
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

			resp, err := client.Debit(context.Background(), request, b2bToken, b2b2cToken, externalID)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.AdditionalInfo.PublicUserID).Should(Equal("TEST"))
		})
	})
})
