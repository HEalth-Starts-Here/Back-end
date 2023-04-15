package commentusecase

import (
	"hesh/internal/pkg/domain"
	"strings"
)

type CommentUsecase struct {
	commentRepo domain.CommentRepository
}

func trimTitle(title *string) {
	*title = strings.Trim(*title, " ")
}

// Try to import with this function
func InitCommentUsc(сr domain.CommentRepository) domain.CommentUsecase {
	return &CommentUsecase{
		commentRepo: сr,
	}
}

func (cu CommentUsecase) CheckUserRole(userId uint64) (bool, bool, error) {
	userExist, isMedic, err := cu.commentRepo.CheckUserRole(userId)
	if err != nil {
		return false, false, err
	}
	return userExist, isMedic, nil
}

func (cu CommentUsecase) CreateComment(diaryId uint64, userId uint64, commentCreateRequest domain.BasicCommentInfo) (domain.CommentCreateResponse, error) {
	if !commentCreateRequest.IsValid() {
		return domain.CommentCreateResponse{}, domain.Err.ErrObj.InvalidTitleOrDescription
	}
	isExist, isMedic, err := cu.CheckUserRole(userId)
	if err != nil {
		return domain.CommentCreateResponse{}, err
	}
	if !isExist {
		return domain.CommentCreateResponse{}, domain.Err.ErrObj.UserNotExist
	}

	DiaryCreateResponse, err := cu.commentRepo.CreateComment(diaryId, isMedic, commentCreateRequest)
	if err != nil {
		return domain.CommentCreateResponse{}, err
	}
	return DiaryCreateResponse, nil
}
