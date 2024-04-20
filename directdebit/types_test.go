package directdebit_test

import (
	"testing"

	"github.com/praswicaksono/ayoconnect-direct-debit-go/directdebit"
)

func TestResponseErrorReturnCorrectError(t *testing.T) {
	var err error

	resp := directdebit.ResponseError{ResponseCode: "4033305", ResponseMessage: "Invalid Param", ResponseDescription: "Invalid Email", StatusCode: 400}
	err = &resp

	if err.Error() != resp.ResponseMessage {
		t.Errorf("Expected %v, but got %v", resp.ResponseMessage, err.Error())
	}

	if !directdebit.IsDebitCardDisabledError(resp.ResponseCode) {
		t.Errorf("Expected %v, got %v", directdebit.IsDebitCardDisabledError(resp.ResponseCode), false)
	}
}

func TestClientSideErrorReturnCorrectError(t *testing.T) {
	var err error

	resp := directdebit.ResponseError{ResponseCode: "4003308", ResponseMessage: "Invalid Param", ResponseDescription: "Invalid Email", StatusCode: 400}
	err = &resp

	if err.Error() != resp.ResponseMessage {
		t.Errorf("Expected %v, but got %v", resp.ResponseMessage, err.Error())
	}

	if !directdebit.IsClientSideError(resp.ResponseCode) {
		t.Errorf("Expected %v, got %v", directdebit.IsClientSideError(resp.ResponseCode), false)
	}
}

func TestCardexpiredResponseCode(t *testing.T) {
	resp := directdebit.ResponseError{ResponseCode: "5000000", ResponseMessage: "Card expired", ResponseDescription: "Card expired", StatusCode: 400}

	if !directdebit.IsCardLinkageTimeoutError(resp.ResponseCode) {
		t.Errorf("Expected %v, got %v", true, false)
	}
}
