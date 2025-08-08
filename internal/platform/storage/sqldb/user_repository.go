package sqldb

import (
	"context"
	"database/sql"
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

// Save saves a user to the SQL database.
func (r *UserRepository) Save(ctx context.Context, user domain.User) error {
	userSQLStruct := sqlbuilder.NewStruct(new(sqlUser)).For(defaultFlavor)

	query, args := userSQLStruct.InsertInto(sqlUserTable, sqlUser{
		ID:    user.ID().String(),
		Name:  user.Name().String(),
		Email: user.Email().String(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("failed to save user: %v", err)
	}

	return nil
}

// Find retrieves a user by ID from the SQL database.
func (r *UserRepository) Find(ctx context.Context, id domain.UserID) (domain.User, error) {
	userSQLStruct := sqlbuilder.NewStruct(new(sqlUser)).For(defaultFlavor)

	query, _ := userSQLStruct.SelectFrom(sqlUserTable).Where("id = ?").Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	row := r.db.QueryRowContext(ctxTimeout, query, id.String())
	var userSQL sqlUser
	if err := row.Scan(userSQLStruct.Addr(&userSQL)...); err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, fmt.Errorf("user not found: %v", id)
		}
		return domain.User{}, fmt.Errorf("failed to find user: %v", err)
	}

	return domain.NewUser(userSQL.ID, userSQL.Name, userSQL.Email)
}
