package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/auth"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/server/handler/categories"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/server/handler/groups"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/server/handler/health"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/server/handler/movies"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/server/handler/session"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/server/handler/themes"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/server/handler/tracks"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/server/handler/users"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/server/middleware/admin"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/server/middleware/jwt"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/internal/platform/server/middleware/log_server"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/command"
	"github.com/AlexFJ498/middle-earth-leitmotifs-api/kit/query"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	// shutdownTimeout is the duration to wait for graceful shutdown.
	shutdownTimeout time.Duration

	// jwtKey is the key used to sign the JWT tokens.
	jwtKey auth.JWTKey

	// deps
	commandBus command.Bus
	queryBus   query.Bus
}

func New(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, commandBus command.Bus, queryBus query.Bus, jwtKey auth.JWTKey) (context.Context, Server) {
	srv := Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.New(),

		shutdownTimeout: shutdownTimeout,
		jwtKey:          jwtKey,

		commandBus: commandBus,
		queryBus:   queryBus,
	}

	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("Server running on", s.httpAddr)

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server shut down: %s\n", err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func (s *Server) registerRoutes() {
	s.engine.Use(log_server.Middleware(), gin.Recovery(), gin.Logger())
	s.engine.GET("/health", health.CheckHandler())

	// Public routes
	s.engine.POST("/login", session.LoginHandler(s.queryBus))
	s.engine.GET("/movies", movies.ListHandler(s.queryBus))
	s.engine.GET("/groups", groups.ListHandler(s.queryBus))
	s.engine.GET("/categories", categories.ListHandler(s.queryBus))
	s.engine.GET("/tracks", tracks.ListHandler(s.queryBus))
	s.engine.GET("/themes", themes.ListHandler(s.queryBus))

	// Protected routes
	auth := s.engine.Group("")
	auth.Use(jwt.Middleware(s.jwtKey), admin.Middleware())
	{
		auth.POST("/users", users.CreateHandler(s.commandBus))
		auth.GET("/users", users.ListHandler(s.queryBus))

		auth.POST("/movies", movies.CreateHandler(s.commandBus))
		auth.PUT("/movies/:id", movies.UpdateHandler(s.commandBus))
		auth.DELETE("/movies/:id", movies.DeleteHandler(s.commandBus))

		auth.POST("/groups", groups.CreateHandler(s.commandBus))
		auth.PUT("/groups/:id", groups.UpdateHandler(s.commandBus))
		auth.DELETE("/groups/:id", groups.DeleteHandler(s.commandBus))

		auth.POST("/categories", categories.CreateHandler(s.commandBus))
		auth.PUT("/categories/:id", categories.UpdateHandler(s.commandBus))
		auth.DELETE("/categories/:id", categories.DeleteHandler(s.commandBus))

		auth.POST("/tracks", tracks.CreateHandler(s.commandBus))
		auth.PUT("/tracks/:id", tracks.UpdateHandler(s.commandBus))
		auth.DELETE("/tracks/:id", tracks.DeleteHandler(s.commandBus))

		auth.POST("/themes", themes.CreateHandler(s.commandBus))
		auth.PUT("/themes/:id", themes.UpdateHandler(s.commandBus))
		auth.DELETE("/themes/:id", themes.DeleteHandler(s.commandBus))
	}
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		<-c
		cancel()
	}()

	return ctx
}
