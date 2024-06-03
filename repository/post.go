package repository

import (
	"context"
	"fmt"

	dbmodel "vse-go-newsletter-api/repository/sql/model"
	"vse-go-newsletter-api/repository/sql/query"
	"vse-go-newsletter-api/service/model"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostRepository struct {
	pool *pgxpool.Pool
}

func NewPostRepository(pool *pgxpool.Pool) *PostRepository {
	return &PostRepository{
		pool: pool,
	}
}

func (r *PostRepository) CreatePost(ctx context.Context, post model.Post, userId string) (*model.Post, error) {
	var createdPost dbmodel.Post

	err := pgxscan.Get(
		ctx,
		r.pool,
		&createdPost,
		query.CreatePost,
		pgx.NamedArgs{
			"title":         post.Title,
			"content":       post.Content,
			"newsletter_id": post.NewsletterId,
			"user_id":       userId,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}
	newPost := &model.Post{
		ID:           createdPost.ID,
		CreatedAt:    createdPost.CreatedAt,
		UpdatedAt:    createdPost.UpdatedAt,
		Title:        createdPost.Title,
		Content:      createdPost.Content,
		NewsletterId: createdPost.NewsletterId,
		IsPublished:  createdPost.IsPublished,
	}
	return newPost, nil
}

func (r *PostRepository) ListPosts(ctx context.Context) ([]model.Post, error) {
	var posts []dbmodel.Post
	if err := pgxscan.Select(
		ctx,
		r.pool,
		&posts,
		query.ListPost,
	); err != nil {
		return nil, err
	}
	response := make([]model.Post, len(posts))
	for i, post := range posts {
		response[i] = model.Post{
			ID:           post.ID,
			Title:        post.Title,
			Content:      post.Content,
			NewsletterId: post.NewsletterId,
			CreatedAt:    post.CreatedAt,
			UpdatedAt:    post.UpdatedAt,
			IsPublished:  post.IsPublished,
		}
	}
	return response, nil
}

func (r *PostRepository) ReadPost(ctx context.Context, postId string) (*model.Post, error) {
	var post dbmodel.Post
	if err := pgxscan.Get(
		ctx,
		r.pool,
		&post,
		query.ReadPost,
		pgx.NamedArgs{
			"id": postId,
		},
	); err != nil {
		return nil, err
	}
	return &model.Post{
		ID:           post.ID,
		Title:        post.Title,
		Content:      post.Content,
		NewsletterId: post.NewsletterId,
		CreatedAt:    post.CreatedAt,
		UpdatedAt:    post.UpdatedAt,
		IsPublished:  post.IsPublished,
	}, nil

}

func (r *PostRepository) UpdatePost(ctx context.Context, post model.Post, userId string) (*model.Post, error) {
	var dbPost dbmodel.Post

	if err := pgxscan.Get(
		ctx,
		r.pool,
		&dbPost,
		query.UpdatePost,
		pgx.NamedArgs{
			"id":				post.ID,
			"title":		post.Title,
			"content":	post.Content,
			"user_id":	userId,
		},
	); err != nil {
		return nil, err
	}

	updatedPost := &model.Post{
		ID:           dbPost.ID,
		Title:        dbPost.Title,
		Content:      dbPost.Content,
		NewsletterId: dbPost.NewsletterId,
		CreatedAt:    dbPost.CreatedAt,
		UpdatedAt:    dbPost.UpdatedAt,
		IsPublished:  dbPost.IsPublished,
	}

	return updatedPost, nil
}

func (r *PostRepository) DeletePost(ctx context.Context, postId string, userId string) (string, error) {
	result, err := r.pool.Exec(ctx, query.DeletePost, pgx.NamedArgs{
		"id":      postId,
		"user_id": userId,
	})

	if err != nil {
		message := fmt.Sprintf("post not deleted! ID: %s", postId)
		return message, err
	}

	if result.RowsAffected() == 0 {
		message := fmt.Sprintf("no post found to delete or you are not allowed to delete it. ID: %s", postId)
		return message, fmt.Errorf("no rows affected")
	}
	message := fmt.Sprintf("post deleted successfully! ID: %s", postId)
	return message, nil
}

func (r *PostRepository) PublishPost(ctx context.Context, postId string, userId string) (*model.Post, error) {
	var dbPost dbmodel.Post

	if err := pgxscan.Get(
		ctx,
		r.pool,
		&dbPost,
		query.PublishPost,
		pgx.NamedArgs{"id": postId,
			"user_id": userId,
		},
	); err != nil {
		return nil, err
	}

	updatedPost := &model.Post{
		ID:           dbPost.ID,
		Title:        dbPost.Title,
		Content:      dbPost.Content,
		NewsletterId: dbPost.NewsletterId,
		CreatedAt:    dbPost.CreatedAt,
		UpdatedAt:    dbPost.UpdatedAt,
		IsPublished:  dbPost.IsPublished,
	}

	return updatedPost, nil
}
