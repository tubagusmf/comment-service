package grpc

import (
	"context"
	"errors"

	pb "github.com/tubagusmf/comment-service/pb/comment"
)

func (s *CommentService) FindComments(ctx context.Context, in *pb.CommentRequest) (*pb.CommentList, error) {
	if in.StoryId == "" {
		return nil, errors.New("storyId is required")
	}
	comments, err := s.commentUsecase.FindComments(in.StoryId, "postgres")
	if err != nil {
		return nil, err
	}

	var commentPbs []*pb.Comment

	for _, comment := range comments {
		commentPbs = append(commentPbs, &pb.Comment{
			Id:        comment.Id,
			StoryId:   comment.StoryId,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		})
	}
	return &pb.CommentList{Comments: commentPbs}, nil
}
