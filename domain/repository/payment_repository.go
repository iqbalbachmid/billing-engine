package repository

import "github.com/iqbalbachmid/billing-engine/domain/entity"

//go:generate mockery --name=PaymentRepository --output=../../mocks/domain/repository --with-expecter=true
type PaymentRepository interface {
	GetByLoanID(loanID int) ([]entity.Payment, error)
	CreatePaymentAndUpdateLoanSchedules(payment entity.Payment, loanSchedules []entity.LoanSchedule) (int, error)
}
