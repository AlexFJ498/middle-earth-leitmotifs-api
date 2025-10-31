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

type TrackThemeDB struct {
	TrackID     string `db:"track_id"`
	ThemeID     string `db:"theme_id"`
	StartSecond int    `db:"start_second"`
	EndSecond   int    `db:"end_second"`
	IsVariant   bool   `db:"is_variant"`
}

var trackThemeFKMap = map[string]error{
	"tracks_themes_track_id_fkey": domain.ErrTrackNotFound,
	"tracks_themes_theme_id_fkey": domain.ErrThemeNotFound,
}

var sqlTrackThemeTable = "tracks_themes"
var trackThemeSQLStruct = sqlbuilder.NewStruct(new(TrackThemeDB)).For(defaultFlavor)

type TrackThemeRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

func NewTrackThemeRepository(db *sql.DB, timeout time.Duration) *TrackThemeRepository {
	return &TrackThemeRepository{
		db:        db,
		dbTimeout: timeout,
	}
}

func trackThemeToDTO(tt domain.TrackTheme) TrackThemeDB {
	return TrackThemeDB{
		TrackID:     tt.TrackID().String(),
		ThemeID:     tt.ThemeID().String(),
		StartSecond: tt.StartSecond().Int(),
		EndSecond:   tt.EndSecond().Int(),
		IsVariant:   tt.IsVariant().Bool(),
	}
}

func trackThemeToDomain(dto TrackThemeDB) (domain.TrackTheme, error) {
	return domain.NewTrackTheme(dto.TrackID, dto.ThemeID, dto.StartSecond, dto.EndSecond, dto.IsVariant)
}

func (r *TrackThemeRepository) Save(ctx context.Context, trackTheme domain.TrackTheme) error {
	row := trackThemeToDTO(trackTheme)
	sb := trackThemeSQLStruct.InsertInto(sqlTrackThemeTable, row)
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		constraint := extractConstraintName(err)

		err = mapSQLError(extractSQLErrorCode(err))
		if errors.Is(err, ErrForeignKeyViolation) {
			if fkErr, ok := trackThemeFKMap[constraint]; ok {
				return fkErr
			}
		}

		return fmt.Errorf("failed to save track theme: %v", err)
	}

	return nil
}

func (r *TrackThemeRepository) Find(ctx context.Context, trackID domain.TrackID, themeID domain.ThemeID, startSecond domain.StartSecond) (domain.TrackTheme, error) {
	sb := trackThemeSQLStruct.SelectFrom(sqlTrackThemeTable)
	sb.Where(
		sb.Equal("track_id", trackID.String()),
		sb.Equal("theme_id", themeID.String()),
		sb.Equal("start_second", startSecond.Int()),
	)
	query, args := sb.Build()
	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	var trackThemeDTO TrackThemeDB
	err := r.db.QueryRowContext(ctxTimeout, query, args...).Scan(trackThemeSQLStruct.Addr(&trackThemeDTO)...)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.TrackTheme{}, domain.ErrTrackThemeNotFound
	}
	if err != nil {
		return domain.TrackTheme{}, fmt.Errorf("failed to find track theme: %v", err)
	}

	return trackThemeToDomain(trackThemeDTO)
}

func (r *TrackThemeRepository) FindByTrack(ctx context.Context, trackID domain.TrackID) ([]domain.TrackTheme, error) {
	sb := trackThemeSQLStruct.SelectFrom(sqlTrackThemeTable)
	sb.Where(sb.Equal("track_id", trackID.String()))
	sb.OrderBy("start_second ASC")
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctxTimeout, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to find track themes by track ID: %v", err)
	}
	defer rows.Close()

	var trackThemes []domain.TrackTheme
	for rows.Next() {
		var trackThemeDTO TrackThemeDB
		if err := rows.Scan(trackThemeSQLStruct.Addr(&trackThemeDTO)...); err != nil {
			return nil, fmt.Errorf("failed to scan track theme: %v", err)
		}

		trackTheme, err := trackThemeToDomain(trackThemeDTO)
		if err != nil {
			return nil, fmt.Errorf("failed to convert track theme to domain: %v", err)
		}

		trackThemes = append(trackThemes, trackTheme)
	}

	return trackThemes, nil
}

func (r *TrackThemeRepository) Delete(ctx context.Context, trackID domain.TrackID, themeID domain.ThemeID, startSecond domain.StartSecond) error {
	sb := trackThemeSQLStruct.DeleteFrom(sqlTrackThemeTable)
	sb.Where(
		sb.Equal("track_id", trackID.String()),
		sb.Equal("theme_id", themeID.String()),
		sb.Equal("start_second", startSecond.Int()),
	)
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	result, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete track theme: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return domain.ErrTrackThemeNotFound
	}

	return nil
}

func (r *TrackThemeRepository) Update(ctx context.Context, trackTheme domain.TrackTheme) error {
	row := trackThemeToDTO(trackTheme)
	sb := trackThemeSQLStruct.Update(sqlTrackThemeTable, row)
	sb.Where(
		sb.Equal("track_id", row.TrackID),
		sb.Equal("theme_id", row.ThemeID),
		sb.Equal("start_second", row.StartSecond),
	)
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	result, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update track theme: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return domain.ErrTrackThemeNotFound
	}

	return nil
}
