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

type TrackDB struct {
	ID         string  `db:"id"`
	Name       string  `db:"name"`
	MovieID    string  `db:"movie_id"`
	SpotifyURL *string `db:"spotify_url"`
}

var sqlTrackTable = "tracks"
var trackSQLStruct = sqlbuilder.NewStruct(new(TrackDB)).For(defaultFlavor)

type TrackRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

func NewTrackRepository(db *sql.DB, timeout time.Duration) *TrackRepository {
	return &TrackRepository{
		db:        db,
		dbTimeout: timeout,
	}
}

func trackToDTO(track domain.Track) TrackDB {
	return TrackDB{
		ID:         track.ID().String(),
		Name:       track.Name().String(),
		MovieID:    track.MovieID().String(),
		SpotifyURL: track.SpotifyURL().AsStringPtr(),
	}
}

func trackToDomain(dto TrackDB) (domain.Track, error) {
	return domain.NewTrackWithID(dto.ID, dto.Name, dto.MovieID, dto.SpotifyURL)
}

func (r *TrackRepository) Save(ctx context.Context, track domain.Track) error {
	row := trackToDTO(track)
	sb := trackSQLStruct.InsertInto(sqlTrackTable, row)
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		err = mapSQLError(extractSQLErrorCode(err))

		if errors.Is(err, ErrForeignKeyViolation) {
			return domain.ErrMovieNotFound
		}

		return fmt.Errorf("failed to save track: %v", err)
	}

	return nil
}

func (r *TrackRepository) Find(ctx context.Context, id domain.TrackID) (domain.Track, error) {
	sb := trackSQLStruct.SelectFrom(sqlTrackTable)
	sb.Where(sb.Equal("id", id.String()))
	query, args := sb.Build()
	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	var trackDTO TrackDB
	err := r.db.QueryRowContext(ctxTimeout, query, args...).Scan(trackSQLStruct.Addr(&trackDTO)...)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Track{}, domain.ErrTrackNotFound
	}
	if err != nil {
		return domain.Track{}, fmt.Errorf("failed to find track: %v", err)
	}

	return trackToDomain(trackDTO)
}

func (r *TrackRepository) FindAll(ctx context.Context) ([]domain.Track, error) {
	sb := trackSQLStruct.SelectFrom(sqlTrackTable)
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctxTimeout, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to find tracks: %v", err)
	}
	defer rows.Close()

	var tracks []domain.Track
	for rows.Next() {
		var trackDTO TrackDB
		err := rows.Scan(trackSQLStruct.Addr(&trackDTO)...)
		if err != nil {
			return nil, fmt.Errorf("failed to scan track: %v", err)
		}
		track, err := trackToDomain(trackDTO)
		if err != nil {
			return nil, fmt.Errorf("failed to convert track: %v", err)
		}
		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (r *TrackRepository) FindByMovie(ctx context.Context, movieID domain.MovieID) ([]domain.Track, error) {
	sb := trackSQLStruct.SelectFrom(sqlTrackTable)
	sb.Where(sb.Equal("movie_id", movieID.String()))
	sb.OrderBy("created_at ASC")
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctxTimeout, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to find tracks by movie: %v", err)
	}
	defer rows.Close()

	var tracks []domain.Track
	for rows.Next() {
		var trackDTO TrackDB
		err := rows.Scan(trackSQLStruct.Addr(&trackDTO)...)
		if err != nil {
			return nil, fmt.Errorf("failed to scan track: %v", err)
		}
		track, err := trackToDomain(trackDTO)
		if err != nil {
			return nil, fmt.Errorf("failed to convert track: %v", err)
		}
		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (r *TrackRepository) Delete(ctx context.Context, id domain.TrackID) error {
	sb := trackSQLStruct.DeleteFrom(sqlTrackTable)
	sb.Where(sb.Equal("id", id.String()))
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	result, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete track: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return domain.ErrTrackNotFound
	}

	return nil
}

func (r *TrackRepository) Update(ctx context.Context, track domain.Track) error {
	row := trackToDTO(track)
	sb := trackSQLStruct.Update(sqlTrackTable, row)
	sb.Where(sb.Equal("id", row.ID))
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	result, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update track: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return domain.ErrTrackNotFound
	}

	return nil
}
