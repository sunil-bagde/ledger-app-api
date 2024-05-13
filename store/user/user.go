package user

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sunil-bagde/ledger-app/database"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
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

func (s *Store) CreateUser(ctx *fiber.Ctx, newUser NewUser, now time.Time) (User, error) {
	// add validation logic here  61-62

	hash, err := hashPassword(newUser.Password)
	if err != nil {
		return User{}, err
	}
	user := User{
		Id:           GenerateID(),
		FirstName:    newUser.FirstName,
		LastName:     newUser.LastName,
		Username:     newUser.Username,
		Email:        newUser.Email,
		PasswordHash: hash,
		DateCreated:  now.Format(time.RFC3339), // Convert now to string using a specific format
		DateUpdated:  now.Format(time.RFC3339), // Convert now to string using a specific format
	}
	const query = `INSERT INTO users (id, first_name, last_name, username, email, password_hash, date_created, date_updated)
	VALUES (:id, :first_name, :last_name, :username, :email, :password_hash, :date_created, :date_updated)`

	if err := database.NamedExecContext(ctx.Context(), s.log, s.db, query, user); err != nil {
		return User{}, fmt.Errorf("inserting user: %w", err)
	}

	return user, nil

}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// GenerateID generate a unique id for entities.
func GenerateID() string {
	return uuid.NewString()
}
