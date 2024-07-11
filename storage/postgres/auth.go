package postgres

import (
	pb "auth_service/genproto/auth"
	"auth_service/models"
	"log"
	"time"

	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type AuthRepo struct {
	DB *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{DB: db}
}

func (us *AuthRepo) Register(user *pb.User) (*pb.Void, error) {
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
		return &pb.Void{}, err
	}

	return &pb.Void{}, nil
}

func (us *AuthRepo) Login(logreq *models.User) (*models.User, error) {
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

func (us *AuthRepo) UpdateProfile(user *pb.User) (*pb.Void, error) {
	_, err := us.DB.Exec("update users set full_name=$1, is_admin=$2, email=$3, password=$4, updated_at=$6 where id=$5", user.FullName, user.IsAdmin, user.Email, user.Password, user.Id, time.Now())
	if err != nil {
		return &pb.Void{}, err
	}

	return &pb.Void{}, nil

}

func (us *AuthRepo) DeleteProfile(rep *pb.Id) (*pb.Void, error) {
	_, err := uuid.Parse(rep.Id)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return &pb.Void{}, err
	}
	_, err = us.DB.Exec("update users set delete_at=$1 where id=$2", time.Now(), rep.Id)
	if err != nil {
		return &pb.Void{}, err
	}
	return &pb.Void{}, nil
}

func (us *AuthRepo) ValidateUserId(rep *pb.Id) (*pb.Exists, error) {
	query := `select 
	            case 
				    when id = $1 then true 
				else 
				    false 
				end 
			from 
			    users 
			where 
			    id = $1 and deletad_at is null`
	res := pb.Exists{}
	err := us.DB.QueryRow(query, rep.Id).Scan(&res.Exists)
	return &res, err
}
func (us *AuthRepo) Logout(token *pb.Token) error {
	_, err := us.DB.Exec("update reflesh_tokens set deleted_at=$1 where token=$2", time.Now(), token.Token)
	if err != nil {
		return err
	}
	return nil
}


func (us *AuthRepo) ShowProfile(id *pb.Id) (*pb.Profile, error) {
	userP := pb.Profile{}
	err := us.DB.QueryRow("select full_name, is_admin, email, created_at, updated_at from users where deleted_at is nul").Scan(
		&userP.FullName, &userP.IsAdmin, &userP.Email, &userP.CreatedAt, &userP.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &userP, nil
}


