package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/authenticating"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/creating"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/deleting"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/listing"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/auth"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/bus/inmemory"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/server"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/storage/sqldb"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/updating"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type config struct {
	// Server configuration
	Host            string
	Port            uint
	Shutdowntimeout time.Duration

	// Database configuration
	Dbuser     string
	Dbpassword string
	Dbhost     string
	Dbport     int
	Dbname     string
	Dbtimeout  time.Duration

	// Login configuration
	Jwtkey     auth.JWTKey
	Jwtexpires time.Duration
}

func Run() error {

	if os.Getenv("MELA_ENV") == "" {
		if err := godotenv.Load(".env.local"); err != nil {
			return fmt.Errorf("error loading .env file: %w", err)
		}
	}

	var cfg config
	err := envconfig.Process("mela", &cfg)
	if err != nil {
		return err
	}

	postgreURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.Dbuser, cfg.Dbpassword, cfg.Dbhost, cfg.Dbport, cfg.Dbname)
	db, err := sql.Open("postgres", postgreURI)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	var (
		commandBus = inmemory.NewCommandBus()
		queryBus   = inmemory.NewQueryBus()
		eventBus   = inmemory.NewEventBus()
	)

	userRepository := sqldb.NewUserRepository(db, cfg.Dbtimeout)
	movieRepository := sqldb.NewMovieRepository(db, cfg.Dbtimeout)

	authenticatingService := authenticating.NewLoginService(userRepository, cfg.Jwtkey, cfg.Jwtexpires)
	queryBus.Register(authenticating.LoginQueryType, authenticating.NewLoginQueryHandler(authenticatingService))

	creatingUserService := creating.NewUserService(userRepository, eventBus)
	commandBus.Register(creating.UserCommandType, creating.NewUserCommandHandler(creatingUserService))
	commandBus.Register(creating.MovieCommandType, creating.NewMovieCommandHandler(creating.NewMovieService(sqldb.NewMovieRepository(db, cfg.Dbtimeout))))

	listingUserService := listing.NewUserService(userRepository)
	listingMovieService := listing.NewMovieService(movieRepository)
	queryBus.Register(listing.UsersQueryType, listing.NewUsersQueryHandler(listingUserService))
	queryBus.Register(listing.MoviesQueryType, listing.NewMoviesQueryHandler(listingMovieService))

	updatingMovieService := updating.NewMovieService(movieRepository)
	commandBus.Register(updating.MovieCommandType, updating.NewMovieCommandHandler(updatingMovieService))

	deletingMovieService := deleting.NewMovieService(movieRepository)
	commandBus.Register(deleting.MovieCommandType, deleting.NewMovieCommandHandler(deletingMovieService))

	// At the moment, this is not implemented. It shows how an inmemory event bus can be used to handle events.
	// increasingUserCounterService := increasing.NewUserCounterIncreaserService()
	// eventBus.Subscribe(
	// 	domain.UserCreatedEventType,
	// 	creating.NewIncreaseUsersCounterOnUserCreated(increasingUserCounterService),
	// )

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.Shutdowntimeout, commandBus, queryBus, cfg.Jwtkey)
	return srv.Run(ctx)
}
