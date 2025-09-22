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

type ThemeDB struct {
	ID          string  `db:"id"`
	Name        string  `db:"name"`
	FirstHeard  string  `db:"first_heard"`
	GroupID     string  `db:"group_id"`
	Description string  `db:"description"`
	CategoryID  *string `db:"category_id"`
}

var fkMap = map[string]error{
	"themes_category_id_fkey": domain.ErrCategoryNotFound,
	"themes_group_id_fkey":    domain.ErrGroupNotFound,
	"themes_first_heard_fkey": domain.ErrCategoryNotFound,
}

var sqlThemeTable = "themes"
var themeSQLStruct = sqlbuilder.NewStruct(new(ThemeDB)).For(defaultFlavor)

type ThemeRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

func NewThemeRepository(db *sql.DB, timeout time.Duration) *ThemeRepository {
	return &ThemeRepository{
		db:        db,
		dbTimeout: timeout,
	}
}

func themeToDTO(theme domain.Theme) ThemeDB {
	var categoryID *string
	if theme.CategoryID() != nil {
		str := theme.CategoryID().String()
		categoryID = &str
	}

	return ThemeDB{
		ID:          theme.ID().String(),
		Name:        theme.Name().String(),
		FirstHeard:  theme.FirstHeard().String(),
		GroupID:     theme.GroupID().String(),
		Description: theme.Description().String(),
		CategoryID:  categoryID,
	}
}

func themeToDomain(dto ThemeDB) (domain.Theme, error) {
	return domain.NewThemeWithID(
		dto.ID,
		dto.Name,
		dto.FirstHeard,
		dto.GroupID,
		dto.Description,
		dto.CategoryID)
}

func (r *ThemeRepository) Save(ctx context.Context, theme domain.Theme) error {
	row := themeToDTO(theme)
	sb := themeSQLStruct.InsertInto(sqlThemeTable, row)
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		constraint := extractConstraintName(err)
		err = mapSQLError(extractSQLErrorCode(err))
		if errors.Is(err, ErrForeignKeyViolation) {
			if fkErr, ok := fkMap[constraint]; ok {
				return fkErr
			}
		}

		return fmt.Errorf("failed to save theme: %v", err)
	}

	return nil
}

func (r *ThemeRepository) Find(ctx context.Context, id domain.ThemeID) (domain.Theme, error) {
	sb := themeSQLStruct.SelectFrom(sqlThemeTable)
	sb.Where(sb.Equal("id", id.String()))
	query, args := sb.Build()
	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	var themeDTO ThemeDB
	err := r.db.QueryRowContext(ctxTimeout, query, args...).Scan(themeSQLStruct.Addr(&themeDTO)...)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Theme{}, domain.ErrThemeNotFound
	}
	if err != nil {
		return domain.Theme{}, fmt.Errorf("failed to find theme: %v", err)
	}

	return themeToDomain(themeDTO)
}

func (r *ThemeRepository) FindAll(ctx context.Context) ([]domain.Theme, error) {
	sb := themeSQLStruct.SelectFrom(sqlThemeTable)
	sb.OrderBy("created_at ASC")
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctxTimeout, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to find all themes: %v", err)
	}
	defer rows.Close()

	var themes []domain.Theme
	for rows.Next() {
		var themeDTO ThemeDB
		if err := rows.Scan(themeSQLStruct.Addr(&themeDTO)...); err != nil {
			return nil, fmt.Errorf("failed to scan theme: %v", err)
		}

		theme, err := themeToDomain(themeDTO)
		if err != nil {
			return nil, fmt.Errorf("failed to convert theme: %v", err)
		}

		themes = append(themes, theme)
	}

	return themes, nil
}

func (r *ThemeRepository) FindByGroup(ctx context.Context, groupID domain.GroupID) ([]domain.Theme, error) {
	sb := themeSQLStruct.SelectFrom(sqlThemeTable)
	sb.Where(sb.Equal("group_id", groupID.String()))
	sb.OrderBy("created_at ASC")
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctxTimeout, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to find themes by group: %v", err)
	}
	defer rows.Close()

	var themes []domain.Theme
	for rows.Next() {
		var themeDTO ThemeDB
		if err := rows.Scan(themeSQLStruct.Addr(&themeDTO)...); err != nil {
			return nil, fmt.Errorf("failed to scan theme: %v", err)
		}

		theme, err := themeToDomain(themeDTO)
		if err != nil {
			return nil, fmt.Errorf("failed to convert theme: %v", err)
		}

		themes = append(themes, theme)
	}

	return themes, nil
}

func (r *ThemeRepository) Delete(ctx context.Context, id domain.ThemeID) error {
	sb := themeSQLStruct.DeleteFrom(sqlThemeTable)
	sb.Where(sb.Equal("id", id.String()))
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	result, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete theme: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return domain.ErrThemeNotFound
	}

	return nil
}

func (r *ThemeRepository) Update(ctx context.Context, theme domain.Theme) error {
	row := themeToDTO(theme)
	sb := themeSQLStruct.Update(sqlThemeTable, row)
	sb.Where(sb.Equal("id", row.ID))
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	result, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update theme: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return domain.ErrThemeNotFound
	}

	return nil
}
