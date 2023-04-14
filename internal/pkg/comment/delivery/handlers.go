package commentdelivery

import (
	"hesh/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type CommentHandler struct {
	CommentUsecase domain.CommentUsecase
}

func SetCommentHandlers(router *mux.Router, cu domain.CommentUsecase) {
	handler := &CommentHandler{
		CommentUsecase: cu,
	}
	router.HandleFunc(CreateCommentUrl, handler.CreateComment).Methods("POST", "OPTIONS")

}
