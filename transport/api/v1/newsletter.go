package v1

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"vse-go-newsletter-api/pkg/id"
	"vse-go-newsletter-api/service/model"
	model2 "vse-go-newsletter-api/transport/api/v1/model"
	"vse-go-newsletter-api/transport/util"

	"github.com/go-chi/chi"
)

func (h *Handler) CreateNewsletter(w http.ResponseWriter, r *http.Request) {
	var newsletter model2.NewsLetter

	if err := json.NewDecoder(r.Body).Decode(&newsletter); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	ctx := r.Context()
	userData := ctx.Value("user").(map[string]interface{})

	if userData == nil {
		http.Error(w, "User not logged in!", http.StatusForbidden)
		return
	}
	userId := userData["userID"].(string)

	created, err := h.service.CreateNewsletter(ctx, newsletter.Name, newsletter.Description, userId)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
	}
	util.WriteResponse(w, http.StatusOK, created)
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
	var newsletter model2.NewsLetter
	if err := json.NewDecoder(r.Body).Decode(&newsletter); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	//TODO update only if the updater is owner
	ctx := r.Context()
	userData := ctx.Value("user").(map[string]interface{})

	var serviceNewsletter = model.Newsletter{ID: newsletterID, Name: newsletter.Name, Description: newsletter.Description,
		OwnerId: userData["user_id"].(string)}

	updatedNewsletter, err := h.service.UpdateNewsletter(r.Context(), serviceNewsletter)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, updatedNewsletter)
}

func (h *Handler) DeleteNewsletter(w http.ResponseWriter, r *http.Request) {
	var newsletterID id.Newsletter
	if err := newsletterID.FromString(chi.URLParam(r, "id")); err != nil {
		http.Error(w, "invalid newsletter ID", http.StatusBadRequest)
		return
	}

	// decode JSON body request
	var newsletter model2.NewsLetter
	if err := json.NewDecoder(r.Body).Decode(&newsletter); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	userData := ctx.Value("user").(map[string]interface{})

	var serviceNewsletter = model.Newsletter{ID: newsletterID, Name: newsletter.Name, Description: newsletter.Description,
		OwnerId: userData["user_id"].(string)}

	err := h.service.DeleteNewsletter(r.Context(), serviceNewsletter)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, newsletterID)
}
