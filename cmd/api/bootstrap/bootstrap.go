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
	Port            uint `envconfig:"PORT"`
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

	// Frontend configuration
	Frontendurl string
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

	var postgreURI string
	if os.Getenv("DATABASE_URL") != "" {
		postgreURI = os.Getenv("DATABASE_URL")
	} else {
		postgreURI = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.Dbuser, cfg.Dbpassword, cfg.Dbhost, cfg.Dbport, cfg.Dbname)
	}
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
	groupRepository := sqldb.NewGroupRepository(db, cfg.Dbtimeout)
	categoryRepository := sqldb.NewCategoryRepository(db, cfg.Dbtimeout)
	trackRepository := sqldb.NewTrackRepository(db, cfg.Dbtimeout)
	themeRepository := sqldb.NewThemeRepository(db, cfg.Dbtimeout)

	authenticatingService := authenticating.NewLoginService(userRepository, cfg.Jwtkey, cfg.Jwtexpires)
	queryBus.Register(authenticating.LoginQueryType, authenticating.NewLoginQueryHandler(authenticatingService))

	creatingUserService := creating.NewUserService(userRepository, eventBus)
	creatingMovieService := creating.NewMovieService(movieRepository)
	creatingGroupService := creating.NewGroupService(groupRepository)
	creatingCategoryService := creating.NewCategoryService(categoryRepository)
	creatingTrackService := creating.NewTrackService(trackRepository)
	creatingThemeService := creating.NewThemeService(themeRepository)
	commandBus.Register(creating.UserCommandType, creating.NewUserCommandHandler(creatingUserService))
	commandBus.Register(creating.MovieCommandType, creating.NewMovieCommandHandler(creatingMovieService))
	commandBus.Register(creating.GroupCommandType, creating.NewGroupCommandHandler(creatingGroupService))
	commandBus.Register(creating.CategoryCommandType, creating.NewCategoryCommandHandler(creatingCategoryService))
	commandBus.Register(creating.TrackCommandType, creating.NewTrackCommandHandler(creatingTrackService))
	commandBus.Register(creating.ThemeCommandType, creating.NewThemeCommandHandler(creatingThemeService))

	listingUserService := listing.NewUserService(userRepository)
	listingMovieService := listing.NewMovieService(movieRepository)
	listingGroupService := listing.NewGroupService(groupRepository)
	listingCategoryService := listing.NewCategoryService(categoryRepository)
	listingTrackService := listing.NewTrackService(trackRepository, listingMovieService)
	listingThemeService := listing.NewThemeService(themeRepository, listingTrackService, listingGroupService, listingCategoryService)
	queryBus.Register(listing.UsersQueryType, listing.NewUsersQueryHandler(listingUserService))
	queryBus.Register(listing.MoviesQueryType, listing.NewMoviesQueryHandler(listingMovieService))
	queryBus.Register(listing.GroupsQueryType, listing.NewGroupsQueryHandler(listingGroupService))
	queryBus.Register(listing.CategoriesQueryType, listing.NewCategoriesQueryHandler(listingCategoryService))
	queryBus.Register(listing.TracksQueryType, listing.NewTracksQueryHandler(listingTrackService))
	queryBus.Register(listing.ThemesQueryType, listing.NewThemesQueryHandler(listingThemeService))

	updatingMovieService := updating.NewMovieService(movieRepository)
	updatingGroupService := updating.NewGroupService(groupRepository)
	updatingCategoryService := updating.NewCategoryService(categoryRepository)
	updatingTrackService := updating.NewTrackService(trackRepository)
	updatingThemeService := updating.NewThemeService(themeRepository)
	commandBus.Register(updating.MovieCommandType, updating.NewMovieCommandHandler(updatingMovieService))
	commandBus.Register(updating.GroupCommandType, updating.NewGroupCommandHandler(updatingGroupService))
	commandBus.Register(updating.CategoryCommandType, updating.NewCategoryCommandHandler(updatingCategoryService))
	commandBus.Register(updating.TrackCommandType, updating.NewTrackCommandHandler(updatingTrackService))
	commandBus.Register(updating.ThemeCommandType, updating.NewThemeCommandHandler(updatingThemeService))

	deletingMovieService := deleting.NewMovieService(movieRepository)
	deletingGroupService := deleting.NewGroupService(groupRepository)
	deletingCategoryService := deleting.NewCategoryService(categoryRepository)
	deletingTrackService := deleting.NewTrackService(trackRepository)
	deletingThemeService := deleting.NewThemeService(themeRepository)
	commandBus.Register(deleting.MovieCommandType, deleting.NewMovieCommandHandler(deletingMovieService))
	commandBus.Register(deleting.GroupCommandType, deleting.NewGroupCommandHandler(deletingGroupService))
	commandBus.Register(deleting.CategoryCommandType, deleting.NewCategoryCommandHandler(deletingCategoryService))
	commandBus.Register(deleting.TrackCommandType, deleting.NewTrackCommandHandler(deletingTrackService))
	commandBus.Register(deleting.ThemeCommandType, deleting.NewThemeCommandHandler(deletingThemeService))

	// At the moment, this is not implemented. It shows how an inmemory event bus can be used to handle events.
	// increasingUserCounterService := increasing.NewUserCounterIncreaserService()
	// eventBus.Subscribe(
	// 	domain.UserCreatedEventType,
	// 	creating.NewIncreaseUsersCounterOnUserCreated(increasingUserCounterService),
	// )

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.Shutdowntimeout, commandBus, queryBus, cfg.Jwtkey, cfg.Frontendurl)
	return srv.Run(ctx)
}
