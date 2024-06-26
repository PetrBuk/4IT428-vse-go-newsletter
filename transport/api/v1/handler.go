package v1

import (
	"vse-go-newsletter-api/transport/middleware"

	"github.com/go-chi/chi"
)

type Handler struct {
	*chi.Mux

	service       RouteService
	authenticator middleware.Authenticator
}

func NewHandler(
	service RouteService,
	authenticator middleware.Authenticator,
) *Handler {
	h := &Handler{
		service:       service,
		authenticator: authenticator,
	}
	h.initRouter()
	return h
}

func (h *Handler) initRouter() {
	r := chi.NewRouter()

	authenticate := middleware.NewAutheticate(h.authenticator)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", h.Login)
		r.Post("/register", h.Register)
		r.Post("/verify", h.Verify)
		r.Post("/refresh", h.RefreshToken)
		r.With(authenticate).Post("/change-password", h.ChangePassword)
	})

	r.Route("/newsletters", func(r chi.Router) {
		r.With(authenticate).Get("/", h.ListNewsletters)
		r.With(authenticate).Post("/", h.CreateNewsletter)
		r.With(authenticate).Get("/{id}", h.GetNewsletter)
		r.With(authenticate).Put("/{id}", h.UpdateNewsletter)
		r.With(authenticate).Delete("/{id}", h.DeleteNewsletter)
		r.Get("/{id}/subscribe", h.SubscribeNewsletter)
		r.Get("/{id}/unsubscribe", h.UnsubscribeNewsletter)
		r.Get("/{id}/confirm", h.ConfirmSubscription)
	})

	r.Route("/posts", func(r chi.Router) {
		r.With(authenticate).Get("/", h.ListPosts)
		r.With(authenticate).Post("/", h.CreatePost)
		r.With(authenticate).Get("/{id}", h.GetPost)
		r.With(authenticate).Put("/{id}", h.UpdatePost)
		r.With(authenticate).Delete("/{id}", h.DeletePost)
		r.With(authenticate).Put("/{id}/publish", h.PublishPost)
	})

	h.Mux = r
}
