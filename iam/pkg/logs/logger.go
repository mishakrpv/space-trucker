package logs

import (
	"io"
	stdlog "log"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/space-trucker/iam/pkg/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	// hide the first logs before the setup of the logger.
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
}

func SetupLogger(logCfg *config.Log) {
	// configure log format
	w := getLogWriter(logCfg)

	// configure log level
	logLevel := getLogLevel(logCfg)

	// create logger
	logCtx := zerolog.New(w).With().Timestamp()
	if logLevel <= zerolog.DebugLevel {
		logCtx = logCtx.Caller()
	}

	log.Logger = logCtx.Logger().Level(logLevel)
	zerolog.DefaultContextLogger = &log.Logger
	zerolog.SetGlobalLevel(logLevel)

	// configure default standard log.
	stdlog.SetFlags(stdlog.Lshortfile | stdlog.LstdFlags)
	stdlog.SetOutput(NoLevel(log.Logger, zerolog.DebugLevel))
}

func getLogWriter(logCfg *config.Log) io.Writer {
	var w io.Writer = os.Stdout

	if logCfg != nil && len(logCfg.FilePath) > 0 {
		_, _ = os.OpenFile(logCfg.FilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
		w = &lumberjack.Logger{
			Filename:   logCfg.FilePath,
			MaxSize:    logCfg.MaxSize,
			MaxBackups: logCfg.MaxBackups,
			MaxAge:     logCfg.MaxAge,
			Compress:   true,
		}
	}

	if logCfg == nil || logCfg.Format != "json" {
		w = zerolog.ConsoleWriter{
			Out:        w,
			TimeFormat: time.RFC3339,
			NoColor:    logCfg != nil && (logCfg.NoColor || len(logCfg.FilePath) > 0),
		}
	}

	return w
}

func getLogLevel(logCfg *config.Log) zerolog.Level {
	levelStr := "error"
	if logCfg != nil && logCfg.Level != "" {
		levelStr = strings.ToLower(logCfg.Level)
	}

	logLevel, err := zerolog.ParseLevel(strings.ToLower(levelStr))
	if err != nil {
		log.Error().Err(err).
			Str("logLevel", levelStr).
			Msg("Unspecified or invalid log level, setting the level to default (ERROR)...")

		logLevel = zerolog.ErrorLevel
	}

	return logLevel
}
