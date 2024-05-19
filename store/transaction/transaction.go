package transaction

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sunil-bagde/ledger-app/constant"
	"github.com/sunil-bagde/ledger-app/database"
	"github.com/sunil-bagde/ledger-app/store/account"
	"go.uber.org/zap"
)

type Store struct {
	db  *sqlx.DB
	log *zap.SugaredLogger
}

const (
	INSERT_TRANSACTION_QUERY = `
	INSERT INTO "public"."transactions" ("account_id", "user_id", "amount", "type", "date_created", "date_updated")
	VALUES (:account_id, :user_id, :amount, :type, :date_created, :date_updated);
	`
)

func NewStore(db *sqlx.DB, log *zap.SugaredLogger) *Store {
	return &Store{
		db:  db,
		log: log,
	}
}

func (s *Store) Deposit(ctx *fiber.Ctx, newAcc NewDeposit, now time.Time) (AccountActivity, error) {
	// add validation logic here  61-62
	acc, err := s.QueryByID(ctx, newAcc.AccountId)
	if err != nil {
		return AccountActivity{}, fmt.Errorf("querying account: %w", err)
	}
	newDepositeCreate := AccountActivity{
		AccountId:       newAcc.AccountId,
		Amount:          newAcc.Amount,
		Type:            string(newAcc.Type),
		UserId:          newAcc.UserId,
		TransactionType: constant.CREDIT,
		DateCreated:     now.Format(time.RFC3339),
		DateUpdated:     now.Format(time.RFC3339),
	}

	s.Update(ctx, newAcc.AccountId, account.UpdateAccount{
		AvailableAmount: acc.AvailableAmount + newAcc.Amount,
		DateUpdated:     now.Format(time.RFC3339),
	}, now) // Add the missing argument "now" to the function call.

	const query = `
	INSERT INTO "public"."transactions" ("account_id", "user_id", "amount", "type", "transaction_type", "date_created", "date_updated")
	VALUES (:account_id, :user_id, :amount, :type, :transaction_type, :date_created, :date_updated);
	`
	if err := database.NamedExecContext(ctx.Context(), s.log, s.db, query, &newDepositeCreate); err != nil {
		return AccountActivity{}, fmt.Errorf("inserting deposite: %w", err)
	}

	return newDepositeCreate, nil

}
func (s *Store) Withdraw(ctx *fiber.Ctx, newAcc NewWithdraw, now time.Time) (AccountActivity, error) {
	// add validation logic here  61-62

	acc, err := s.QueryByID(ctx, newAcc.AccountId)
	if err != nil {
		return AccountActivity{}, fmt.Errorf("querying account: %w", err)
	}
	if acc.AvailableAmount < newAcc.Amount {
		return AccountActivity{}, errors.New("low balance in account")
	}
	newDepositeCreate := AccountActivity{
		AccountId:       newAcc.AccountId,
		Amount:          newAcc.Amount,
		Type:            string(newAcc.Type),
		UserId:          newAcc.UserId,
		TransactionType: constant.DEBIT,
		DateCreated:     now.Format(time.RFC3339),
		DateUpdated:     now.Format(time.RFC3339),
	}

	s.Update(ctx, newAcc.AccountId, account.UpdateAccount{
		AvailableAmount: acc.AvailableAmount - newAcc.Amount,
		DateUpdated:     now.Format(time.RFC3339),
	}, now) // Add the missing argument "now" to the function call.

	const query = `
	INSERT INTO "public"."transactions" ("account_id", "user_id", "amount", "type", "transaction_type", "date_created", "date_updated")
	VALUES (:account_id, :user_id, :amount, :type, :transaction_type, :date_created, :date_updated);
	`
	if err := database.NamedExecContext(ctx.Context(), s.log, s.db, query, newDepositeCreate); err != nil {
		return AccountActivity{}, fmt.Errorf("inserting deposite: %w", err)
	}

	return newDepositeCreate, nil

}

