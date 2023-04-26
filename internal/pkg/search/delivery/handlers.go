package searchdelivery

import (
	"hesh/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type SearchHandler struct {
	SearchUsecase domain.SearchUsecase
}

func SetSearchHandlers(router *mux.Router, nu domain.SearchUsecase) {
	handler := &SearchHandler{
		SearchUsecase: nu,
	}
	router.HandleFunc(SearchDiaryUrl, handler.SearchDiary).Methods("GET", "OPTIONS")
}
