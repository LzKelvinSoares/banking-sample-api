package app

import (
	"banking/constants"
	"banking/dto"
	"banking/errs"
	"banking/mocks/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

var router *mux.Router
var ch CustomerHandler
var mockService *service.MockCustomerService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = service.NewMockCustomerService(ctrl)
	ch = CustomerHandler{
		service: mockService,
	}

	router = mux.NewRouter()
	return func() {
		router = nil
		defer ctrl.Finish()
	}

}

func Test_should_return_list_of_customers_with_status_200(t *testing.T) {
	// Arrange
	tearDown := setup(t)
	defer tearDown()

	dummyCustomers := []dto.CustomerResponse{
		{Id: "1001", Name: "Kelvin Luz", City: "João Pessoa", Zipcode: "58075000", DateofBirth: "30/07/1991", Status: "1"},
		{Id: "1002", Name: "Teste Luz", City: "João Pessoa", Zipcode: "58075000", DateofBirth: "08/03/1991", Status: "1"},
	}
	mockService.EXPECT().GetAllCustomers(nil).Return(dummyCustomers, nil)

	router.HandleFunc(constants.GET_CUSTOMERS, ch.GetAllCustomers)

	request, _ := http.NewRequest(http.MethodGet, constants.GET_CUSTOMERS, nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Faile while testing the status code")
	}
}

func Test_should_return_status_code_500_with_error_message(t *testing.T) {
	// Arrange
	tearDown := setup(t)
	defer tearDown()

	mockService.EXPECT().GetAllCustomers(nil).Return(nil, errs.NewUnexpectedError("Some DB error"))

	router.HandleFunc(constants.GET_CUSTOMERS, ch.GetAllCustomers)

	request, _ := http.NewRequest(http.MethodGet, constants.GET_CUSTOMERS, nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Faile while testing the status code")
	}

}
