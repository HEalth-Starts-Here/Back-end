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

func (nu NoteUsecase) CreateNote(medicId uint64, isMedicRecord bool, recordId uint64, noteCreateRequest *domain.BasicNoteInfo) (domain.NoteCreateResponse, error) {
	if !noteCreateRequest.IsValid() {
		return domain.NoteCreateResponse{}, domain.Err.ErrObj.InvalidText
	}
	// isExist, isMedic, err := nu.CheckUserRole(userId)
	// if err != nil {
	// 	return domain.CommentCreateResponse{}, err
	// }
	// if !isExist {
	// 	return domain.CommentCreateResponse{}, domain.Err.ErrObj.UserNotExist
	// }

	DiaryCreateResponse, err := nu.NoteRepo.CreateNote(isMedicRecord, recordId, *noteCreateRequest)
	if err != nil {
		return domain.NoteCreateResponse{}, err
	}
	return DiaryCreateResponse, nil
}

func (nu NoteUsecase) GetNote(medicId uint64, isMedicRecord bool, recordId uint64) (domain.GetNoteResponse, error) {
	// TODO check is user have access to this diary
	GetNoteResponse, err := nu.NoteRepo.GetNote(isMedicRecord, recordId)
	if err != nil {
		return domain.GetNoteResponse{}, err
	}
	fmt.Printf("isMedicRecord: %v\n", isMedicRecord)
	return GetNoteResponse, nil
}

func (nu NoteUsecase) DeleteNote(medicId uint64, isMedicRecord bool, commentId uint64) error {
	err := nu.NoteRepo.DeleteNote(commentId)
	if err != nil {
		return err
	}

	return nil
}
