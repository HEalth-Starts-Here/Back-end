package commentrepository

import (
	"hesh/internal/pkg/database"
	"hesh/internal/pkg/domain"

	"hesh/internal/pkg/utils/cast"
	"hesh/internal/pkg/utils/log"

	"time"
)

type dbcommentrepository struct {
	dbm *database.DBManager
}

func InitCommentRep(manager *database.DBManager) domain.CommentRepository {
	return &dbcommentrepository{
		dbm: manager,
	}
}

func (er *dbcommentrepository) CheckUserRole(userId uint64) (bool, bool, error) {
	query := queryCheckUserRole
	resp, err := er.dbm.Query(query,
		userId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return false, false, err
	}
	if len(resp) == 0 {
		return false, false, nil
	}
	return true, cast.ToBool(resp[0][0]), nil
}

func (er *dbcommentrepository) CreateComment(diaryId uint64, authorIsMedic bool, commentrequest domain.BasicCommentInfo) (domain.CommentCreateResponse, error) {
	query := queryCreateComment
	resp, err := er.dbm.Query(query,
		authorIsMedic,
		time.Now().Format("2006.01.02 15:04:05"),
		diaryId,
		commentrequest.Text)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.CommentCreateResponse{}, err
	}

	return domain.CommentCreateResponse{
		Id:					cast.ToUint64(resp[0][0]),
		AuthorIsMedic:      cast.ToBool(resp[0][1]),
		IsReaded:			cast.ToBool(resp[0][2]),
		CreatingDate: 		cast.TimeToStr(cast.ToTime(resp[0][3]), true),
		DiaryId:			cast.ToUint64(resp[0][4]),
		BasicCommentInfo: domain.BasicCommentInfo{
			Text:       cast.ToString(resp[0][5]),
		},
	}, nil
}
