package user

import (
	"time"

	"github.com/gofiber/fiber/v2"
	userCore "github.com/sunil-bagde/ledger-app/core/user"
	"github.com/sunil-bagde/ledger-app/store/user"
)

// Handlers manages the set of user enpoints.
type Handlers struct {
	User userCore.Core
}

func NewHandlers(user userCore.Core) *Handlers {
	return &Handlers{
		User: user,
	}
}

func (h *Handlers) CreateUserHandler(ctx *fiber.Ctx) error {
	var params user.NewUser
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	newUser := user.NewUser{
		FirstName:       params.FirstName,
		LastName:        params.LastName,
		Email:           params.Email,
		Password:        params.Password,
		Username:        params.Username,
		ConfirmPassword: params.ConfirmPassword,
		DateCreated:     time.Now().Format(time.RFC3339),
		DateUpdated:     time.Now().Format(time.RFC3339),
	}

	// add validation logic here
	h.User.Create(ctx, newUser, time.Now())
	return ctx.JSON(fiber.Map{"message": "User created successfully"})
}
