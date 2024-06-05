package v1

import (
	"errors"
	"html/template"
	"net/http"
	"os"
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
		_, err := h.service.UnsubscribeNewsletter(r.Context(), newsletterId, email)
		if err != nil {
			util.WriteErrResponse(w, http.StatusInternalServerError, err)
			return
		}

		tmpl, err := template.ParseFiles("templates/pages/unsubscribe_success.html")
		if err != nil {
				util.WriteErrResponse(w, http.StatusInternalServerError, err)
				return
		}

		if err := tmpl.Execute(w, nil); err != nil {
      util.WriteErrResponse(w, http.StatusInternalServerError, err)
			return
    }
	}
}

func (h *Handler) ConfirmSubscription(w http.ResponseWriter, r *http.Request) {
	newsletterId := getNewsletterId(w, r)
	email := getEmail(w, r)
	baseUrl := os.Getenv("BASE_URL")

	if email != "" {
		_, err := h.service.ConfirmSubscription(r.Context(), newsletterId, email)

		if err != nil {
			util.WriteErrResponse(w, http.StatusInternalServerError, err)
			return
		}

		tmpl, err := template.ParseFiles("templates/pages/confirm_success.html")
		if err != nil {
				util.WriteErrResponse(w, http.StatusInternalServerError, err)
				return
		}

		templateData := struct {
		SubscriberEmail string
		UnsubscribeLink string
	}{
		SubscriberEmail: email,
		UnsubscribeLink: baseUrl + "/api/v1/newsletters/" + newsletterId.String() + "/unsubscribe?email=" + email,
	}

		if err := tmpl.Execute(w, templateData); err != nil {
      util.WriteErrResponse(w, http.StatusInternalServerError, err)
			return
    }
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
