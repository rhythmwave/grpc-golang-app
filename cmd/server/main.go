package main

import (
	"context"
	"fmt"
	"go-bponline/m/config"
	"go-bponline/m/internal/database"
	"go-bponline/m/internal/route"
	"net"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/lpernett/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Config struct {
	Port string
}

type App struct {
	Config Config
}

const (
	component = "go-bponline"
)

// InterceptorLogger adapts go-kit logger to interceptor logger.
func interceptorLogger(l log.Logger) logging.Logger {
	return logging.LoggerFunc(func(_ context.Context, lvl logging.Level, msg string, fields ...any) {
		largs := append([]any{"msg", msg}, fields...)
		switch lvl {
		case logging.LevelDebug:
			_ = level.Debug(l).Log(largs...)
		case logging.LevelInfo:
			_ = level.Info(l).Log(largs...)
		case logging.LevelWarn:
			_ = level.Warn(l).Log(largs...)
		case logging.LevelError:
			_ = level.Error(l).Log(largs...)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

func (app *App) initGRPCServer(logger log.Logger) error {
	port := app.Config.Port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		level.Info(logger).Log("msg", "Failed to listen: %v", err)
	}
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}

	rpcLogger := log.With(logger, "service", "gRPC/server", "component", component)
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(interceptorLogger(rpcLogger), opts...),
		),
		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(interceptorLogger(rpcLogger), opts...),
		),
	)
	reflection.Register(grpcServer)
	route.GrpcRoute(grpcServer)
	level.Info(logger).Log("msg", "GRPC Server listening", "addr", lis.Addr())
	return grpcServer.Serve(lis)
}
func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	// Load env file
	err := godotenv.Load()
	if err != nil {
		level.Error(logger).Log("msg", "Error loading .env file")
	}

	//load config
	config.Init()

	// Connect with database
	database.ConnectWithDatabase()

	// Create Config struct
	config := Config{
		Port: os.Getenv("PORT"),
	}

	app := &App{
		Config: config,
	}

	err = app.initGRPCServer(logger)
	if err != nil {
		level.Error(logger).Log("msg", "Failed to serve", "err", err)
		return
	}
}

/*
TODOS
- Up Database
- Implement hot reload
- Separate command to generate model
- Implement auto generate service
- Setup colorful logger both grpc or gorm
- Setup deployment
- Create Markdown
- Improve Your Creativity
*/
