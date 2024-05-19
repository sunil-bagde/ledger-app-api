package transaction

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/sunil-bagde/ledger-app/constant"
	tCore "github.com/sunil-bagde/ledger-app/core/transaction"
	tStore "github.com/sunil-bagde/ledger-app/store/transaction"
)

// Handlers manages the set of user enpoints.
type Handlers struct {
	Transaction tCore.Core
}

func NewHandlers(t tCore.Core) *Handlers {
	return &Handlers{
		Transaction: t,
	}
}

/*
 */
func (h *Handlers) CreateDepositHandler(ctx *fiber.Ctx) error {
	var params tStore.NewDeposit
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}

	if params.Type != constant.DEPOSIT {
		return ctx.Status(400).JSON(fiber.Map{"message": "Invalid 'type' request"})
	}
	newDeposite := tStore.NewDeposit{
		AccountId: params.AccountId,
		Amount:    params.Amount,
		Type:      params.Type,
		UserId:    params.UserId,
	}

	// add validation logic here
	_, err := h.Transaction.Deposit(ctx, newDeposite, time.Now())
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"message": "Error occured while depositing money to account"})
	}
	msg := fmt.Sprintf("%.2f has been deposited to %s account!", newDeposite.Amount, newDeposite.AccountId)
	return ctx.JSON(fiber.Map{"message": msg})
}

func (h *Handlers) CreateWithdrawHandler(ctx *fiber.Ctx) error {
	var params tStore.NewWithdraw
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}

	if params.Type != constant.WITHDRAW {
		return ctx.Status(400).JSON(fiber.Map{"message": "Invalid 'type' request"})
	}
	newWithdraw := tStore.NewWithdraw{
		AccountId: params.AccountId,
		Amount:    params.Amount,
		Type:      params.Type,
		UserId:    params.UserId,
	}

	// add validation logic here
	_, err := h.Transaction.Withdraw(ctx, newWithdraw, time.Now())
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Withdraw done successfully"})
}
func (h *Handlers) CreateUserTransferHandler(ctx *fiber.Ctx) error {

	var params tStore.NewTransfer
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}

	if params.Type != constant.TRANSFER {
		return ctx.Status(400).JSON(fiber.Map{"message": "Invalid 'type' request"})
	}

	h.Transaction.UserToAccountTransfer(ctx, params, time.Now())

	return ctx.JSON(fiber.Map{"message": "Transfer to account done successfully"})
}
func (h *Handlers) CreateAccountToAccountTransferHandler(ctx *fiber.Ctx) error {

	var params tStore.AccountToAccountTransfer
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}

	if params.Type != constant.TRANSFER {
		return ctx.Status(400).JSON(fiber.Map{"message": "Invalid 'type' request"})
	}

	h.Transaction.CreateAccountToAccountTransfer(ctx, params, time.Now())
	return ctx.JSON(fiber.Map{"message": "Transfer to account done successfully"})
}

func (h *Handlers) VerifyTransaction(ctx *fiber.Ctx) error {
	transactionId := ctx.Params("transactionId")

	_, err := h.Transaction.VerifyTransaction(ctx, transactionId, time.Now())
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Transaction verified successfully"})

}
