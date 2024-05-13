package account

type Account struct {
	Id              string  `db:"id" json:"id"`
	UserId          string  `db:"user_id" json:"user_id"`
	Name            string  `db:"name" json:"name"`
	AvailableAmount float64 `db:"available_amount" json:"available_amount"`
	Type            string  `db:"type" json:"type"`
	Slug            string  `db:"slug" json:"slug"`
	TransactionType string  `db:"transaction_type" json:"transaction_type"`
	DateCreated     string  `db:"date_created" json:"date_created"`
	DateUpdated     string  `db:"date_updated" json:"date_updated"`
}

type NewAccount struct {
	UserId          string  `json:"user_id" validate:"required"`
	Name            string  `json:"name" validate:"required,min=2,max=100"`
	AvailableAmount float64 `json:"available_amount" validate:"required"`
	Type            string  `json:"type" validate:"required"`
	Slug            string  `json:"slug" validate:"required"`
	DateCreated     string  `db:"date_created" json:"date_created"`
}

type UpdateAccount struct {
	AvailableAmount float64 `db:"available_amount" json:"available_amount"`
	DateUpdated     string  `db:"date_updated" json:"date_updated"`
}
