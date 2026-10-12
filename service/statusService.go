package service

import (
	"banking/domain"
	"banking/errs"
)

type StatusService interface {
	ParseStatus(status string) (domain.Status, *errs.AppError)
}

type DefaultStatusService struct{}

func (s DefaultStatusService) ParseStatus(status string) (domain.Status, *errs.AppError) {
	switch status {
	case "active":
		return domain.StatusActive, nil
	case "inactive":
		return domain.StatusInactive, nil
	default:
		return 0, errs.NewUnexpectedError("status inválido: " + status)
	}
}

func NewStatusService() DefaultStatusService {
	return DefaultStatusService{}
}
