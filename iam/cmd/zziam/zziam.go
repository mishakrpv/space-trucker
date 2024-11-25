package main

import (
	"context"
	stdlog "log"
	"os"

	"github.com/space-trucker/iam/cmd"
	"github.com/space-trucker/iam/pkg/config"
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

	setupLogger(ctx, zConfig.Log)

	_, err := setupServer(ctx, &zConfig.Cfg)
	if err != nil {
		return err
	}

	return nil
}

func setupLogger(ctx context.Context, cfg *config.Log) {

}

func setupServer(ctx context.Context, cfg *config.Cfg) (*server.Server, error) {
	return server.New(), nil
}
