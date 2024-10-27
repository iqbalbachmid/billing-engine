package entity

import "time"

const (
	PaymentStatusUnspecified = "unspecified"
	PaymentStatusDue         = "due"
	PaymentStatusPaid        = "paid"
	PaymentStatusOverdue     = "overdue"
)

type LoanSchedule struct {
	ScheduleID      int       `db:"schedule_id"`
	LoanID          int       `db:"loan_id"`
	DueDate         time.Time `db:"due_date"`
	PrincipalAmount float64   `db:"principal_amount"`
	InterestAmount  float64   `db:"interest_amount"`
	TotalDue        float64   `db:"total_due"`
	PaymentStatus   string    `db:"payment_status"`
}

func (l *LoanSchedule) IsUnspecified() bool {
	return l.PaymentStatus == PaymentStatusUnspecified
}

func (l *LoanSchedule) IsDue() bool {
	return l.PaymentStatus == PaymentStatusDue
}

func (l *LoanSchedule) IsPaid() bool {
	return l.PaymentStatus == PaymentStatusPaid
}

func (l *LoanSchedule) IsOverdue() bool {
	return l.PaymentStatus == PaymentStatusOverdue
}
