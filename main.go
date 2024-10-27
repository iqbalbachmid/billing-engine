package main

import (
	"github.com/iqbalbachmid/billing-engine/infrastructure/sql"
)

func main() {
	dbClient := sql.NewSQLite3Client()
	dbClient.CreateTables()
	defer dbClient.Close()

	//service := application.NewLoanService(nil, nil, nil, func() time.Time {
	//	return time.Now()
	//})
}
