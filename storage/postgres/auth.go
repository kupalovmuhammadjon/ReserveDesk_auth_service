package postgres

import (
	pb "auth_service/genproto/auth"
	"auth_service/models"

	"database/sql"
	"errors"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (us *AuthRepo) Register(user *pb.User)  (*pb.Void, error) {
	query := `
		insert into users(
		    full_name,
			is_admin,
			email,
			password	
		) values($1,$2,$3,$4)
		returning full_name, is_admin, email, password
	`
	var User pb.User

	err := us.db.QueryRow(query, user.FullName, user.IsAdmin, user.Email, user.Password).
		Scan(&User.FullName, &User.IsAdmin, &User.Email, &User.Password)
	if err != nil {
		return &pb.Void{},err
	}

	return &pb.Void{},nil
}

func (us *AuthRepo) Login(logreq *pb.UserLogin) (*models.User, error) {
	user := models.User{}
	query := `select email, password  from users where email = $1 and password = $2`
	err := us.db.QueryRow(query, logreq.Email, logreq.Password).Scan(&user.Id, &user.FullName,
		user.IsAdmin, user.Email)
	if err != nil {
		return nil, err
	}
	qualify := true
	if user.Password != logreq.Password || user.Email != logreq.Email {
		qualify = false
	}
	if !qualify {
		return nil, errors.New("username or password incorrect")
	}
	return &user, nil
}