// GenerateID generate a unique id for entities.
func GenerateID() string {
	return uuid.NewString()
}
func (s Store) Update(ctx *fiber.Ctx, accountId string, up account.UpdateAccount, now time.Time) error {

	_, err := s.QueryByID(ctx, accountId)
	if err != nil {
		return fmt.Errorf("updating account balance accountId[%s]: %w", accountId, err)
	}

	data := struct {
		AvailableAmount float64 `db:"available_amount" json:"available_amount"`
		DateUpdated     string  `db:"date_updated" json:"date_updated"`
		Id              string  `db:"id" json:"id"`
	}{
		AvailableAmount: up.AvailableAmount,
		DateUpdated:     now.Format(time.RFC3339),
		Id:              accountId,
	}

	const q = `
	UPDATE
		accounts
	SET
		"available_amount" = :available_amount,
		"date_updated" = :date_updated
	WHERE
	id = :id`

	if err := database.NamedExecContext(ctx.Context(), s.log, s.db, q, data); err != nil {
		return fmt.Errorf("updating product productID[%s]: %w", accountId, err)
	}

	return nil
}
func (s *Store) QueryLastTrsactionByUserId(ctx *fiber.Ctx, userId string) (NewTransfer, error) {

	data := struct {
		UserId string `db:"user_id"`
	}{
		UserId: userId,
	}
	const q = `SELECT * FROM transactions WHERE id=:`

	var prd NewTransfer

	if err := database.NamedQueryStruct(ctx.Context(), s.log, s.db, q, data, &prd); err != nil {
		if err == database.ErrNotFound {
			return NewTransfer{}, database.ErrNotFound
		}

		return NewTransfer{}, fmt.Errorf("selecting product accountId[%q]: %w", userId, err)
	}

	return prd, nil
}

// ctx *fiber.Ctx, newAcc NewDeposit, now time.Time
func (s *Store) QueryByID(ctx *fiber.Ctx, accountId string) (account.Account, error) {

	data := struct {
		AccountId string `db:"account_id"`
	}{
		AccountId: accountId,
	}
	const q = `SELECT
						*
					FROM
						"public"."accounts"
					WHERE
						"id" = :account_id
					LIMIT 1
		`

	var prd account.Account

	if err := database.NamedQueryStruct(ctx.Context(), s.log, s.db, q, data, &prd); err != nil {
		if err == database.ErrNotFound {
			return account.Account{}, database.ErrNotFound
		}

		return account.Account{}, fmt.Errorf("selecting product accountId[%q]: %w", accountId, err)
	}

	return prd, nil
}

func (s *Store) AccountToUserTransfer(ctx *fiber.Ctx, nt NewTransfer, now time.Time) (AccountActivity, error) {
	fromAcc, fromErr := s.QueryByID(ctx, nt.AccountId)
	if fromErr != nil {
		return AccountActivity{}, fmt.Errorf("querying from account: %w", fromErr)
	}
	toAcc, toErr := s.QueryByID(ctx, nt.ToAccount)
	if toErr != nil {
		return AccountActivity{}, fmt.Errorf("querying to account: %w", toErr)
	}
	fromAccBalance := fromAcc.AvailableAmount - nt.Amount
	// from update
	s.Update(ctx, nt.AccountId, account.UpdateAccount{
		AvailableAmount: fromAccBalance,
		DateUpdated:     now.Format(time.RFC3339),
	}, now)
	// to update
	toAccBalance := toAcc.AvailableAmount + nt.Amount
	s.Update(ctx, nt.ToAccount, account.UpdateAccount{
		AvailableAmount: toAccBalance,
		DateUpdated:     now.Format(time.RFC3339),
	}, now)
	date := now.Format(time.RFC3339)

	frmAccount := NewTransfer{
		AccountId:       nt.AccountId,
		UserId:          nt.UserId,
		Amount:          nt.Amount,
		Type:            constant.TRANSFER,
		ToAccount:       nt.ToAccount,
		TransactionType: constant.DEBIT,
		DateCreated:     &date,
		DateUpdated:     &date,
	}

	if err := database.NamedExecContext(ctx.Context(), s.log, s.db, `INSERT INTO "public"."transactions" ("account_id", "user_id", "amount", "type", "transaction_type", "date_created", "date_updated")
	VALUES (:account_id, :user_id, :amount, :type, :transaction_type, :date_created, :date_updated)`, frmAccount); err != nil {
		return AccountActivity{}, fmt.Errorf("inserting transfer account to user: %w", err)
	}
	toAccount := NewTransfer{
		AccountId:       nt.ToAccount,
		Amount:          nt.Amount,
		Type:            nt.Type,
		TransactionType: constant.CREDIT,
		UserId:          nt.UserId,
		DateCreated:     &date,
		DateUpdated:     &date,
	}
	if err := database.NamedExecContext(ctx.Context(), s.log, s.db, INSERT_TRANSACTION_QUERY, toAccount); err != nil {
		return AccountActivity{}, fmt.Errorf("inserting transfer account to user: %w", err)
	}

	return AccountActivity{}, nil
}

