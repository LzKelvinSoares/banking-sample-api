package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_account_type_is_not_checking_or_saving(t *testing.T) {
	// AAA: Arrange Act Assert
	// Arrange
	request := NewAccountRequest{
		AccountType: "invalid",
		Amount:      5000,
	}

	// Act
	appErr := request.Validate()

	// Assert
	if appErr.Message != "Account type should be checking or saving" {
		t.Error("Invalid Message while testing account type")
	}

	if appErr.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid Code while testing account type")
	}
}

func Test_should_return_error_when_amount_is_less_than_five_thousand(t *testing.T) {
	request := NewAccountRequest{
		Amount: 4999,
	}

	appErr := request.Validate()
	if appErr.Message != "To open new account, you need to deposit at least 5000.00" {
		t.Error("Invalid Message while validating amount")
	}

}
