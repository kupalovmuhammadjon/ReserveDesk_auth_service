package handler

import (
	// pb "auth_service/genproto/auth"

	"auth_service/pkg/logger"
	"auth_service/storage/postgres"
	"database/sql"
	"log"
	"log/slog"
)

type Hendler struct {
	Auth   postgres.AuthRepo
	Logger *slog.Logger
}

func NewHendler(db *sql.DB) *Hendler {
	l, err := logger.New()
	if err != nil {
		log.Fatal(err)
	}
	return &Hendler{
		Auth:   *postgres.NewAuthRepo(db),
		Logger: l,
	}
}
