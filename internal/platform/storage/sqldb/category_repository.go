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

type CategoryDB struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

var sqlCategoryTable = "categories"
var categorySQLStruct = sqlbuilder.NewStruct(new(CategoryDB)).For(defaultFlavor)

// CategoryRepository implements the CategoryRepository interface for SQL.
type CategoryRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewCategoryRepository creates a new CategoryRepository.
func NewCategoryRepository(db *sql.DB, dbTimeout time.Duration) *CategoryRepository {
	return &CategoryRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

func categoryToDTO(category domain.Category) CategoryDB {
	return CategoryDB{
		ID:   category.ID().String(),
		Name: category.Name().String(),
	}
}
func categoryToDomain(dto CategoryDB) (domain.Category, error) {
	return domain.NewCategoryWithID(
		dto.ID,
		dto.Name,
	)
}

func (r *CategoryRepository) Save(ctx context.Context, category domain.Category) error {
	row := categoryToDTO(category)
	query, args := categorySQLStruct.InsertInto(sqlCategoryTable, row).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("failed to save category: %v", err)
	}

	return nil
}

func (r *CategoryRepository) Find(ctx context.Context, id domain.CategoryID) (domain.Category, error) {
	sb := categorySQLStruct.SelectFrom(sqlCategoryTable)
	sb.Where(sb.Equal("id", id.String()))
	query, args := sb.Build()
	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	var categoryDTO CategoryDB
	err := r.db.QueryRowContext(ctxTimeout, query, args...).Scan(categorySQLStruct.Addr(&categoryDTO)...)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Category{}, domain.ErrCategoryNotFound
	}
	if err != nil {
		return domain.Category{}, fmt.Errorf("failed to find category: %v", err)
	}

	return categoryToDomain(categoryDTO)
}

func (r *CategoryRepository) FindAll(ctx context.Context) ([]domain.Category, error) {
	sb := categorySQLStruct.SelectFrom(sqlCategoryTable)
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctxTimeout, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to find categories: %v", err)
	}
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		var categoryDTO CategoryDB
		if err := rows.Scan(categorySQLStruct.Addr(&categoryDTO)...); err != nil {
			return nil, fmt.Errorf("failed to scan category: %v", err)
		}
		category, err := categoryToDomain(categoryDTO)
		if err != nil {
			return nil, fmt.Errorf("failed to convert category: %v", err)
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *CategoryRepository) Delete(ctx context.Context, id domain.CategoryID) error {
	sb := categorySQLStruct.DeleteFrom(sqlCategoryTable)
	sb.Where(sb.Equal("id", id.String()))
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	result, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete category: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return domain.ErrCategoryNotFound
	}

	return nil
}

func (r *CategoryRepository) Update(ctx context.Context, category domain.Category) error {
	row := categoryToDTO(category)
	sb := categorySQLStruct.Update(sqlCategoryTable, row)
	sb.Where(sb.Equal("id", row.ID))
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	result, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update category: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return domain.ErrCategoryNotFound
	}

	return nil
}
