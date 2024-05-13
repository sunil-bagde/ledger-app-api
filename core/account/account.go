package account

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sunil-bagde/ledger-app/store/account"
	"go.uber.org/zap"
)

type Core struct {
	log     *zap.SugaredLogger
	account account.Store
}

func NewCore(log *zap.SugaredLogger, db *sqlx.DB) *Core {
	return &Core{
		log:     log,
		account: *account.NewStore(db, log),
	}
}
func (c Core) Create(ctx *fiber.Ctx, na account.NewAccount, now time.Time) (account.Account, error) {

	// PERFORM PRE BUSINESS OPERATIONS

	acc, err := c.account.CreateAccount(ctx, na, now)

	if err != nil {
		return account.Account{}, fmt.Errorf("create: %w", err)
	}

	// PERFORM POST BUSINESS OPERATIONS

	return acc, nil
}
