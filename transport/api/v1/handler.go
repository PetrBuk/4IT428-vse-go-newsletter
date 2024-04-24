package v1

import (
	"vse-go-newsletter-api/transport/middleware"

	"github.com/go-chi/chi"
)

type Handler struct {
	*chi.Mux

	service Service
	authenticator middleware.Authenticator
}

func NewHandler(
	service Service,
	authenticator middleware.Authenticator,
) *Handler {
	h := &Handler{
		service: service,
		authenticator: authenticator,
	}
	h.initRouter()
	return h
}

func (h *Handler) initRouter() {
	r := chi.NewRouter()

	authenticate := middleware.NewAutheticate(h.authenticator)

	r.Route("/newsletters", func(r chi.Router) {
		r.With(authenticate).Get("/", h.ListNewsletters)
		r.With(authenticate).Post("/", h.CreateNewsletter)
		r.With(authenticate).Get("/{id}", h.GetNewsletter)
		r.With(authenticate).Put("/{id}", h.UpdateNewsletter)
		r.With(authenticate).Delete("/{id}", h.DeleteNewsletter)
	})

	h.Mux = r
}
