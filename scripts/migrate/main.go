package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sunil-bagde/ledger-app/database"
	"github.com/sunil-bagde/ledger-app/migrations"
)

func main() {

	if err := migrate(); err != nil {
		fmt.Println(err)
	}
}
func migrate() error {
	config := database.Config{
		User:         "ledgerapp",
		Password:     "ledgerapp",
		Host:         "localhost",
		Name:         "ledgerapp",
		MaxIdleConns: 0,
		MaxOpenConns: 0,
		DisableTLS:   true,
	}
	db, err := database.Open(config)
	if err != nil {
		return err
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := migrations.Migrate(ctx, db); err != nil {
		cRed := "\033[31m"
		return fmt.Errorf(cRed+"migrate: %w", err)
	}
	cGreen := "\033[32m"

	fmt.Println(cGreen + "migrations complete")
	return nil
}
