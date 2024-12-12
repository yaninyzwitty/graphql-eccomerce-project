package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/yaninyzwitty/gqlgen-eccomerce-project/graph"
	"github.com/yaninyzwitty/gqlgen-eccomerce-project/internal/database"
	"github.com/yaninyzwitty/gqlgen-eccomerce-project/pkg"
)

var (
	password string
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	file, err := os.Open("config.yml") //refactor with docker
	if err != nil {
		slog.Error("failed to open config.yaml ")
		os.Exit(1)
	}
	var cfg pkg.Config
	if err := cfg.LoadConfig(file); err != nil {
		slog.Error("failed to load config.yaml")
		os.Exit(1)
	}
	// NOT NEEDED WHEN YOU ARE USING COMPOSE
	// err = godotenv.Load("dev.env")
	// if err != nil {
	// 	slog.Error("failed to load dev.env")
	// 	os.Exit(1)
	// }

	fmt.Println(os.Getenv("DB_PASSWORD"))

	if s := os.Getenv("DB_PASSWORD"); s != "" {
		password = s
	}

	databaseCfg := database.NewDbConfig(cfg.Database.User, password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Database, cfg.Database.SSLMode)
	pool, err := databaseCfg.MakeNewPgxPool(ctx, 30)
	if err != nil {
		slog.Error("failed to make new pgx pool", "error", err)
		os.Exit(1)
	}

	defer pool.Close()
	// ping db
	err = databaseCfg.Ping(ctx)
	if err != nil {
		slog.Error("failed to ping db", "error", err)
		os.Exit(1)
	}

	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	resolvers := &graph.Resolver{
		Pool: pool,
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolvers}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: mux,
	}

	stopCH := make(chan os.Signal, 1)
	signal.Notify(stopCH, os.Interrupt, syscall.SIGTERM)
	go func() {
		slog.Info("server is listening on :" + fmt.Sprintf("%d", cfg.Server.Port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("failed to start server")
			os.Exit(1)
		}

	}()
	<-stopCH
	slog.Info("shuttting down the server...")
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server")
		os.Exit(1)
	} else {
		slog.Info("server stopped down gracefully")

	}

}