func createSenderTransactionHash(t NewTransfer) string {
	data := fmt.Sprintf("%.2f%s%s%s%s", t.Amount, t.UserId, t.AccountId, t.Type, t.PreviousHash)
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
func createReceiverTransactionHash(t NewTransfer) string {

	data := fmt.Sprintf("%.2f%s%s%s%s", t.Amount, t.ToUser, t.ToAccount, t.Type, t.PreviousHash)

	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// HERE
func (s *Store) UserToAccountTransfer(ctx *fiber.Ctx, nt NewTransfer, now time.Time) (AccountActivity, error) {
	tx, err := s.db.Beginx()
	if err != nil {

		return AccountActivity{}, err
	}
	fromAcc, fromErr := s.QueryByID(ctx, nt.AccountId)

	if fromErr != nil {

		return AccountActivity{}, fmt.Errorf("querying from account: %w", fromErr)
	}
	toAcc, toErr := s.QueryByID(ctx, nt.ToAccount)
	if toErr != nil {

		return AccountActivity{}, fmt.Errorf("querying to account: %w", toErr)
	}

	fromAccBalance := fromAcc.AvailableAmount - nt.Amount
	// from update
	s.Update(ctx, nt.AccountId, account.UpdateAccount{
		AvailableAmount: fromAccBalance,
		DateUpdated:     now.Format(time.RFC3339),
	}, now)
	// to update
	toAccBalance := toAcc.AvailableAmount + nt.Amount
	s.Update(ctx, nt.ToAccount, account.UpdateAccount{
		AvailableAmount: toAccBalance,
		DateUpdated:     now.Format(time.RFC3339),
	}, now)
	date := now.Format(time.RFC3339)
	type Transaction1 struct {
		PreviousHash string `db:"previous_hash"`
		Hash         string `db:"hash"`
	}
	var lastTransaction Transaction1
	errLastTx := tx.Get(&lastTransaction, "SELECT hash, previous_hash FROM transactions WHERE date_created = (SELECT MAX(date_created) FROM transactions) FOR UPDATE")

	if errLastTx != nil && errLastTx != sql.ErrNoRows {
		_ = tx.Rollback()

		return AccountActivity{}, errLastTx
	}
	if err != sql.ErrNoRows {

		lastTransaction.PreviousHash = lastTransaction.Hash
	} else {
		lastTransaction.PreviousHash = ""
	}
	nt.PreviousHash = lastTransaction.PreviousHash
	fromHash := createSenderTransactionHash(nt)

	frmAccount := FromAccountToTransfer{
		AccountId:       nt.AccountId,
		Amount:          nt.Amount,
		Type:            nt.Type,
		DateCreated:     &date,
		DateUpdated:     &date,
		Hash:            fromHash,
		PreviousHash:    lastTransaction.Hash,
		TransactionType: constant.DEBIT,
		UserId:          nt.UserId,
	}
	if err := database.NamedExecContext(ctx.Context(), s.log, s.db, `INSERT INTO "public"."transactions" ("account_id", "user_id", "amount", "type", "transaction_type", "hash",
	"previous_hash","date_created", "date_updated")
	VALUES (:account_id, :user_id, :amount, :type, :transaction_type,:hash, :previous_hash, :date_created, :date_updated)`, frmAccount); err != nil {

		return AccountActivity{}, fmt.Errorf("inserting transfer account to user: %w", err)
	}
	nt.PreviousHash = fromHash
	toHash := createReceiverTransactionHash(nt)
	toAccount := ToAccountToTransfer{
		ToUser:          nt.ToUser,
		ToAccount:       nt.ToAccount,
		Amount:          nt.Amount,
		Type:            nt.Type,
		TransactionType: constant.CREDIT,
		Hash:            toHash,
		PreviousHash:    fromHash,
		DateCreated:     &date,
		DateUpdated:     &date,
	}

	toInsertQuery := `INSERT INTO "public"."transactions" ("to_user_id", "to_account_id", "amount", "type", "transaction_type", "hash",
	"previous_hash", "date_created", "date_updated")
	VALUES (:to_user, :to_account, :amount, :type, :transaction_type, :hash, :previous_hash,:date_created, :date_updated)`

	if err := database.NamedExecContext(ctx.Context(), s.log, s.db, toInsertQuery, toAccount); err != nil {

		return AccountActivity{}, err
	}
	// Commit the transaction
	err = tx.Commit()
	if err != nil {

		return AccountActivity{}, err
	}
	return AccountActivity{}, nil
}

func (s *Store) AccountToAccountTransfer(ctx *fiber.Ctx, nt AccountToAccountTransfer, now time.Time) (AccountActivity, error) {
	fromAcc, fromErr := s.QueryByID(ctx, nt.FromAccount)

	if fromErr != nil {
		return AccountActivity{}, fmt.Errorf("querying from account: %w", fromErr)
	}
	toAcc, toErr := s.QueryByID(ctx, nt.ToAccount)
	if toErr != nil {
		return AccountActivity{}, fmt.Errorf("querying to account: %w", toErr)
	}
	fromAccBalance := fromAcc.AvailableAmount - nt.Amount
	// from update
	s.Update(ctx, nt.FromAccount, account.UpdateAccount{
		AvailableAmount: fromAccBalance,
		DateUpdated:     now.Format(time.RFC3339),
	}, now)
	// to update
	toAccBalance := toAcc.AvailableAmount + nt.Amount
	s.Update(ctx, nt.ToAccount, account.UpdateAccount{
		AvailableAmount: toAccBalance,
		DateUpdated:     now.Format(time.RFC3339),
	}, now)
	date := now.Format(time.RFC3339)

	frmAccount := AccountToAccountTransfer{
		FromAccount:     nt.FromAccount,
		Amount:          nt.Amount,
		Type:            nt.Type,
		DateCreated:     &date,
		DateUpdated:     &date,
		TransactionType: constant.DEBIT,
	}
	FROM_ACCOUNT_QUERY := `INSERT INTO "public"."transactions" ("account_id", "amount", "type", "transaction_type","date_created", "date_updated")
	VALUES (:account_id, :amount, :type, :transaction_type, :date_created, :date_updated)`
	if err := database.NamedExecContext(ctx.Context(), s.log, s.db, FROM_ACCOUNT_QUERY, frmAccount); err != nil {
		fmt.Print(err)
		return AccountActivity{}, fmt.Errorf("inserting transfer account to user: %w", err)
	}

	toAccount := AccountToAccountTransfer{
		ToAccount:       nt.ToAccount,
		Amount:          nt.Amount,
		Type:            nt.Type,
		TransactionType: constant.CREDIT,
		DateCreated:     &date,
		DateUpdated:     &date,
	}

	toInsertQuery := `INSERT INTO "public"."transactions" ( "to_account_id", "amount", "type", "transaction_type", "date_created", "date_updated")
	   	VALUES ( :to_account, :amount, :type, :transaction_type, :date_created, :date_updated)`

	if err := database.NamedExecContext(ctx.Context(), s.log, s.db, toInsertQuery, toAccount); err != nil {
		return AccountActivity{}, fmt.Errorf("inserting transfer account to user: %w", err)
	}

	return AccountActivity{}, nil
}
func (s *Store) VerifyTransaction(ctx *fiber.Ctx, transactionId string, now time.Time) (AccountActivity, error) {
	data := struct {
		TransactionId string `db:"transaction_id"`
	}{
		TransactionId: transactionId,
	}
	const q = `SELECT to_account_id, to_user_id, account_id, user_id, amount, type, hash, previous_hash FROM transactions WHERE id = :transaction_id`

	var createHash CreateHash

	if err := database.NamedQueryStruct(ctx.Context(), s.log, s.db, q, data, &createHash); err != nil {
		if err == database.ErrNotFound {
			return AccountActivity{}, database.ErrNotFound
		}
		return AccountActivity{}, fmt.Errorf("selecting product accountId[%q]: %w", transactionId, err)
	}

	var prevHash string
	var acId string
	var uId string
	var toAc string
	var toUser string

	if createHash.PreviousHash.Valid {
		prevHash = createHash.PreviousHash.String
	}
	if createHash.AccountId.Valid {
		acId = createHash.AccountId.String
	}
	if createHash.ToAccount.Valid {
		toAc = createHash.ToAccount.String
	}
	if createHash.UserId.Valid {
		uId = createHash.UserId.String
	}
	if createHash.ToUser.Valid {
		toUser = createHash.ToUser.String
	}

	var t = NewTransfer{
		Amount:       createHash.Amount,
		UserId:       uId,
		AccountId:    acId,
		Type:         createHash.Type,
		ToAccount:    toAc,
		PreviousHash: prevHash,
	}
	t2 := NewTransfer{
		Amount:       createHash.Amount,
		ToUser:       toUser,
		Type:         createHash.Type,
		ToAccount:    toAc,
		PreviousHash: prevHash,
	}

	var isValid bool = false

	fmt.Println(createHash.Hash, " ", createReceiverTransactionHash(t2))

	if createHash.Hash == createSenderTransactionHash(t) {
		isValid = true
	}
	fmt.Println(isValid)
	if createHash.Hash == createReceiverTransactionHash(t2) {
		isValid = true
	}
	fmt.Println(isValid)
	if !isValid {
		return AccountActivity{}, errors.New("transaction is not valid")
	}
	return AccountActivity{
		AccountId:       acId,
		Amount:          createHash.Amount,
		Type:            createHash.Type,
		UserId:          toAc,
		TransactionType: constant.CREDIT,
	}, nil
}
