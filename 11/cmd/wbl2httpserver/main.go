package main

import (
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	application := app.New(log)

	application.MustRun(cfg.HTTPServer, cfg.PostgresConfig, cfg.AppConfig)

}

func setupLogger(env string) *slog.Logger {

	return setupPrettySlog()
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
