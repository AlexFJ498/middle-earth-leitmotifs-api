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

type GroupDB struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	ImageURL    string `db:"image_url"`
}

var sqlGroupTable = "groups"
var groupSQLStruct = sqlbuilder.NewStruct(new(GroupDB)).For(defaultFlavor)

// GroupRepository implements the GroupRepository interface for SQL.
type GroupRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewGroupRepository creates a new GroupRepository.
func NewGroupRepository(db *sql.DB, dbTimeout time.Duration) *GroupRepository {
	return &GroupRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

func groupToDTO(group domain.Group) GroupDB {
	return GroupDB{
		ID:          group.ID().String(),
		Name:        group.Name().String(),
		Description: group.Description().String(),
		ImageURL:    group.ImageURL().String(),
	}
}
func groupToDomain(dto GroupDB) (domain.Group, error) {
	return domain.NewGroupWithID(
		dto.ID,
		dto.Name,
		dto.Description,
		dto.ImageURL,
	)
}

func (r *GroupRepository) Save(ctx context.Context, group domain.Group) error {
	row := groupToDTO(group)
	query, args := groupSQLStruct.InsertInto(sqlGroupTable, row).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("failed to save group: %v", err)
	}

	return nil
}

func (r *GroupRepository) Find(ctx context.Context, id domain.GroupID) (domain.Group, error) {
	sb := groupSQLStruct.SelectFrom(sqlGroupTable)
	sb.Where(sb.Equal("id", id.String()))
	query, args := sb.Build()
	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	var groupDTO GroupDB
	err := r.db.QueryRowContext(ctxTimeout, query, args...).Scan(groupSQLStruct.Addr(&groupDTO)...)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Group{}, domain.ErrGroupNotFound
	}
	if err != nil {
		return domain.Group{}, fmt.Errorf("failed to find group: %v", err)
	}

	return groupToDomain(groupDTO)
}

func (r *GroupRepository) FindAll(ctx context.Context) ([]domain.Group, error) {
	sb := groupSQLStruct.SelectFrom(sqlGroupTable)
	sb.OrderBy("created_at ASC")
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctxTimeout, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to find groups: %v", err)
	}
	defer rows.Close()

	var groups []domain.Group
	for rows.Next() {
		var groupDTO GroupDB
		if err := rows.Scan(groupSQLStruct.Addr(&groupDTO)...); err != nil {
			return nil, fmt.Errorf("failed to scan group: %v", err)
		}
		group, err := groupToDomain(groupDTO)
		if err != nil {
			return nil, fmt.Errorf("failed to convert group: %v", err)
		}
		groups = append(groups, group)
	}

	return groups, nil
}

func (r *GroupRepository) Delete(ctx context.Context, id domain.GroupID) error {
	sb := groupSQLStruct.DeleteFrom(sqlGroupTable)
	sb.Where(sb.Equal("id", id.String()))
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	result, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete group: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return domain.ErrGroupNotFound
	}

	return nil
}

func (r *GroupRepository) Update(ctx context.Context, group domain.Group) error {
	row := groupToDTO(group)
	sb := groupSQLStruct.Update(sqlGroupTable, row)
	sb.Where(sb.Equal("id", row.ID))
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	result, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update group: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return domain.ErrGroupNotFound
	}

	return nil
}
