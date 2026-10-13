package dto

import (
	"banking/constants"
	"net/http"
	"testing"
)

func Test_should_return_error_when_transaction_type_is_not_deposit_or_withdrawal(t *testing.T) {
	// AAA: Arrange Act Assert
	// Arrange
	request := NewTransactionRequest{
		TransactionType: "invalid",
		Amount:          5,
	}

	// Act
	appErr := request.Validate()

	// Assert
	if appErr.Message != "Transaction type should be Withdrawal or Deposit" {
		t.Error("Invalid Message while testing transaction type")
	}

	if appErr.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid Code while testing transaction type")
	}
}

func Test_should_return_error_when_amount_is_less_than_zero(t *testing.T) {
	request := NewTransactionRequest{
		TransactionType: constants.DEPOSIT,
		Amount:          -10,
	}

	appErr := request.Validate()
	if appErr.Message != "The new transaction cannot be under zero" {
		t.Error("Invalid Message while validating amount")
	}

}
