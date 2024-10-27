package repository

import "github.com/iqbalbachmid/billing-engine/domain/entity"

//go:generate mockery --name=BorrowerRepository --output=../../mocks/domain/repository --with-expecter=true
type BorrowerRepository interface {
	GetByID(id int) (entity.Borrower, error)
	GetAll() ([]entity.Borrower, error)
	Create(borrower entity.Borrower) (int, error)
	Update(borrower entity.Borrower) error
	Delete(id int) error
}
