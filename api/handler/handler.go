package handler

import (
	// pb "auth_service/genproto/auth"
	"auth_service/storage/postgres"
	"database/sql"
)

type Hendler struct {
	Auth postgres.AuthRepo
}

func NewHendler(db *sql.DB) *Hendler {
	return &Hendler{*postgres.NewAuthRepo(db)}
}


