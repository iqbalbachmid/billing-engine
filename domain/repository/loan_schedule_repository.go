package repository

import "github.com/iqbalbachmid/billing-engine/domain/entity"

//go:generate mockery --name=LoanScheduleRepository --output=../../mocks/domain/repository --with-expecter=true
type LoanScheduleRepository interface {
	GetByLoanID(loanID int) ([]entity.LoanSchedule, error)
	Create(schedule entity.LoanSchedule) (int, error)
	Update(schedule entity.LoanSchedule) error
}
