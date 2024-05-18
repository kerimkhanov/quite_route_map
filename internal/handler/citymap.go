package handler

import (
	"context"
	"net/http"
)

func (h *Handler) GetMap(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.log.Error()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	err = h.uc.GetMap(ctx)
	if err != nil {
		return
	}
}
