package commentusecase

import (
	"hesh/internal/pkg/domain"
	// setter "hesh/internal/pkg/utils/setter"
	// diaryrepository "hesh/internal/pkg/diary/repository"
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
	// svs := setter.Services{}
	// isExist, isMedic, err := diaryrepository.InitDiaryRep(svs.Diary.Db).GetUserRole(userId)
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

func (cu CommentUsecase) GetComment (userId uint64, diaryId uint64) (domain.GetCommentResponse, error) {
	// TODO check is user have access to this diary
	GetCommentResponse, err := cu.commentRepo.GetComment(diaryId)
	if err != nil {
		return domain.GetCommentResponse{}, err
	}
	return GetCommentResponse, nil
}
