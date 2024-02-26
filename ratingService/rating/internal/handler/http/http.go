package http

import (
	"books_microservices/ratingService/rating/internal/controller"
	"books_microservices/ratingService/rating/internal/repository"
	"books_microservices/ratingService/rating/pkg/model"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	ctrl *controller.Controller
}

func New(ctrl *controller.Controller) *Handler {
	return &Handler{ctrl}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	recordID := model.RecordID(r.FormValue("id"))
	if recordID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	recordType := model.RecordType(r.FormValue("type"))
	if recordType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	switch r.Method {
	case http.MethodGet:
		v, err := h.ctrl.GetAggregatedRating(r.Context(), recordID, recordType)
		if err != nil && errors.Is(err, repository.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("Response encode error: %v\n", err)
		}

	case http.MethodPut:
		userId := model.UserID(r.FormValue("userId"))
		v, err := strconv.ParseFloat(r.FormValue("value"), 64)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := h.ctrl.PutRating(r.Context(), recordID, recordType, &model.Rating{UserID: userId, Value: model.RatingValue(v)}); err != nil {
			log.Printf("Repository put error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
