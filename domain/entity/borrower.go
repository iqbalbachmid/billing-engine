package entity

import "time"

type Borrower struct {
	BorrowerID    int       `db:"borrower_id"`
	FirstName     string    `db:"first_name"`
	LastName      string    `db:"last_name"`
	Email         string    `db:"email"`
	Phone         string    `db:"phone"`
	Address       string    `db:"address"`
	DateOfBirth   time.Time `db:"date_of_birth"`
	AccountStatus string    `db:"account_status"`
}
