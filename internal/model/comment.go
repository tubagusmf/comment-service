package model

type Comment struct {
	Id        string `json:"id"`
	StoryId   string `json:"story_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CommentUsecase interface {
	FindComments(storyId string, from string) ([]Comment, error)
}

type CommentRepository interface {
	FindComments(storyId string) ([]Comment, error)
	FindCommentFromDocument(storyId string) ([]Comment, error)
}
