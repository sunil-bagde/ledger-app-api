package user

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sunil-bagde/ledger-app/store/user"
	"go.uber.org/zap"
)

type Core struct {
	log  *zap.SugaredLogger
	user user.Store
}

func NewCore(log *zap.SugaredLogger, db *sqlx.DB) *Core {
	return &Core{
		log:  log,
		user: *user.NewStore(db, log),
	}
}

func (c Core) Create(ctx *fiber.Ctx, nu user.NewUser, now time.Time) (user.User, error) {

	// PERFORM PRE BUSINESS OPERATIONS

	usr, err := c.user.CreateUser(ctx, nu, now)
	fmt.Println(err)
	if err != nil {
		return user.User{}, fmt.Errorf("create: %w", err)
	}

	// PERFORM POST BUSINESS OPERATIONS

	return usr, nil
}
