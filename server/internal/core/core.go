package core

import (
	"database/sql"

	"go.uber.org/zap"

	"github.com/SavelyBerdnik/Airlines_amonic/server/internal/infrastructure/config"
)

type Core struct {
	logger *zap.Logger
	db     *sql.DB
	cfg    *config.Config
}

func NewCore(l *zap.Logger, db *sql.DB, c *config.Config) *Core {
	return &Core{
		logger: l,
		db:     db,
		cfg:    c,
	}
}
