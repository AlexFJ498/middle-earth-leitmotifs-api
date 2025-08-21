package sqldb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"
	"github.com/huandu/go-sqlbuilder"
)

// UserRepository implements the UserRepository interface for SQL.
type UserRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewUserRepository creates a new UserRepository instance.
func NewUserRepository(db *sql.DB, dbTimeout time.Duration) *UserRepository {
	return &UserRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

type UserDB struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

var sqlUserTable = "users"
var userSQLStruct = sqlbuilder.NewStruct(new(UserDB)).For(defaultFlavor)

func toDTO(user domain.User) UserDB {
	return UserDB{
		ID:       user.ID().String(),
		Name:     user.Name().String(),
		Email:    user.Email().String(),
		Password: user.Password().String(),
	}
}

func toDomain(dto UserDB) (domain.User, error) {
	return domain.NewUserWithID(
		dto.ID,
		dto.Name,
		dto.Email,
		dto.Password,
	)
}

// Save saves a user to the SQL database.
func (r *UserRepository) Save(ctx context.Context, user domain.User) error {
	// First, check if the user already exists
	_, err := r.FindByEmail(ctx, user.Email())
	if err == nil {
		// User already exists, return an error
		return domain.ErrUserAlreadyExists
	}
	if !errors.Is(err, domain.ErrUserNotFound) {
		// Some other error occurred
		return fmt.Errorf("failed to check if user exists: %v", err)
	}

	row := toDTO(user)
	query, args := userSQLStruct.InsertInto(sqlUserTable, row).Build()
	fmt.Println(query, args)
	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err = r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("failed to save user: %v", err)
	}

	return nil
}

// Find retrieves a user by ID from the SQL database.
func (r *UserRepository) Find(ctx context.Context, id domain.UserID) (domain.User, error) {

	sb := userSQLStruct.SelectFrom(sqlUserTable)
	sb.Where(sb.Equal("id", id.String()))
	query, args := sb.Build()
	fmt.Println(query, args)
	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	var userDTO UserDB
	err := r.db.QueryRowContext(ctxTimeout, query, args...).Scan(userSQLStruct.Addr(&userDTO)...)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.User{}, domain.ErrUserNotFound
	}
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to find user: %v", err)
	}

	return toDomain(userDTO)
}

func (r *UserRepository) FindByEmail(ctx context.Context, email domain.UserEmail) (domain.User, error) {

	sb := userSQLStruct.SelectFrom(sqlUserTable)
	sb.Where(sb.Equal("email", email.String()))
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	var userDTO UserDB
	err := r.db.QueryRowContext(ctxTimeout, query, args...).Scan(userSQLStruct.Addr(&userDTO)...)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.User{}, domain.ErrUserNotFound
	}
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to find user: %v", err)
	}

	return toDomain(userDTO)
}
