package postgres

import (
	pb "auth_service/genproto/auth"
	"auth_service/models"
	"time"

	"database/sql"
	"errors"
)

type AuthRepo struct {
	DB *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{DB: db}
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

	err := us.DB.QueryRow(query, user.FullName, user.IsAdmin, user.Email, user.Password).
		Scan(&User.FullName, &User.IsAdmin, &User.Email, &User.Password)
	if err != nil {
		return &pb.Void{},err
	}

	return &pb.Void{},nil
}

func (us *AuthRepo) Login(logreq *pb.UserLogin) (*models.User, error) {
	user := models.User{}
	query := `select email, password  from users where email = $1 and password = $2 and revoked=false `
	err := us.DB.QueryRow(query, logreq.Email, logreq.Password).Scan(&user.Id, &user.FullName,
		user.IsAdmin, user.Email)
	if err != nil {
		return nil, err
	}
	if user.Password != logreq.Password || user.Email != logreq.Email {
		return nil, errors.New("username or password incorrect")
	}

	return &user, nil
}

func (us *AuthRepo) Logout(token string) error {
	_, err := us.DB.Exec("update reflesh_tokens set deleted_at=$1 where token=$2", time.Now(), token)
	if err != nil {
		return err
	}
	return nil
}

func (us *AuthRepo) ShowProfile(id string) (*pb.Profile, error) {
	userP := pb.Profile{}
	err := us.DB.QueryRow("select full_name, is_admin, email, created_at, updated_at from users where deleted_at is nul").Scan(
							&userP.FullName, &userP.IsAdmin, &userP.Email, &userP.CreatedAt, &userP.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &userP, nil
}

