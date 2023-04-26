package searchdelivery

import (
	"hesh/internal/pkg/domain"
	"hesh/internal/pkg/utils/log"
	"hesh/internal/pkg/utils/sanitizer"

	"strconv"

	"net/http"

	"github.com/mailru/easyjson"
)

func (handler *SearchHandler) SearchDiary (w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// categoryString := r.URL.Query().Get("category")
	// categories := strings.Split(categoryString, " ")
	queryParameters := r.URL.Query()
	userId, err := strconv.ParseUint(queryParameters.Get("vk_user_id"), 10, 64)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	text := queryParameters.Get("text")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	searchParams := new(domain.SearchDiaryRequest)
	searchParams.SetDefault()

	searchParams.Text = text

	sanitizer.SanitizeSearchDiaryParams(searchParams)

	diaryList, err := handler.SearchUsecase.SearchDiary(userId, *searchParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := easyjson.Marshal(diaryList)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(out)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
