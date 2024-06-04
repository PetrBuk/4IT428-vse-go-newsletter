package v1

import (
	"net/http"
	"vse-go-newsletter-api/service/model"
	"vse-go-newsletter-api/transport/util"
)

func (h *Handler) SubscribeNewsletter(w http.ResponseWriter, r *http.Request) {
	newsletterId := getNewsletterId(w, r)

	ctx, userId, isUnauthenticated := getUserId(w, r)
	if isUnauthenticated {
		return
	}

	subscription, err := h.service.SubscribeNewsletter(ctx, newsletterId, userId)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	util.WriteResponse(w, http.StatusOK, model.Subscription{ID: subscription.ID, UserId: subscription.UserId, NewsletterId: subscription.NewsletterId, CreatedAt: subscription.CreatedAt, IsConfirmed: subscription.IsConfirmed})
}

func (h *Handler) UnsubscribeNewsletter(w http.ResponseWriter, r *http.Request) {
	newsletterId := getNewsletterId(w, r)

	ctx, userId, isUnauthenticated := getUserId(w, r)
	if isUnauthenticated {
		return
	}

	unsubscription, err := h.service.UnsubscribeNewsletter(ctx, newsletterId, userId)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	util.WriteResponse(w, http.StatusOK, unsubscription)
}

func (h *Handler) ConfirmSubscription(w http.ResponseWriter, r *http.Request) {
	newsletterId := getNewsletterId(w, r)

	ctx, userId, isUnauthenticated := getUserId(w, r)
	if isUnauthenticated {
		return
	}

	subscription, err := h.service.ConfirmSubscription(ctx, newsletterId, userId)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	util.WriteResponse(w, http.StatusOK, model.Subscription{ID: subscription.ID, UserId: subscription.UserId, NewsletterId: subscription.NewsletterId, CreatedAt: subscription.CreatedAt, IsConfirmed: subscription.IsConfirmed})
}
