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

type UserDB struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	IsAdmin  bool   `db:"is_admin"`
}

var sqlUserTable = "users"
var userSQLStruct = sqlbuilder.NewStruct(new(UserDB)).For(defaultFlavor)

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

func userToDTO(user domain.User) UserDB {
	return UserDB{
		ID:       user.ID().String(),
		Name:     user.Name().String(),
		Email:    user.Email().String(),
		Password: user.Password().String(),
		IsAdmin:  user.IsAdmin().Bool(),
	}
}

func userToDomain(dto UserDB) (domain.User, error) {
	return domain.NewUserWithID(
		dto.ID,
		dto.Name,
		dto.Email,
		dto.Password,
		dto.IsAdmin,
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

	row := userToDTO(user)
	query, args := userSQLStruct.InsertInto(sqlUserTable, row).Build()

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

	return userToDomain(userDTO)
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

	return userToDomain(userDTO)
}

func (r *UserRepository) FindAll(ctx context.Context) ([]domain.User, error) {
	sb := userSQLStruct.SelectFrom(sqlUserTable)
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctxTimeout, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to find users: %v", err)
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var userDTO UserDB
		if err := rows.Scan(userSQLStruct.Addr(&userDTO)...); err != nil {
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		user, err := userToDomain(userDTO)
		if err != nil {
			return nil, fmt.Errorf("failed to convert user: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}
