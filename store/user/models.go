package user

type User struct {
	Id           string `db:"id" json:"id"`
	FirstName    string `db:"first_name" json:"first_name"`
	LastName     string `db:"last_name" json:"last_name"`
	Username     string `db:"username" json:"username"`
	Email        string `db:"email" json:"email"`
	PasswordHash string `db:"password_hash" json:"-"`
	DateCreated  string `db:"date_created" json:"date_created"`
	DateUpdated  string `db:"date_updated" json:"date_updated"`
}

type NewUser struct {
	FirstName       string `json:"first_name" validate:"required,min=2,max=100"`
	LastName        string `json:"last_name" validate:"required,min=2,max=100"`
	Username        string `json:"username" validate:"required,min=2,max=100"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,max=100"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8,max=100"`
	DateCreated     string `db:"date_created" json:"date_created"`
	DateUpdated     string `db:"date_updated" json:"date_updated"`
}

type UpdateUser struct {
	FirstName       *string `json:"first_name" validate:"omitempty,min=2,max=100"`
	LastName        *string `json:"last_name" validate:"omitempty,min=2,max=100"`
	Username        *string `json:"username" validate:"omitempty,min=2,max=100"`
	Email           *string `json:"email" validate:"omitempty,email"`
	Password        *string `json:"password" validate:"omitempty,min=8,max=100"`
	ConfirmPassword *string `json:"confirm_password" validate:"omitempty,min=8,max=100"`
	DateCreated     string  `db:"date_created" json:"date_created"`
	DateUpdated     string  `db:"date_updated" json:"date_updated"`
}
