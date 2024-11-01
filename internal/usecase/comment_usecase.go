package usecase

import "github.com/tubagusmf/comment-service/internal/model"

type commentUsecase struct {
	commentRepo model.CommentRepository
}

func NewCommentUsecase(cr model.CommentRepository) model.CommentUsecase {
	return &commentUsecase{commentRepo: cr}
}

func (cu *commentUsecase) FindComments(storyId string, from string) ([]model.Comment, error) {
	return cu.commentRepo.FindComments(storyId)
}
