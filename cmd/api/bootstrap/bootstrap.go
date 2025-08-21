package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/authenticating"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/creating"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/auth"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/bus/inmemory"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/server"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/storage/sqldb"
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

	creatingUserService := creating.NewUserService(userRepository, eventBus)
	commandBus.Register(creating.UserCommandType, creating.NewUserCommandHandler(creatingUserService))

	authenticatingService := authenticating.NewLoginService(userRepository, cfg.Jwtkey, cfg.Jwtexpires)
	queryBus.Register(authenticating.LoginQueryType, authenticating.NewLoginQueryHandler(authenticatingService))

	// At the moment, this is not implemented. It shows how an inmemory event bus can be used to handle events.
	// increasingUserCounterService := increasing.NewUserCounterIncreaserService()
	// eventBus.Subscribe(
	// 	domain.UserCreatedEventType,
	// 	creating.NewIncreaseUsersCounterOnUserCreated(increasingUserCounterService),
	// )

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.Shutdowntimeout, commandBus, queryBus)
	return srv.Run(ctx)
}
