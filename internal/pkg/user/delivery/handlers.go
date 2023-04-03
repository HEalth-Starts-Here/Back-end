package userdelivery

import (
	"hesh/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func SetUserHandlers(router *mux.Router, uu domain.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: uu,
	}
	router.HandleFunc(UserInitUrl, handler.UserInit).Methods("POST", "OPTIONS")
}
