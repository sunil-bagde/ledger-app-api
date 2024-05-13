package account

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sunil-bagde/ledger-app/database"
	"go.uber.org/zap"
)

type Store struct {
	db  *sqlx.DB
	log *zap.SugaredLogger
}

func NewStore(db *sqlx.DB, log *zap.SugaredLogger) *Store {
	return &Store{
		db:  db,
		log: log,
	}
}

func (s *Store) CreateAccount(ctx *fiber.Ctx, newAccount NewAccount, now time.Time) (Account, error) {
	// add validation logic here  61-62

	account := Account{
		UserId:          newAccount.UserId,
		Name:            newAccount.Name,
		AvailableAmount: newAccount.AvailableAmount,
		Type:            newAccount.Type,
		Slug:            newAccount.Slug,
		DateCreated:     now.Format(time.RFC3339),
	}
	const query = `INSERT INTO "public"."accounts" ("user_id", "name", "available_amount", "type", "slug", "date_created")
	VALUES (:user_id, :name, :available_amount, :type, :slug, :date_created);`

	if err := database.NamedExecContext(ctx.Context(), s.log, s.db, query, account); err != nil {
		return Account{}, fmt.Errorf("inserting user: %w", err)
	}

	return account, nil

}
