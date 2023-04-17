package noteusecase

import (
	"fmt"
	"hesh/internal/pkg/domain"
)

type NoteUsecase struct {
	NoteRepo domain.NoteRepository
}

func InitNoteUsc(nr domain.NoteRepository) domain.NoteUsecase {
	return &NoteUsecase{
		NoteRepo: nr,
	}
}

// func (cu CommentUsecase) CheckUserRole(userId uint64) (bool, bool, error) {
// 	userExist, isMedic, err := cu.commentRepo.CheckUserRole(userId)
// 	if err != nil {
// 		return false, false, err
// 	}
// 	return userExist, isMedic, nil
// }

// func (nu NoteUsecase) CreateNote(diaryId uint64, userId uint64, commentCreateRequest domain.BasicCommentInfo) (domain.CommentCreateResponse, error) {
// 	if !commentCreateRequest.IsValid() {
// 		return domain.CommentCreateResponse{}, domain.Err.ErrObj.InvalidTitleOrDescription
// 	}
// 	// svs := setter.Services{}
// 	// isExist, isMedic, err := diaryrepository.InitDiaryRep(svs.Diary.Db).GetUserRole(userId)
// 	isExist, isMedic, err := cu.CheckUserRole(userId) 
// 	if err != nil {
// 		return domain.CommentCreateResponse{}, err
// 	}
// 	if !isExist {
// 		return domain.CommentCreateResponse{}, domain.Err.ErrObj.UserNotExist
// 	}

// 	DiaryCreateResponse, err := cu.commentRepo.CreateComment(diaryId, isMedic, commentCreateRequest)
// 	if err != nil {
// 		return domain.CommentCreateResponse{}, err
// 	}
// 	return DiaryCreateResponse, nil
// }

func (nu NoteUsecase) GetNote (medicId uint64, isMedicRecord bool, recordId uint64) (domain.GetNoteResponse, error) {
	// TODO check is user have access to this diary
	GetNoteResponse, err := nu.NoteRepo.GetNote(isMedicRecord, recordId)
	if err != nil {
		return domain.GetNoteResponse{}, err
	}
	fmt.Printf("isMedicRecord: %v\n", isMedicRecord)
	return GetNoteResponse, nil
}

// func (du CommentUsecase) DeleteComment(userId, commentId uint64) error {
// 	err := du.commentRepo.DeleteComment(commentId)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
