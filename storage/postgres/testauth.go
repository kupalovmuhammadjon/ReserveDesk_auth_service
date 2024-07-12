package postgres

import (
	"testing"

	pb "auth_service/genproto/auth"
)

func newAuthRepo(t *testing.T) *AuthRepo {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	return &AuthRepo{DB: db}
}

func TestRegister(t *testing.T) {
	a := newAuthRepo(t)

	user := pb.User{
		FullName: "QWERTY",
		IsAdmin:  false,
		Email:    "fgfrf@gmail.com",
		Password: "gfghre43",
	}

	_, err := a.Register(&user)
	if err != nil {
		t.Fatalf("Failed to register user: %v", err)
	}
}

func TestRegisterDuplicate(t *testing.T) {
	a := newAuthRepo(t)

	user := pb.User{
		FullName: "QWERTY",
		IsAdmin:  false,
		Email:    "fgfrf@gmail.com",
		Password: "gfghre43",
	}

	_, err := a.Register(&user)
	if err != nil {
		t.Fatalf("Failed to register user: %v", err)
	}

	// Trying to register the same user again to check for duplicate handling
	_, err = a.Register(&user)
	if err == nil {
		t.Fatal("Expected an error when registering a duplicate user, but got none")
	}
}
