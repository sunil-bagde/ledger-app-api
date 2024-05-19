package transaction

import "database/sql"

/*
	from_account_id

to_account_id
account_id
user_id
to_user_id
amount
name
slug
TYPE
date_created
date_updated
*/
type Deposit struct {
	AccountId string  `db:"account_id" json:"account_id"`
	Amount    float64 `db:"amount" json:"amount"`
	Type      string  `db:"type" json:"type"`
	UserId    string  `db:"user_id" json:"user_id"`
}
type AccountActivity struct {
	AccountId       string  `db:"account_id" json:"account_id"`
	ToAccount       string  `db:"to_account_id" json:"to_account_id"`
	Amount          float64 `db:"amount" json:"amount"`
	Type            string  `db:"type" json:"type"`
	TransactionType string  `db:"transaction_type" json:"transaction_type"`
	UserId          string  `db:"user_id" json:"user_id"`
	Hash            string  `db:"hash" json:"hash"`
	PreviousHash    string  `db:"previous_hash" json:"previous_hash"`
	DateCreated     string  `db:"date_created" json:"date_created"`
	DateUpdated     string  `db:"date_updated" json:"date_updated"`
}

type NewDeposit struct {
	AccountId string  `json:"account_id"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"`
	UserId    string  `json:"user_id"`
}

type NewWithdraw struct {
	AccountId string  `json:"account_id"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"`
	UserId    string  `json:"user_id"`
}

type NewTransfer struct {
	UserId          string  `db:"user_id" json:"user_id"`
	AccountId       string  `db:"account_id" json:"account_id"`
	Amount          float64 `db:"amount" json:"amount"`
	Type            string  `db:"type" json:"type"`
	TransactionType string  `db:"transaction_type" json:"transaction_type"`
	ToUser          string  `db:"to_user_id" json:"to_user"`
	ToAccount       string  `db:"to_account_id" json:"to_account"`
	Hash            string  `db:"hash" json:"hash"`
	PreviousHash    string  `db:"previous_hash" json:"previous_hash"`
	DateCreated     *string `db:"date_created" json:"date_created"`
	DateUpdated     *string `db:"date_updated" json:"date_updated"`
}

type NewAccountToTransfer struct {
	FromAccountId string  `db:"account_id" json:"from_account_id"`
	ToAccountId   string  `db:"to_account_id" json:"to_account_id"`
	Amount        float64 `db:"amount" json:"amount"`
	Type          string  `db:"type" json:"type"`
	DateCreated   *string `db:"date_created" json:"date_created"`
	DateUpdated   *string `db:"date_updated" json:"date_updated"`
}
type FromAccountToTransfer struct {
	UserId          string  `db:"user_id" json:"user_id"`
	AccountId       string  `db:"account_id" json:"account_id"`
	Amount          float64 `db:"amount" json:"amount"`
	Type            string  `db:"type" json:"type"`
	TransactionType string  `db:"transaction_type" json:"transaction_type"`
	Hash            string  `db:"hash" json:"hash"`
	PreviousHash    string  `db:"previous_hash" json:"previous_hash"`
	DateCreated     *string `db:"date_created" json:"date_created"`
	DateUpdated     *string `db:"date_updated" json:"date_updated"`
}

type ToAccountToTransfer struct {
	Amount          float64 `db:"amount" json:"amount"`
	Type            string  `db:"type" json:"type"`
	TransactionType string  `db:"transaction_type" json:"transaction_type"`
	ToUser          string  `db:"to_user" json:"to_user"`
	ToAccount       string  `db:"to_account" json:"to_account"`
	Hash            string  `db:"hash" json:"hash"`
	PreviousHash    string  `db:"previous_hash" json:"previous_hash"`
	DateCreated     *string `db:"date_created" json:"date_created"`
	DateUpdated     *string `db:"date_updated" json:"date_updated"`
}

type AccountToAccountTransfer struct {
	Amount          float64 `db:"amount" json:"amount"`
	Type            string  `db:"type" json:"type"`
	FromAccount     string  `db:"account_id" json:"from_account_id"`
	TransactionType string  `db:"transaction_type" json:"transaction_type"`
	ToAccount       string  `db:"to_account" json:"to_account_id"`
	DateCreated     *string `db:"date_created" json:"date_created"`
	DateUpdated     *string `db:"date_updated" json:"date_updated"`
}

type CreateHash struct {
	AccountId       sql.NullString `db:"account_id" json:"account_id"`
	ToAccount       sql.NullString `db:"to_account_id" json:"to_account_id"`
	Amount          float64        `db:"amount" json:"amount"`
	Type            string         `db:"type" json:"type"`
	TransactionType string         `db:"transaction_type" json:"transaction_type"`
	UserId          sql.NullString `db:"user_id" json:"user_id"`
	ToUser          sql.NullString `db:"to_user_id" json:"to_user_id"`
	Hash            string         `db:"hash" json:"hash"`
	PreviousHash    sql.NullString `db:"previous_hash" json:"previous_hash"`
}
