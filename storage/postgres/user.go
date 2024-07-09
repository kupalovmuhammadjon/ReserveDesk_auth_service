package postgres

import "database/sql"

type UserRepo struct{
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo{
	return &UserRepo{Db: db}
}

