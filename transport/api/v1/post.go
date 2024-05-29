package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	model2 "vse-go-newsletter-api/transport/api/v1/model"
	"vse-go-newsletter-api/transport/util"
)

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post model2.Post

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if post.Title == "" || post.Content == "" {
		http.Error(w,
			fmt.Sprintf("Invalid request body, title or content can´t be nil! \nTitle: %s \nContent: %s",
				post.Title,
				post.Content),
			http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	userData := ctx.Value("user").(map[string]interface{})

	if userData == nil {
		http.Error(w, "User not logged in!", http.StatusForbidden)
		return
	}

	created, err := h.service.CreatePost(ctx, post.Title, post.Content, post.NewsletterId)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
	}
	util.WriteResponse(w, http.StatusOK, created)
}

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	//var newsletterID id.Newsletter
	//if err := newsletterID.FromString(chi.URLParam(r, "id")); err != nil {
	//	http.Error(w, "invalid newsletter ID", http.StatusBadRequest)
	//}
	//
	//newsletter, err := h.service.GetNewsletter(r.Context(), newsletterID)
	//if err != nil {
	//	util.WriteErrResponse(w, http.StatusInternalServerError, err)
	//	return
	//}
	//
	//util.WriteResponse(w, http.StatusOK, newsletter)
}

func (h *Handler) ListPosts(w http.ResponseWriter, r *http.Request) {
	//slog.Info("getting list newsletters")
	//newsletters, err := h.service.ListNewsletters(r.Context())
	//if err != nil {
	//	util.WriteErrResponse(w, http.StatusInternalServerError, err)
	//	return
	//}
	//
	//util.WriteResponse(w, http.StatusOK, newsletters)
}

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	//var newsletterID id.Newsletter
	//if err := newsletterID.FromString(chi.URLParam(r, "id")); err != nil {
	//	http.Error(w, "invalid newsletter ID", http.StatusBadRequest)
	//	return
	//}
	//
	//// decode JSON body request
	//var newsletter model2.NewsLetter
	//if err := json.NewDecoder(r.Body).Decode(&newsletter); err != nil {
	//	http.Error(w, "invalid request body", http.StatusBadRequest)
	//	return
	//}
	//
	////TODO update only if the updater is owner
	//ctx := r.Context()
	//userData := ctx.Value("user").(map[string]interface{})
	//
	//var serviceNewsletter = model.Newsletter{ID: newsletterID, Name: newsletter.Name, Description: newsletter.Description,
	//	OwnerId: userData["user_id"].(string)}
	//
	//updatedNewsletter, err := h.service.UpdateNewsletter(r.Context(), serviceNewsletter)
	//if err != nil {
	//	util.WriteErrResponse(w, http.StatusInternalServerError, err)
	//	return
	//}
	//
	//util.WriteResponse(w, http.StatusOK, updatedNewsletter)
}

func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	//var newsletterID id.Newsletter
	//if err := newsletterID.FromString(chi.URLParam(r, "id")); err != nil {
	//	http.Error(w, "invalid newsletter ID", http.StatusBadRequest)
	//	return
	//}
	//
	//// decode JSON body request
	//var newsletter model2.NewsLetter
	//if err := json.NewDecoder(r.Body).Decode(&newsletter); err != nil {
	//	http.Error(w, "invalid request body", http.StatusBadRequest)
	//	return
	//}
	//
	//ctx := r.Context()
	//userData := ctx.Value("user").(map[string]interface{})
	//
	//var serviceNewsletter = model.Newsletter{ID: newsletterID, Name: newsletter.Name, Description: newsletter.Description,
	//	OwnerId: userData["user_id"].(string)}
	//
	//err := h.service.DeleteNewsletter(r.Context(), serviceNewsletter)
	//if err != nil {
	//	util.WriteErrResponse(w, http.StatusInternalServerError, err)
	//	return
	//}
	//
	//util.WriteResponse(w, http.StatusOK, newsletterID)
}
