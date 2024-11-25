package main

import (
	"context"
	stdlog "log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/space-trucker/iam/cmd"
	"github.com/space-trucker/iam/pkg/api"
	"github.com/space-trucker/iam/pkg/config"
	"github.com/space-trucker/iam/pkg/logs"
	"github.com/space-trucker/iam/pkg/server"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		stdlog.Printf("%s\n", err)
		os.Exit(1)
	}
}

func run(
	ctx context.Context,
) error {
	zConfig := cmd.NewZzIamCmdConfiguration()

	logs.SetupLogger(zConfig.Log)

	srv, err := setupServer(ctx, &zConfig.Cfg)
	if err != nil {
		return err
	}

	ctx, _ = signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)

	srv.Start(ctx)
	defer srv.Close()

	srv.Wait()
	log.Info().Msg("Shutting down")
	return nil
}

func setupServer(ctx context.Context, cfg *config.Cfg) (*server.Server, error) {
	tlsConfig, err := cfg.ClientTLS.CreateTLSConfig(ctx)
	if err != nil {
		return nil, err
	}

	services := api.NewServices(cfg)

	httpServer := &http.Server{
		Addr:      net.JoinHostPort(cfg.Host, cfg.Port),
		TLSConfig: tlsConfig,
		Handler:   api.NewHTTPHandler(services),
	}

	return server.NewServer(httpServer), nil
}
