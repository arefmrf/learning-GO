package controller

import (
	"encoding/json"
	"net/http"
	"prj/internal/modules/item/service"
	"strconv"
)

type Controller struct {
	service service.ItemServiceInterface
}

func NewController() *Controller {
	return &Controller{service: service.NewService()}
}

func (h *Controller) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()
	q := r.URL.Query()
	limit, _ := strconv.Atoi(q.Get("limit"))
	offset, _ := strconv.Atoi(q.Get("offset"))
	if limit == 0 {
		limit = 10
	}

	items, err := h.service.List(ctx, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}
