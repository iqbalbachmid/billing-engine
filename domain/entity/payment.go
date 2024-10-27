package entity

import "time"

const Status = "completed"

type Payment struct {
	PaymentID     int       `db:"payment_id"`
	LoanID        int       `db:"loan_id"`
	PaymentDate   time.Time `db:"payment_date"`
	AmountPaid    float64   `db:"amount_paid"`
	PaymentMethod string    `db:"payment_method"`
	Status        string    `db:"status"`
}
