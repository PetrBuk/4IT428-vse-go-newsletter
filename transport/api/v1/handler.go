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
		r.With(authenticate).Post("/add", h.CreateNewsletter)
		r.With(authenticate).Get("/{id}", h.GetNewsletter)
		r.With(authenticate).Put("/{id}", h.UpdateNewsletter)
		r.With(authenticate).Delete("/{id}", h.DeleteNewsletter)
	})

	r.Route("/post", func(r chi.Router) {
		r.With(authenticate).Get("/", h.ListPosts)
		r.With(authenticate).Post("/add", h.CreatePost)
		r.With(authenticate).Get("/{id}", h.GetPost)
		r.With(authenticate).Put("/{id}", h.UpdatePost)
		r.With(authenticate).Delete("/{id}", h.DeletePost)
	})

	h.Mux = r
}
