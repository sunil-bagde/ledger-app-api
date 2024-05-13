package account

import (
	"time"

	"github.com/gofiber/fiber/v2"
	accountCore "github.com/sunil-bagde/ledger-app/core/account"
	accountStore "github.com/sunil-bagde/ledger-app/store/account"
)

// Handlers manages the set of user enpoints.
type Handlers struct {
	Account accountCore.Core
}

func NewHandlers(account accountCore.Core) *Handlers {
	return &Handlers{
		Account: account,
	}
}
func (h *Handlers) CreateAccountHandler(ctx *fiber.Ctx) error {
	var params accountStore.NewAccount
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	newAccountCreate := accountStore.NewAccount{
		UserId:          params.UserId,
		Name:            params.Name,
		AvailableAmount: params.AvailableAmount,
		Type:            params.Type,
		Slug:            params.Slug,
		DateCreated:     time.Now().Format(time.RFC3339),
	}

	// add validation logic here
	h.Account.Create(ctx, newAccountCreate, time.Now())
	return ctx.JSON(fiber.Map{"message": "Account created successfully"})
}
