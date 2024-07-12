package service

import (
	pb "auth_service/genproto/auth"
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
	return &AuthStorage{Repo: postgres.NewAuthRepo(db)}
}


func (a *AuthStorage) UpdateProfile(ctx context.Context, req *pb.User) (*pb.Void, error) {
	_, err := a.Repo.UpdateProfile(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (a *AuthStorage) DeleteProfile(ctx context.Context, req *pb.Id) (*pb.Void, error) {
	_, err := a.Repo.DeleteProfile(req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (a *AuthStorage) ValidateUserId(ctx context.Context, req *pb.Id) (*pb.Exists, error) {
	exist, err := a.Repo.ValidateUserId(req)
	if err != nil {
		return nil, err
	}

	if !exist.Exists {
		return nil, fmt.Errorf("this id does not user")
	}
	return &pb.Exists{Exists: exist.Exists}, nil
}

func (a *AuthStorage) ShowProfile(cnt context.Context, req *pb.Id) (*pb.Profile, error) {
	userP, err := a.Repo.ShowProfile(req)
	if err != nil {
		return nil, err
	}
	return userP, nil
}