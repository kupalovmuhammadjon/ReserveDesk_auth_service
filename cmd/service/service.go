package main

// import (
// 	"auth_service/service"
// 	pb "auth_service/genproto/auth"
// 	"auth_service/config"
// 	"auth_service/storage/postgres"
// 	"log"
// 	"net"

// 	"google.golang.org/grpc"
// )

// func main() {
// 	cfg := config.Load()
// 	listrner, err := net.Listen("tcp", cfg.AUTH_SERVICE_PORT)
// 	if err != nil{
// 		log.Fatal(err)
// 	}
// 	defer listrner.Close()
// 	db, err := postgres.ConnectDB()
// 	if err != nil{
// 		log.Fatal(err)
// 	}
// 	authService := service.NewAuthStorage(db)
// 	s := grpc.NewServer()
// 	pb.RegisterAuthServer(s, authService)
// }