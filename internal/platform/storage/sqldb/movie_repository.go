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

type MovieDB struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

var sqlMovieTable = "movies"
var movieSQLStruct = sqlbuilder.NewStruct(new(MovieDB)).For(defaultFlavor)

// MovieRepository implements the MovieRepository interface for SQL.
type MovieRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewMovieRepository creates a new MovieRepository.
func NewMovieRepository(db *sql.DB, dbTimeout time.Duration) *MovieRepository {
	return &MovieRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

func movieToDTO(movie domain.Movie) MovieDB {
	return MovieDB{
		ID:   movie.ID().String(),
		Name: movie.Name().String(),
	}
}
func movieToDomain(dto MovieDB) (domain.Movie, error) {
	return domain.NewMovieWithID(
		dto.ID,
		dto.Name,
	)
}

func (r *MovieRepository) Save(ctx context.Context, movie domain.Movie) error {
	row := movieToDTO(movie)
	query, args := movieSQLStruct.InsertInto(sqlMovieTable, row).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("failed to save movie: %v", err)
	}

	return nil
}

func (r *MovieRepository) Find(ctx context.Context, id domain.MovieID) (domain.Movie, error) {
	sb := movieSQLStruct.SelectFrom(sqlMovieTable)
	sb.Where(sb.Equal("id", id.String()))
	query, args := sb.Build()
	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	var movieDTO MovieDB
	err := r.db.QueryRowContext(ctxTimeout, query, args...).Scan(movieSQLStruct.Addr(&movieDTO)...)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Movie{}, domain.ErrMovieNotFound
	}
	if err != nil {
		return domain.Movie{}, fmt.Errorf("failed to find movie: %v", err)
	}

	return movieToDomain(movieDTO)
}

func (r *MovieRepository) FindAll(ctx context.Context) ([]domain.Movie, error) {
	sb := movieSQLStruct.SelectFrom(sqlMovieTable)
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctxTimeout, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to find movies: %v", err)
	}
	defer rows.Close()

	var movies []domain.Movie
	for rows.Next() {
		var movieDTO MovieDB
		if err := rows.Scan(movieSQLStruct.Addr(&movieDTO)...); err != nil {
			return nil, fmt.Errorf("failed to scan movie: %v", err)
		}
		movie, err := movieToDomain(movieDTO)
		if err != nil {
			return nil, fmt.Errorf("failed to convert movie: %v", err)
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (r *MovieRepository) Delete(ctx context.Context, id domain.MovieID) error {
	sb := movieSQLStruct.DeleteFrom(sqlMovieTable)
	sb.Where(sb.Equal("id", id.String()))
	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	result, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete movie: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return domain.ErrMovieNotFound
	}

	return nil
}
