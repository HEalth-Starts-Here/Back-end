package notedelivery

import (
	"hesh/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type NoteHandler struct {
	NoteUsecase domain.NoteUsecase
}

func SetNoteHandlers(router *mux.Router, nu domain.NoteUsecase) {
	handler := &NoteHandler{
		NoteUsecase: nu,
	}
	router.HandleFunc(GetNoteUrl, handler.GetNote).Methods("GET", "OPTIONS")
	router.HandleFunc(CreateNoteUrl, handler.CreateNote).Methods("POST", "OPTIONS")
	// router.HandleFunc(DeleteNoteUrl, handler.DeleteNote).Methods("POST", "OPTIONS")
}
