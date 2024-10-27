package application

import (
	"errors"
	"github.com/iqbalbachmid/billing-engine/domain/entity"
	"github.com/iqbalbachmid/billing-engine/domain/repository"
	"time"
)

const OverdueLimit = 2

type LoanService struct {
	loanRepo         repository.LoanRepository
	loanScheduleRepo repository.LoanScheduleRepository
	paymentRepo      repository.PaymentRepository
	timeNow          func() time.Time
}

func NewLoanService(
	loanRepo repository.LoanRepository,
	loanScheduleRepo repository.LoanScheduleRepository,
	paymentRepo repository.PaymentRepository,
	timeNow func() time.Time,
) *LoanService {
	return &LoanService{
		loanRepo:         loanRepo,
		loanScheduleRepo: loanScheduleRepo,
		paymentRepo:      paymentRepo,
		timeNow:          timeNow,
	}
}

func (s *LoanService) GetOutstanding(loanID int) (float64, error) {
	schedules, err := s.loanScheduleRepo.GetByLoanID(loanID)
	if err != nil {
		return 0, err
	}

	outstanding := 0.0
	for _, schedule := range schedules {
		if !schedule.IsPaid() {
			outstanding += schedule.TotalDue
		}
	}

	return outstanding, nil
}

func (s *LoanService) IsDelinquent(loanID int) (bool, error) {
	schedules, err := s.loanScheduleRepo.GetByLoanID(loanID)
	if err != nil {
		return false, err
	}

	var overdueCounter int
	for _, schedule := range schedules {
		if schedule.IsOverdue() {
			overdueCounter++
		}
		if overdueCounter == OverdueLimit {
			return true, nil
		}
	}

	return false, nil
}

func (s *LoanService) MakePayment(loanID int, paymentAmount float64, paymentMethod string) error {
	schedules, err := s.loanScheduleRepo.GetByLoanID(loanID)
	if err != nil {
		return err
	}

	if err = s.validate(schedules, loanID, paymentAmount); err != nil {
		return err
	}

	payment := entity.Payment{
		LoanID:        loanID,
		PaymentDate:   s.timeNow(),
		AmountPaid:    paymentAmount,
		PaymentMethod: paymentMethod,
		Status:        entity.Status,
	}

	var loanSchedulesToBeUpdated []entity.LoanSchedule
	for _, schedule := range schedules {
		if !schedule.IsPaid() {
			schedule.PaymentStatus = entity.PaymentStatusPaid
			loanSchedulesToBeUpdated = append(loanSchedulesToBeUpdated, schedule)
			paymentAmount -= schedule.TotalDue
		}
		if paymentAmount == 0.0 {
			break
		}
	}

	if _, err = s.paymentRepo.CreatePaymentAndUpdateLoanSchedules(payment, loanSchedulesToBeUpdated); err != nil {
		return err
	}

	return nil
}

func (s *LoanService) validate(
	schedules []entity.LoanSchedule,
	loanID int,
	paymentAmount float64,
) error {
	outstanding, err := s.GetOutstanding(loanID)
	if err != nil {
		return err
	}
	if paymentAmount > outstanding {
		return errors.New("payment amount exceeds outstanding balance")
	}

	if int(paymentAmount)%int(schedules[0].TotalDue) != 0 {
		return errors.New("payment is not a valid multiple of schedule amount")
	}

	return nil
}
