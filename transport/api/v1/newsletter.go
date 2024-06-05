package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"vse-go-newsletter-api/pkg/id"
	"vse-go-newsletter-api/service/errors"
	"vse-go-newsletter-api/service/model"
	transportModel "vse-go-newsletter-api/transport/api/v1/model"
	"vse-go-newsletter-api/transport/util"

	"github.com/go-chi/chi"
)

func (h *Handler) CreateNewsletter(w http.ResponseWriter, r *http.Request) {
	var newsletter transportModel.NewsLetter

	if err := json.NewDecoder(r.Body).Decode(&newsletter); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if newsletter.Name == "" || newsletter.Description == "" {
		http.Error(w,
			fmt.Sprintf("Invalid request body, name or description can´t be nil! \nName: %s \nDescription: %s",
				newsletter.Name,
				newsletter.Description),
			http.StatusBadRequest)
		return
	}

	ctx, userId, isUnauthenticated := getUserId(w, r)
	if isUnauthenticated {
		return
	}

	var serviceNewsletter = model.Newsletter{Name: newsletter.Name, Description: newsletter.Description, OwnerId: userId}

	created, err := h.service.CreateNewsletter(ctx, serviceNewsletter)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusCreated, created)
}

func (h *Handler) GetNewsletter(w http.ResponseWriter, r *http.Request) {
	newsletterID := getNewsletterId(w, r)

	newsletter, err := h.service.GetNewsletter(r.Context(), newsletterID)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, newsletter)
}

func (h *Handler) ListNewsletters(w http.ResponseWriter, r *http.Request) {
	newsletters, err := h.service.ListNewsletters(r.Context())
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, newsletters)
}

func (h *Handler) UpdateNewsletter(w http.ResponseWriter, r *http.Request) {
	newsletterID := getNewsletterId(w, r)

	var newsletter transportModel.NewsLetter

	if err := json.NewDecoder(r.Body).Decode(&newsletter); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx, userId, isUnauthenticated := getUserId(w, r)
	if isUnauthenticated {
		return
	}

	var serviceNewsletter = model.Newsletter{ID: newsletterID, Name: newsletter.Name, Description: newsletter.Description, OwnerId: userId}

	updated, err := h.service.UpdateNewsletter(ctx, serviceNewsletter, userId)

	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, updated)
}

func (h *Handler) DeleteNewsletter(w http.ResponseWriter, r *http.Request) {
	newsletterID := getNewsletterId(w, r)

	ctx, userId, isUnauthenticated := getUserId(w, r)
	if isUnauthenticated {
		return
	}

	deleted, err := h.service.DeleteNewsletter(ctx, newsletterID, userId)

	if err != nil {
		if err.Error() == errors.ErrNotFound.Error() {
			util.WriteErrResponse(w, http.StatusNotFound, err)
			return
		} else if err.Error() == errors.ErrForbidden.Error() {
			util.WriteErrResponse(w, http.StatusForbidden, err)
			return
		} else {
			util.WriteErrResponse(w, http.StatusInternalServerError, err)
			return
		}
	}

	util.WriteResponse(w, http.StatusNoContent, deleted)
}

func getUserId(w http.ResponseWriter, r *http.Request) (context.Context, string, bool) {
	ctx := r.Context()
	userData := ctx.Value("user").(map[string]interface{})
	if userData == nil {
		http.Error(w, "User not logged in!", http.StatusForbidden)
		return nil, "", true
	}
	userId := userData["userID"].(string)
	return ctx, userId, false
}

func getNewsletterId(w http.ResponseWriter, r *http.Request) id.Newsletter {
	var newsletterID id.Newsletter
	if err := newsletterID.FromString(chi.URLParam(r, "id")); err != nil {
		http.Error(w, "invalid newsletter ID", http.StatusBadRequest)
	}
	return newsletterID
}
