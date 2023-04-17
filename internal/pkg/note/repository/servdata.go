package noterepository

import (
	"fmt"
	"hesh/internal/pkg/database"
	"hesh/internal/pkg/domain"
	"hesh/internal/pkg/utils/cast"
	"hesh/internal/pkg/utils/log"
)

type dbnoterepository struct {
	dbm *database.DBManager
}

func InitNoteRep(manager *database.DBManager) domain.NoteRepository {
	return &dbnoterepository{
		dbm: manager,
	}
}

// func (er *dbnoterepository) CheckUserRole(userId uint64) (bool, bool, error) {
// 	query := queryCheckUserRole
// 	resp, err := er.dbm.Query(query,
// 		userId)
// 	if err != nil {
// 		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
// 		log.Error(err)
// 		return false, false, err
// 	}
// 	if len(resp) == 0 {
// 		return false, false, nil
// 	}
// 	return true, cast.ToBool(resp[0][0]), nil
// }

// func (er *dbnoterepository) CreateNote(medicId uint64, isMedic bool, commentrequest domain.BasicCommentInfo) (domain.CommentCreateResponse, error) {
// 	query := queryCreateComment
// 	resp, err := er.dbm.Query(query,
// 		authorIsMedic,
// 		time.Now().Format("2006.01.02 15:04:05"),
// 		diaryId,
// 		commentrequest.Text)
// 	if err != nil {
// 		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
// 		log.Error(err)
// 		return domain.CommentCreateResponse{}, err
// 	}

// 	return domain.CommentCreateResponse{
// 		CommentInListInfo: domain.CommentInListInfo{
// 			Id:					cast.ToUint64(resp[0][0]),
// 			AuthorIsMedic:      cast.ToBool(resp[0][1]),
// 			IsReaded:			cast.ToBool(resp[0][2]),
// 			CreatingDate: 		cast.TimeToStr(cast.ToTime(resp[0][3]), true),
// 			BasicCommentInfo: domain.BasicCommentInfo{
// 				Text:       cast.ToString(resp[0][4]),
// 			},
// 		},
// 		DiaryId:			cast.ToUint64(resp[0][5]),
// 	}, nil
// }

func (cr *dbnoterepository) GetNote(isMedicRecord bool, recordId uint64) (domain.GetNoteResponse, error) {
	var resp []database.DBbyterow
	var err error

	query := queryGetNote3
	var userRecord string
	if isMedicRecord {
		userRecord = "medicrecordid"
	} else {
		userRecord = "patientrecordid"
	}
	// resp, err = cr.dbm.Query(query, userRecord, strconv.Itoa(int(recordId)))
	// resp, err = cr.dbm.Query(query, strconv.Itoa(int(recordId)))
	resp, err = cr.dbm.Query(query, userRecord)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.GetNoteResponse{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		println("len(resp) == 0")
		return domain.GetNoteResponse{}, nil
	}
	notes := make([]domain.NoteInListInfo, 0)
	for i := range resp {
		notes = append(notes, domain.NoteInListInfo{
			Id:           		cast.ToUint64(resp[i][1]),
			CreatingDate:    	cast.TimeToStr(cast.ToTime(resp[i][2]), true),
			BasicNoteInfo: domain.BasicNoteInfo{
				Text: cast.ToString(resp[i][3]),
			},
		})
	}

	out := domain.GetNoteResponse{
		RecordId: cast.ToUint64(resp[0][0]),
		IsMedicRecord: isMedicRecord,
		NoteList: notes,
	}

	return out, nil
}

// func (er *dbcommentrepository) DeleteComment(commentId uint64) error {
// 	query := queryDeleteComment
// 	_, err := er.dbm.Query(query, commentId)
// 	if err != nil {
// 		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
// 		log.Error(err)
// 	}
// 	return err
// }
