package main

import (
	"log"
	"net"

	grpcSvc "github.com/tubagusmf/comment-service/internal/delivery/grpc"
	"github.com/tubagusmf/comment-service/internal/repository"
	"github.com/tubagusmf/comment-service/internal/usecase"
	pb "github.com/tubagusmf/comment-service/pb/comment"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	commentRepo := repository.NewCommentRepository()
	commentUsecase := usecase.NewCommentUsecase(commentRepo)

	commentService := grpcSvc.NewCommentService(commentUsecase)

	pb.RegisterCommentServiceServer(s, commentService)

	listen, err := net.Listen("tcp", ":3100")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening on port 3100")

	err = s.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
