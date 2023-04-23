package noterepository

import (
	"hesh/internal/pkg/database"
	"hesh/internal/pkg/domain"
	"hesh/internal/pkg/utils/cast"
	"hesh/internal/pkg/utils/log"
	"strings"
	"time"
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

func (cr *dbnoterepository) GetNote(isMedicRecord bool, recordId uint64) (domain.GetNoteResponse, error) {
	var resp []database.DBbyterow
	var err error

	// query := queryGetNote3
	var query strings.Builder

	var userRecord string
	if isMedicRecord {
		userRecord = "medicrecordid"
	} else {
		userRecord = "patientrecordid"
	}
	query.Write([]byte(queryGetNoteFirstPart))
	query.Write([]byte(userRecord))
	query.Write([]byte(queryGetNoteSecondPart))
	query.Write([]byte(userRecord))
	query.Write([]byte(queryGetNoteThirdPart))

	resp, err = cr.dbm.Query(query.String(), recordId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query.String())
		log.Error(err)
		return domain.GetNoteResponse{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		return domain.GetNoteResponse{}, nil
	}
	notes := make([]domain.NoteInListInfo, 0)
	for i := range resp {
		notes = append(notes, domain.NoteInListInfo{
			Id:           cast.ToUint64(resp[i][1]),
			CreatingDate: cast.TimeToStr(cast.ToTime(resp[i][2]), true),
			BasicNoteInfo: domain.BasicNoteInfo{
				Text: cast.ToString(resp[i][3]),
			},
		})
	}

	out := domain.GetNoteResponse{
		RecordId:      cast.ToUint64(resp[0][0]),
		IsMedicRecord: isMedicRecord,
		NoteList:      notes,
	}

	return out, nil
}

func (er *dbnoterepository) CreateNote(isMedicRecord bool, recordId uint64, noteRequest domain.BasicNoteInfo) (domain.NoteCreateResponse, error) {
	var query strings.Builder
	var userRecord string
	if isMedicRecord {
		userRecord = "medicrecordid"
	} else {
		userRecord = "patientrecordid"
	}
	query.Write([]byte(queryCreateNoteFirstPart))
	query.Write([]byte(userRecord))
	query.Write([]byte(queryCreateNoteSecondPart))
	query.Write([]byte(userRecord))
	query.Write([]byte(queryCreateNoteThirdPart))

	resp, err := er.dbm.Query(query.String(),
		recordId,
		isMedicRecord,
		time.Now().Format("2006.01.02 15:04:05"),
		noteRequest.Text)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query.String())
		log.Error(err)
		return domain.NoteCreateResponse{}, err
	}
	return domain.NoteCreateResponse{
		NoteInListInfo: domain.NoteInListInfo{
			Id:           cast.ToUint64(resp[0][0]),
			CreatingDate: cast.TimeToStr(cast.ToTime(resp[0][3]), true),
			BasicNoteInfo: domain.BasicNoteInfo{
				Text: cast.ToString(resp[0][4]),
			},
		},
		RecordId:      cast.ToUint64(resp[0][1]),
		IsMedicRecord: cast.ToBool(resp[0][2]),
	}, nil
}

func (er *dbnoterepository) DeleteNote(noteId uint64) error {
	query := queryDeleteNote
	_, err := er.dbm.Query(query,
		noteId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
	}
	return err
}
