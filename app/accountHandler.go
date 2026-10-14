package app

import (
	"banking/dto"
	"banking/service"
	"encoding/json"
	"net/http"

	"github.com/LzKelvinSoares/lzk-go-lib/errs"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service            service.AccountService
	transactionService service.TransactionService
}

func (ah AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appErr := ah.service.NewAccount(request)
		if appErr != nil {
			writeResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}
}

func (ah AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	transactionType := vars["transaction_type"]
	var request dto.NewTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.AccountId = accountId
		request.TransactionType = transactionType
		var transaction *dto.NewTransactionResponse
		var appErr *errs.AppError

		if request.IsTransactionTypeWithdrawal() {
			transaction, appErr = ah.transactionService.Withdrawal(request)
		} else {
			transaction, appErr = ah.transactionService.Deposit(request)
		}

		if appErr != nil {
			writeResponse(w, appErr.Code, appErr.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, transaction)
		}
	}
}
