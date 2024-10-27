package repository

import "github.com/iqbalbachmid/billing-engine/domain/entity"

//go:generate mockery --name=LoanRepository --output=../../mocks/domain/repository --with-expecter=true
type LoanRepository interface {
	GetByID(id int) (entity.Loan, error)
	GetAll() ([]entity.Loan, error)
	Create(loan entity.Loan) (int, error)
	Update(loan entity.Loan) error
	Delete(id int) error
}
