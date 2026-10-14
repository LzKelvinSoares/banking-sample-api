package service

import (
	realdomain "banking/domain"
	"banking/dto"
	"banking/mocks/domain"
	"testing"

	"github.com/LzKelvinSoares/lzk-go-lib/errs"

	"github.com/golang/mock/gomock"
)

var mockRepo *domain.MockAccountRepository
var service AccountService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockRepo = domain.NewMockAccountRepository(ctrl)
	service = NewAccountService(mockRepo)
	return func() {
		service = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_a_validation_error_response_when_the_request_is_not_validated(t *testing.T) {
	request := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      0,
	}
	service := NewAccountService(nil)

	_, appErr := service.NewAccount(request)

	if appErr == nil {
		t.Error("Failed while testing the new account validation")
	}
}

func Test_should_return_an_error_from_the_server_side_if_the_new_account_cannot_be_created(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	request := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      6000,
	}
	account := realdomain.NewAccount(request.CustomerId, request.AccountType, request.Amount)
	mockRepo.EXPECT().Save(account).Return(nil, errs.NewUnexpectedError("Unexpected DB error"))

	_, appErr := service.NewAccount(request)

	if appErr == nil {
		t.Error("Failed while validating error for new account")
	}

}

func Test_should_return_new_account_response_when_a_new_account_is_saved_successfully(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	request := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      6000,
	}
	account := realdomain.NewAccount(request.CustomerId, request.AccountType, request.Amount)
	accountWithId := account
	accountWithId.AccountId = "201"
	mockRepo.EXPECT().Save(account).Return(&accountWithId, nil)

	newAccount, appErr := service.NewAccount(request)

	if appErr != nil {
		t.Error("Failed while creating new account")
	}

	if newAccount.AccountId != accountWithId.AccountId {
		t.Error("Failed while matching new account id")
	}
}
