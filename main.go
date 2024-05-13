package main

import (
	"github.com/ardanlabs/conf/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	accountCore "github.com/sunil-bagde/ledger-app/core/account"
	transactionCore "github.com/sunil-bagde/ledger-app/core/transaction"
	userCore "github.com/sunil-bagde/ledger-app/core/user"
	"github.com/sunil-bagde/ledger-app/database"
	"github.com/sunil-bagde/ledger-app/handler"
	"github.com/sunil-bagde/ledger-app/handler/account"
	transactionHandler "github.com/sunil-bagde/ledger-app/handler/transaction"
	userHandler "github.com/sunil-bagde/ledger-app/handler/user"
	"go.uber.org/zap"
)

func main() {
	/* migrate()
	return */
	cfg := struct {
		conf.Version

		DB struct {
			User         string `conf:"default:ledgerapp"`
			Password     string `conf:"default:ledgerapp,mask"`
			Host         string `conf:"default:localhost"`
			Name         string `conf:"default:ledgerapp"`
			MaxIdleConns int    `conf:"default:0"`
			MaxOpenConns int    `conf:"default:0"`
			DisableTLS   bool   `conf:"default:true"`
		}
	}{
		Version: conf.Version{
			Build: "v1.0.0",
			Desc:  "Ledger app",
		},
	}
	const prefix = "LEDGER"
	conf.Parse(prefix, &cfg)

	// Database Sup	port

	// Create connectivity to the database.

	log.Infow("startup", "status", "initializing database support", "host", cfg.DB.Host)
	configDb := database.Config{
		User:         cfg.DB.User,
		Password:     cfg.DB.Password,
		Host:         cfg.DB.Host,
		Name:         cfg.DB.Name,
		MaxIdleConns: cfg.DB.MaxIdleConns,
		MaxOpenConns: cfg.DB.MaxOpenConns,
		DisableTLS:   cfg.DB.DisableTLS,
	}
	db, err := database.Open(configDb)
	if err != nil {
		log.Fatalw("startup", "status", "failed to connect to database", "host", cfg.DB.Host, "error", err)
		return
	}
	db.Ping()
	defer func() {
		log.Infow("shutdown", "status", "stopping database support", "host", cfg.DB.Host)
		db.Close()
	}()
	app := fiber.New()

	// User Handler
	ugh := userHandler.NewHandlers(*userCore.NewCore(zap.S(), db))
	// Account Handler
	accCore := account.NewHandlers(*accountCore.NewCore(zap.S(), db))

	// Deposite Handler
	createTransactionHandler := transactionHandler.NewHandlers(*transactionCore.NewCore(zap.S(), db))
	tHandler := transactionHandler.NewHandlers(createTransactionHandler.Transaction)
	handler := handler.NewHomeHandler(db)
	userHandler := userHandler.NewHandlers(ugh.User)
	accountHandler := account.NewHandlers(accCore.Account)
	// API
	app.Get("/", handler.HomeHandler)
	app.Post("/api/users", userHandler.CreateUserHandler)
	app.Post("/api/accounts", accountHandler.CreateAccountHandler)
	app.Post("/api/deposit", tHandler.CreateDepositHandler)
	app.Post("/api/withdraw", tHandler.CreateWithdrawHandler)
	app.Post("/api/transfer/account-to-user", tHandler.CreateUserTransferHandler)
	app.Post("/api/transfer/account-to-account", tHandler.CreateAccountToAccountTransferHandler)

	// Start the server on port 3000
	app.Listen(":3000") // Remove the second argument from the app.Listen function call

}
