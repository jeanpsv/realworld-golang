package router

import (
	"log/slog"

	"github.com/jeanpsv/realworld-golang/internal/data"
)

type HttpConfig struct {
	Port int
	Env  string
	Db   struct {
		Dsn string
	}
}

type HttpApplication struct {
	Config HttpConfig
	Logger *slog.Logger
	Models data.Models
}
