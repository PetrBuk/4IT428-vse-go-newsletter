package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"vse-go-newsletter-api/service/model"
	model2 "vse-go-newsletter-api/transport/api/v1/model"
	"vse-go-newsletter-api/transport/util"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5"
)

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post model2.Post

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if post.Title == "" || post.Content == "" || post.NewsletterId.IsEmpty() {
		http.Error(w,
			fmt.Sprintf("Invalid request body, title or content can´t be nil! \nTitle: %s \nContent: %s \nNewsLetterId: %s",
				post.Title,
				post.Content,
				post.NewsletterId),
			http.StatusBadRequest)
		return
	}

	ctx, userId, isUnauthenticated := getUserId(w, r)
	if isUnauthenticated {
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
	ctx, userId, isUnauthenticated := getUserId(w, r)
	if isUnauthenticated {
		return
	}

	if post.Title == "" || post.Content == "" {
		http.Error(w,
			fmt.Sprintf("Invalid request body, title, content or newsletterId can´t be nil! \nTitle: %s \nContent: %s",
				post.Title,
				post.Content),
			http.StatusBadRequest)
		return
	}

	var servicePost = model.Post{ID: postId, Title: post.Title, Content: post.Content}

	updated, err := h.service.UpdatePost(ctx, servicePost, userId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.WriteResponse(w, http.StatusOK, "Post couldn't be updated. Either you are not the owner of the newsletter or the post was already published.")
		} else {
			util.WriteErrResponse(w, http.StatusInternalServerError, err)
		}
		return
	}

	util.WriteResponse(w, http.StatusOK, updated)
}

func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	postId, isPostIdNotOk := getPostId(w, r)
	if isPostIdNotOk {
		return
	}

	ctx, userId, isUnauthenticated := getUserId(w, r)
	if isUnauthenticated {
		return
	}

	var post model2.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	deleted, err := h.service.DeletePost(ctx, postId, userId)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
	}
	util.WriteResponse(w, http.StatusOK, deleted)
}

func (h *Handler) PublishPost(w http.ResponseWriter, r *http.Request) {
	postId, isPostIdNotOk := getPostId(w, r)
	if isPostIdNotOk {
		return
	}

	ctx, userId, isUnauthenticated := getUserId(w, r)
	if isUnauthenticated {
		return
	}

	post, err := h.service.PublishPost(ctx, postId, userId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			util.WriteResponse(w, http.StatusOK, "Post couldn't be published. Either you are not the owner of the newsletter or the post was already published.")
		} else {
			util.WriteErrResponse(w, http.StatusInternalServerError, err)
		}
		return
	}
	util.WriteResponse(w, http.StatusOK, post)
}

func getPostId(w http.ResponseWriter, r *http.Request) (string, bool) {
	postId := chi.URLParam(r, "id")
	if postId == "" {
		http.Error(w, "invalid post ID", http.StatusBadRequest)
		return "", true
	}
	return postId, false
}
