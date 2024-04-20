package directdebit

import (
	"net/http"

	"golang.org/x/exp/slices"

	"log/slog"
)

var (
	CardExpiredResponseCode = []string{
		"4033318",
		"4033305",
		"4033307",
	}

	ClientSideErrorHTTPCode = []string{
		"400",
		"412",
		"401",
		"404",
		"500",
	}

	ClientSideErrorResponseCode = []string{
		"4033320", // validation or ayoconnect configuration error
	}

	CardLinkageTimeoutResponseCode = []string{
		"5000000",
	}
)

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

type SeamlessData struct {
	MobileNumber string `json:"mobileNumber"`
	BankCode     string `json:"bankCode,omitempty"`
}

type GetAuthCode struct {
	Client Client
}

type GetAuthCodeRequest struct {
	RedirectURL        string       `url:"redirectUrl"`
	FailureRedirectURL string       `url:"failureRedirectUrl"`
	Scopes             string       `url:"scopes"`
	State              string       `url:"state"`
	Lang               string       `url:"lang"`
	MerchantID         string       `url:"merchantId"`
	SeamlessData       SeamlessData `url:"seamlessData"`
}

type GetAuthCodeResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	AuthCode        string `json:"authCode"`
	State           string `json:"state"`
}

type GetCardsRequest struct {
	PartnerReferenceNo string
	MerchantID         string
	PublicUserID       string
}

type DebitTransactionRequest struct {
	PartnerReferenceNo string
	MerchantID         string
	BankCardToken      string
}

type GetBusinessAccessTokenAdditionalInfo struct {
	MerchantID string `json:"merchantId"`
}

type GetBusinessAccessTokenRequest struct {
	GrantType      string                               `json:"grantType"`
	AdditionalInfo GetBusinessAccessTokenAdditionalInfo `json:"additionalInfo"`
}

type GetCustomerAccessTokenAdditionalInfo struct {
	MerchantID string `json:"merchantId"`
}

type GetCustomerAccessTokenRequest struct {
	GrantType      string                               `json:"grantType"`
	AuthCode       string                               `json:"authCode"`
	AdditionalInfo GetCustomerAccessTokenAdditionalInfo `json:"additionalInfo"`
}

type GetAccessTokenResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	TokenType       string `json:"tokenType"`
	ResponseTime    string `json:"responseTime"`
	AccessToken     string `json:"accessToken"`
	ExpiredIn       int    `json:"expiresIn"`
}

type RequestHeader struct {
	Authorization         string
	AuthorizationCustomer string
	Timestamp             string
	ClientKey             string
	Signature             string
	PartnerID             string
	ExternalID            string
	ChannelID             string
}

type Client struct {
	Config *Config
}

type AccountBindingRequest struct {
	PartnerReferenceNo string `json:"partnerReferenceNo"`
	AuthCode           string `json:"authCode"`
	MerchantID         string `json:"merchantId"`
}

type UserInfo struct {
	PublicUserID string `json:"publicUserId"`
}

type AccountBindingAdditionalInfo struct {
	MaskedCard string `json:"maskedCard"`
	BankCode   string `json:"bankCode"`
}

type AccountBindingResponse struct {
	ResponseCode       string                       `json:"responseCode"`
	ResponseMessage    string                       `json:"responseMessage"`
	PartnerReferenceNo string                       `json:"partnerReferenceNo"`
	AccountToken       string                       `json:"accountToken"`
	TokenStatus        string                       `json:"tokenStatus"`
	AuthCode           string                       `json:"authCode"`
	UserInfo           UserInfo                     `json:"userInfo"`
	AdditionalInfo     AccountBindingAdditionalInfo `json:"additionalInfo"`
}

type DebitResponse struct {
	ResponseCode       string              `json:"responseCode"`
	ResponseMessage    string              `json:"responseMessage"`
	PartnerReferenceNo string              `json:"partnerReferenceNo"`
	ReferenceNo        string              `json:"referenceNo"`
	Amount             Amount              `json:"amount"`
	AdditionalInfo     DebitAdditionalInfo `json:"additionalInfo"`
}

type DebitRequest struct {
	PartnerReferenceNo string              `json:"partnerReferenceNo"`
	BankCardToken      string              `json:"bankCardToken"`
	MerchantID         string              `json:"merchantId"`
	URLParam           []URLParam          `json:"urlParam"`
	Amount             Amount              `json:"amount"`
	AdditionalInfo     DebitAdditionalInfo `json:"additionalInfo"`
}

type DebitAdditionalInfo struct {
	PublicUserID  string `json:"publicUserId"`
	Remarks       string `json:"remarks"`
	BankCode      string `json:"bankCode,omitempty"`
	OtpAllowed    string `json:"otpAllowed,omitempty"`
	PaymentResult string `json:"paymentResult,omitempty"`
}

type URLParam struct {
	URL        string `json:"url"`
	Type       string `json:"type"`
	IsDeepLink string `json:"isDeepLink"`
}

type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type AccountUnbindRequest struct {
	PartnerReferenceNo string                             `json:"partnerReferenceNo"`
	MerchantID         string                             `json:"merchantId"`
	AdditionalInfo     AccountUnbindRequestAdditionalInfo `json:"additionalInfo"`
}

type AccountUnbindRequestAdditionalInfo struct {
	PublicUserID string `json:"publicUserId"`
	AccountToken string `json:"accountToken"`
	BankCode     string `json:"bankCode"`
}

type AccountUnbindResponse struct {
	ResponseCode       string                              `json:"responseCode"`
	ResponseMessage    string                              `json:"responseMessage"`
	PartnerReferenceNo string                              `json:"partnerReferenceNo"`
	ReferenceNo        string                              `json:"referenceNo"`
	UnlinkResult       string                              `json:"unlinkResult"`
	AdditionalInfo     AccountUnbindResponseAdditionalInfo `json:"additionalInfo"`
}

type AccountUnbindResponseAdditionalInfo struct {
	UnlinkOtpToken string `json:"unlinkOtpToken"`
}

type ResponseError struct {
	ResponseCode        string `json:"responseCode"`
	ResponseMessage     string `json:"responseMessage"`
	ResponseDescription string `json:"responseDescription"`
	StatusCode          int
}

func (e *ResponseError) Error() string {
	return e.ResponseMessage
}

func IsDebitCardDisabledError(responseCode string) bool {
	return slices.Contains(CardExpiredResponseCode, responseCode)
}

func IsClientSideError(responseCode string) bool {
	if slices.Contains(ClientSideErrorResponseCode, responseCode) {
		return true
	}
	httpCode := responseCode[0:3]

	return slices.Contains(ClientSideErrorHTTPCode, httpCode)
}

func IsCardLinkageTimeoutError(responseCode string) bool {
	return slices.Contains(CardLinkageTimeoutResponseCode, responseCode)
}
