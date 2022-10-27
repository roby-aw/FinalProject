package interfaces

import (
	"final-project/dto"
)

type CommentRepository interface {
	CreateComment(comment *dto.CreateComment) (dto.CreateComment, error)
	GetComment() ([]dto.GetAllComment, error)
	UpdateComment(id uint, comment *dto.UpdateComment) (dto.UpdateComment, error)
	GetPhotoId(id uint) (dto.GetPhotoUser, error)
	DeleteComment(id uint) error
	GetCommentId(id uint) (dto.GetComment, error)
}

type CommentService interface {
	CreateComment(comment *dto.CreateComment) (dto.GetComment, error)
	GetComment() ([]dto.GetAllComment, error)
	UpdateComment(id uint, comment *dto.UpdateComment) (*dto.GetPhotoUser, error)
	GetPhotoId(id uint) (dto.GetPhotoUser, error)
	DeleteComment(id uint) error
	GetCommentId(id uint) (dto.GetComment, error)
}
