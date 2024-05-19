package transaction

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sunil-bagde/ledger-app/store/transaction"

	"go.uber.org/zap"
)

type Core struct {
	log         *zap.SugaredLogger
	transaction transaction.Store
}

func NewCore(log *zap.SugaredLogger, db *sqlx.DB) *Core {
	return &Core{
		log:         log,
		transaction: *transaction.NewStore(db, log),
	}
}

func (c Core) Deposit(ctx *fiber.Ctx, nt transaction.NewDeposit, now time.Time) (transaction.AccountActivity, error) {

	// PERFORM PRE BUSINESS OPERATIONS

	usr, err := c.transaction.Deposit(ctx, nt, now)
	fmt.Println(err)
	if err != nil {
		return transaction.AccountActivity{}, fmt.Errorf("create: %w", err)
	}

	// PERFORM POST BUSINESS OPERATIONS

	return usr, nil
}

func (c Core) Withdraw(ctx *fiber.Ctx, nt transaction.NewWithdraw, now time.Time) (transaction.AccountActivity, error) {

	// PERFORM PRE BUSINESS OPERATIONS

	usr, err := c.transaction.Withdraw(ctx, nt, now)

	if err != nil {
		return transaction.AccountActivity{}, fmt.Errorf("create: %w", err)
	}

	// PERFORM POST BUSINESS OPERATIONS

	return usr, nil
}

/*
	 func (c Core) AccountToAccountTransfer(ctx *fiber.Ctx, nt transaction.NewAccountToTransfer, now time.Time) (transaction.AccountActivity, error) {

		// PERFORM PRE BUSINESS OPERATIONS

		usr, err := c.transaction.AccountToAccountTransfer(ctx, nt, now)

		if err != nil {
			return transaction.AccountActivity{}, fmt.Errorf("create: %w", err)
		}

		// PERFORM POST BUSINESS OPERATIONS

		return usr, nil
	}
*/
func (c Core) UserToAccountTransfer(ctx *fiber.Ctx, nt transaction.NewTransfer, now time.Time) (transaction.AccountActivity, error) {

	// PERFORM PRE BUSINESS OPERATIONS

	usr, err := c.transaction.UserToAccountTransfer(ctx, nt, now)

	if err != nil {
		return transaction.AccountActivity{}, fmt.Errorf("create: %w", err)
	}

	// PERFORM POST BUSINESS OPERATIONS

	return usr, nil
}
func (c Core) CreateAccountToAccountTransfer(ctx *fiber.Ctx, nt transaction.AccountToAccountTransfer, now time.Time) (transaction.AccountActivity, error) {

	// PERFORM PRE BUSINESS OPERATIONS

	usr, err := c.transaction.AccountToAccountTransfer(ctx, nt, now)

	if err != nil {
		return transaction.AccountActivity{}, fmt.Errorf("create: %w", err)
	}

	// PERFORM POST BUSINESS OPERATIONS

	return usr, nil
}

func (c Core) VerifyTransaction(ctx *fiber.Ctx, transactionId string, now time.Time) (transaction.AccountActivity, error) {

	// PERFORM PRE BUSINESS OPERATIONS

	usr, err := c.transaction.VerifyTransaction(ctx, transactionId, now)

	if err != nil {
		return transaction.AccountActivity{}, fmt.Errorf("create: %w", err)
	}

	// PERFORM POST BUSINESS OPERATIONS

	return usr, nil
}
