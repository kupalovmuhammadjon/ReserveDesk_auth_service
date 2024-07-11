package service

import (
	pb "auth_service/genproto/auth"
	"auth_service/models"
	"auth_service/storage/postgres"
	"context"
	"database/sql"
	"fmt"
)

type AuthStorage struct {
	pb.UnimplementedAuthServer
	Repo *postgres.AuthRepo
}

func NewAuthStorage(db *sql.DB) *AuthStorage {
	return &AuthStorage{
		Repo: postgres.NewAuthRepo(db),
	}
}

func (a *AuthStorage) Register(ctx context.Context, rep *pb.User) (*pb.Void, error) {
	_, err := a.Repo.Register(rep)
	if err != nil {
		return &pb.Void{}, err
	}
	return &pb.Void{}, nil
}

func (a *AuthStorage) Login(ctx context.Context, rep *pb.UserLogin) (*models.User, error) {
	user, err := a.Repo.Login(rep)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *AuthStorage) UpdateProfile(ctx context.Context,req *pb.User) (*pb.Void, error) {
	
	return &pb.Void{}, nil
}

func (a *AuthStorage) DeleteProfile(ctx context.Context,req *pb.Id) (*pb.Void, error) {
	return &pb.Void{}, nil
}

func (a *AuthStorage) ValidateUserId(ctx context.Context,req *pb.Id) (*pb.Exists, error) {
	exist, err := a.Repo.ValidateUserId(req)
	if err != nil {
		return nil, err
	}

	if !exist.Exists {
		return nil, fmt.Errorf("this id does not user")
	}
	return &pb.Exists{Exists: exist.Exists}, nil 
}










	