package v1

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"log/slog"
	"net/http"
	"vse-go-newsletter-api/service/model"
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

	ctx, userId, isUserIdNotOk := getUserId(w, r)
	if isUserIdNotOk {
		return
	}

	var servicePost = model.Post{Title: post.Title, Content: post.Content, NewsletterId: post.NewsletterId}

	created, err := h.service.CreatePost(ctx, servicePost, userId)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
	}
	util.WriteResponse(w, http.StatusOK, created)
}

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	postId, isPostIdNotOk := getPostId(w, r)
	if isPostIdNotOk {
		return
	}

	newsletter, err := h.service.GetPost(r.Context(), postId)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, newsletter)
}

func (h *Handler) ListPosts(w http.ResponseWriter, r *http.Request) {
	slog.Info("getting list posts")
	posts, err := h.service.ListPosts(r.Context())
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, posts)
}

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	var post model2.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	postId, isPostIdNotOk := getPostId(w, r)
	if isPostIdNotOk {
		return
	}
	ctx, userId, isUserIdNotOk := getUserId(w, r)
	if isUserIdNotOk {
		return
	}

	if post.Title == "" || post.Content == "" || post.NewsletterId.IsEmpty() {
		http.Error(w,
			fmt.Sprintf("Invalid request body, title, content or newsletterId can´t be nil! \nTitle: %s \nContent: %s \nNewsletterId: %s",
				post.Title,
				post.Content,
				post.NewsletterId),
			http.StatusBadRequest)
		return
	}

	var servicePost = model.Post{ID: postId, Title: post.Title, Content: post.Content, NewsletterId: post.NewsletterId}

	updated, err := h.service.UpdatePost(ctx, servicePost, userId)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
	}
	message := fmt.Sprintf("post updated successfully! ID: %s, Title: %s, Content: %s, NewsletterId: %s, CreatedAt: %s, UpdatedAt: %s",
		updated.ID, updated.Title, updated.Content, updated.NewsletterId, updated.CreatedAt, updated.UpdatedAt)
	util.WriteResponse(w, http.StatusOK, message)
}

func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	postId, isPostIdNotOk := getPostId(w, r)
	if isPostIdNotOk {
		return
	}

	ctx, userId, isUserIdNotOk := getUserId(w, r)
	if isUserIdNotOk {
		return
	}

	var post model2.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	if post.NewsletterId.IsEmpty() {
		http.Error(w,
			fmt.Sprintf("Invalid request body, newsletterId can´t be nil! \nNewsletterId: %s",
				post.NewsletterId),
			http.StatusBadRequest)
		return
	}

	var servicePost = model.Post{ID: postId, NewsletterId: post.NewsletterId}

	deleted, err := h.service.DeletePost(ctx, servicePost, userId)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
	}
	util.WriteResponse(w, http.StatusOK, deleted)
}

func getPostId(w http.ResponseWriter, r *http.Request) (string, bool) {
	postId := chi.URLParam(r, "id")
	if postId == "" {
		http.Error(w, "invalid post ID", http.StatusBadRequest)
		return "", true
	}
	return postId, false
}
