package v1

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"vse-go-newsletter-api/pkg/id"
	"vse-go-newsletter-api/service/model"
	"vse-go-newsletter-api/transport/util"

	"github.com/go-chi/chi"
)

func (h *Handler) CreateNewsletter(w http.ResponseWriter, r *http.Request) {
	panic("not implemented - CreateNewsletter")
}

func (h *Handler) GetNewsletter(w http.ResponseWriter, r *http.Request) {
	var newsletterID id.Newsletter
	if err := newsletterID.FromString(chi.URLParam(r, "id")); err != nil {
		http.Error(w, "invalid newsletter ID", http.StatusBadRequest)
	}

	newsletter, err := h.service.GetNewsletter(r.Context(), newsletterID)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, newsletter)
}

func (h *Handler) ListNewsletters(w http.ResponseWriter, r *http.Request) {
	slog.Info("getting list newsletters")
	newsletters, err := h.service.ListNewsletters(r.Context())
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, newsletters)
}

func (h *Handler) UpdateNewsletter(w http.ResponseWriter, r *http.Request) {
	var newsletterID id.Newsletter
	if err := newsletterID.FromString(chi.URLParam(r, "id")); err != nil {
		http.Error(w, "invalid newsletter ID", http.StatusBadRequest)
		return
	}

	// decode JSON body request
	var newsletter model.Newsletter
	if err := json.NewDecoder(r.Body).Decode(&newsletter); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	newsletter.ID = newsletterID

	//TODO update only if the updater is owner

	updatedNewsletter, err := h.service.UpdateNewsletter(r.Context(), newsletterID, newsletter)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, updatedNewsletter)
}

func (h *Handler) DeleteNewsletter(w http.ResponseWriter, r *http.Request) {
	panic("not implemented - DeleteNewsletter")
}
