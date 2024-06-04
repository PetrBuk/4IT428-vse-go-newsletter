package v1

import (
	"errors"
	"net/http"
	"vse-go-newsletter-api/service/model"
	"vse-go-newsletter-api/transport/util"
)

func (h *Handler) SubscribeNewsletter(w http.ResponseWriter, r *http.Request) {
	newsletterId := getNewsletterId(w, r)
	email := getEmail(w, r)

	if email != "" {
		subscription, err := h.service.SubscribeNewsletter(r.Context(), newsletterId, email)
		if err != nil {
			util.WriteErrResponse(w, http.StatusInternalServerError, err)
			return
		}
		util.WriteResponse(w, http.StatusOK, model.Subscription{ID: subscription.ID, Email: subscription.Email, NewsletterId: subscription.NewsletterId, CreatedAt: subscription.CreatedAt, IsConfirmed: subscription.IsConfirmed})

	}
}

func (h *Handler) UnsubscribeNewsletter(w http.ResponseWriter, r *http.Request) {
	newsletterId := getNewsletterId(w, r)

	email := getEmail(w, r)
	if email != "" {
		unsubscription, err := h.service.UnsubscribeNewsletter(r.Context(), newsletterId, email)
		if err != nil {
			util.WriteErrResponse(w, http.StatusInternalServerError, err)
			return
		}
		util.WriteResponse(w, http.StatusOK, unsubscription)
	}
}

func (h *Handler) ConfirmSubscription(w http.ResponseWriter, r *http.Request) {
	newsletterId := getNewsletterId(w, r)
	email := getEmail(w, r)
	if email != "" {
		subscription, err := h.service.ConfirmSubscription(r.Context(), newsletterId, email)
		if err != nil {
			util.WriteErrResponse(w, http.StatusInternalServerError, err)
			return
		}
		util.WriteResponse(w, http.StatusOK, model.Subscription{ID: subscription.ID, Email: subscription.Email, NewsletterId: subscription.NewsletterId, CreatedAt: subscription.CreatedAt, IsConfirmed: subscription.IsConfirmed})

	}

}

func getEmail(w http.ResponseWriter, r *http.Request) string {
	var email string = r.URL.Query().Get("email")

	if email == "" {
		util.WriteErrResponse(w, http.StatusBadRequest, errors.New("email is required"))
		return ""
	}

	return email
}
