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

func (us *AuthRepo) UpdateProfile(user *pb.User) (*pb.Void, error) {
	_, err := us.db.Exec("update users set full_name=$1, is_admin=$2, email=$3, password=$4, updated_at=$6 where id=$5", user.FullName, user.IsAdmin, user.Email, user.Password, user.Id, time.Now())
	if err != nil {
		return &pb.Void{}, err
	}

	return &pb.Void{}, nil

}

func (us *AuthRepo) DeleteProfile(rep *pb.Id) (*pb.Void, error){
	_, err := uuid.Parse(rep.Id)
	if err != nil {
		log.Printf("Error parsing UUID: %v", err)
		return &pb.Void{},err
	}
	_, err = us.db.Exec("update users set delete_at=$1 where id=$2", time.Now(), rep.Id)
	if err != nil {
		return &pb.Void{},err
	}	
	return &pb.Void{},nil
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
	err := us.db.QueryRow(query, rep.Id).Scan(&res.Exists)
	return &res, err
}
