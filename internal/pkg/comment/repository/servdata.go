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
		CommentInListInfo: domain.CommentInListInfo{
			Id:            cast.ToUint64(resp[0][0]),
			AuthorIsMedic: cast.ToBool(resp[0][1]),
			IsReaded:      cast.ToBool(resp[0][2]),
			CreatingDate:  cast.TimeToStr(cast.ToTime(resp[0][3]), true),
			BasicCommentInfo: domain.BasicCommentInfo{
				Text: cast.ToString(resp[0][4]),
			},
		},
		DiaryId: cast.ToUint64(resp[0][5]),
	}, nil
}

func (cr *dbcommentrepository) GetComment(diaryId uint64) (domain.GetCommentResponse, error) {
	var resp []database.DBbyterow
	var err error

	query := queryGetComment
	resp, err = cr.dbm.Query(query, diaryId)

	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.GetCommentResponse{}, domain.Err.ErrObj.InternalServer
	}

	comments := make([]domain.CommentInListInfo, 0)
	for i := range resp {
		comments = append(comments, domain.CommentInListInfo{
			Id:               cast.ToUint64(resp[i][0]),
			AuthorIsMedic:    cast.ToBool(resp[i][1]),
			IsReaded:         cast.ToBool(resp[i][2]),
			CreatingDate:     cast.TimeToStr(cast.ToTime(resp[0][3]), true),
			BasicCommentInfo: domain.BasicCommentInfo{Text: cast.ToString(resp[i][4])},
		})
	}
	var out domain.GetCommentResponse
	if len(resp) != 0 {
		out = domain.GetCommentResponse{
			DiaryId:     diaryId,
			CommentList: comments,
		}
	}

	return out, nil
}

func (er *dbcommentrepository) DeleteComment(commentId uint64) error {
	query := queryDeleteComment
	_, err := er.dbm.Query(query, commentId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
	}
	return err
}
