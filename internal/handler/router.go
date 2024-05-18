package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/city", h.GetMap).Methods(http.MethodPost)

	return r
}
