package entity

import "time"

type Loan struct {
	LoanID        int       `db:"loan_id"`
	BorrowerID    int       `db:"borrower_id"`
	LoanAmount    float64   `db:"loan_amount"`
	InterestRate  float64   `db:"interest_rate"`
	LoanStartDate time.Time `db:"loan_start_date"`
	LoanEndDate   time.Time `db:"loan_end_date"`
	LoanStatus    string    `db:"loan_status"`
}
