package grpc

import (
	"github.com/tubagusmf/comment-service/internal/model"
	pb "github.com/tubagusmf/comment-service/pb/comment"
)

type CommentService struct {
	pb.UnimplementedCommentServiceServer
	commentUsecase model.CommentUsecase
}

func NewCommentService(commentUsecase model.CommentUsecase) *CommentService {
	return &CommentService{commentUsecase: commentUsecase}
}
