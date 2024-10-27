package sql

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DbClient struct {
	DB *sql.DB
}

func NewSQLite3Client() *DbClient {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	return &DbClient{
		DB: db,
	}
}

func (c *DbClient) CreateTables() {
	createTableSQL := `
    CREATE TABLE borrowers (
	  borrower_id INTEGER PRIMARY KEY AUTOINCREMENT,
	  first_name VARCHAR(255),
	  last_name VARCHAR(255),
	  email VARCHAR(255),
	  phone VARCHAR(20),
	  address TEXT,
	  date_of_birth DATE,
	  account_status TEXT CHECK(account_status IN ('active', 'delinquent', 'closed'))
	);
	CREATE TABLE loans (
	  loan_id INTEGER PRIMARY KEY AUTOINCREMENT,
	  borrower_id INTEGER,
	  loan_amount DECIMAL(15, 2),
	  interest_rate DECIMAL(5, 2),
	  loan_start_date DATE,
	  loan_end_date DATE,
	  loan_status TEXT CHECK(loan_status IN ('active', 'paid'))
	);
	CREATE TABLE loan_schedule (
	  schedule_id INTEGER PRIMARY KEY AUTOINCREMENT,
	  loan_id INTEGER,
	  due_date DATE,
	  principal_amount DECIMAL(15, 2),
	  interest_amount DECIMAL(15, 2),
	  total_due DECIMAL(15, 2),
	  payment_status TEXT CHECK(payment_status IN ('unspecified', 'due', 'paid', 'overdue'))
	);
	CREATE TABLE payments (
	  payment_id INTEGER PRIMARY KEY AUTOINCREMENT,
	  loan_id INTEGER,
	  payment_date DATE,
	  amount_paid DECIMAL(15, 2),
	  payment_method TEXT CHECK(payment_method IN ('bank_transfer')),
	  status TEXT CHECK(status IN ('completed'))
	);`
	if _, err := c.DB.Exec(createTableSQL); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

func (c *DbClient) Close() {
	if err := c.DB.Close(); err != nil {
		log.Fatalf("Failed to close db: %v", err)
	}
}
