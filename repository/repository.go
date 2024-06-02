package repository

import (
	"context"
	"fmt"

	"vse-go-newsletter-api/pkg/id"
	dbmodel "vse-go-newsletter-api/repository/sql/model"
	"vse-go-newsletter-api/repository/sql/query"
	"vse-go-newsletter-api/service/model"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	*NewsletterRepository
	*PostRepository
}

func New(pool *pgxpool.Pool) (*Repository, error) {
	return &Repository{
		NewsletterRepository: NewNewsletterRepository(pool),
		PostRepository:       NewPostRepository(pool),
	}, nil
}

type NewsletterRepository struct {
	pool *pgxpool.Pool
}

func NewNewsletterRepository(pool *pgxpool.Pool) *NewsletterRepository {
	return &NewsletterRepository{
		pool: pool,
	}
}

type PostRepository struct {
	pool *pgxpool.Pool
}

func NewPostRepository(pool *pgxpool.Pool) *PostRepository {
	return &PostRepository{
		pool: pool,
	}
}

func (r *NewsletterRepository) ReadNewsletter(ctx context.Context, newsletterID id.Newsletter) (*model.Newsletter, error) {
	var newsletter dbmodel.Newsletter
	if err := pgxscan.Get(
		ctx,
		r.pool,
		&newsletter,
		query.ReadNewsletter,
		pgx.NamedArgs{
			"id": newsletterID,
		},
	); err != nil {
		return nil, err
	}
	return &model.Newsletter{
		ID:          newsletter.ID,
		Name:        newsletter.Name,
		Description: newsletter.Description,
		OwnerId:     newsletter.OwnerId,
		CreateAt:    newsletter.CreatedAt,
		UpdatedAt:   newsletter.UpdatedAt,
	}, nil
}

func (r *NewsletterRepository) ListNewsletter(ctx context.Context) ([]model.Newsletter, error) {
	var newsletters []dbmodel.Newsletter
	if err := pgxscan.Select(
		ctx,
		r.pool,
		&newsletters,
		query.ListNewsletter,
	); err != nil {
		return nil, err
	}
	response := make([]model.Newsletter, len(newsletters))
	for i, newsletter := range newsletters {
		response[i] = model.Newsletter{
			ID:          newsletter.ID,
			Name:        newsletter.Name,
			Description: newsletter.Description,
			OwnerId:     newsletter.OwnerId,
			CreateAt:    newsletter.CreatedAt,
			UpdatedAt:   newsletter.UpdatedAt,
		}
	}
	return response, nil
}

func (r *NewsletterRepository) UpdateNewsletter(ctx context.Context, newsletter model.Newsletter) (*model.Newsletter, error) {
	var dbNewsletter dbmodel.Newsletter

	if err := pgxscan.Get(
		ctx,
		r.pool,
		&dbNewsletter,
		query.UpdateNewsletter,
		pgx.NamedArgs{"id": newsletter.ID,
			"name":        newsletter.Name,
			"description": newsletter.Description,
			"owner_id":    newsletter.OwnerId,
		},
	); err != nil {
		return nil, err
	}

	updatedNewsletter := &model.Newsletter{
		ID:          dbNewsletter.ID,
		Name:        dbNewsletter.Name,
		Description: dbNewsletter.Description,
		OwnerId:     dbNewsletter.OwnerId,
		UpdatedAt:   dbNewsletter.UpdatedAt,
	}

	return updatedNewsletter, nil
}

func (r *NewsletterRepository) DeleteNewsletter(ctx context.Context, newsletter model.Newsletter) (string, error) {
	result, err := r.pool.Exec(
		ctx,
		query.DeleteNewsletter,
		pgx.NamedArgs{"id": newsletter.ID,
			"owner_id": newsletter.OwnerId,
		},
	)

	if err != nil {
		message := fmt.Sprintf("newsletter not deleted! ID: %s", newsletter.ID)
		return message, err
	}

	if result.RowsAffected() == 0 {
		message := fmt.Sprintf("no newsletter found to delete or you are not allowed to delete it. ID: %s", newsletter.ID)
		return message, fmt.Errorf("no rows affected\n")
	}
	message := fmt.Sprintf("newsletter deleted successfully! ID: %s", newsletter.ID)
	return message, nil
}

func (r *NewsletterRepository) CreateNewsletter(ctx context.Context, newsletter model.Newsletter) (*model.Newsletter, error) {
	var createdNewsletter dbmodel.Newsletter

	// Execute the SQL insert query with RETURNING clause
	err := pgxscan.Get(
		ctx,
		r.pool,
		&createdNewsletter,
		query.CreateNewsletter,
		pgx.NamedArgs{"name": newsletter.Name,
			"description": newsletter.Description,
			"owner_id":    newsletter.OwnerId,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create newsletter: %w", err)
	}

	// Convert the db model to the application model
	newNewsletter := &model.Newsletter{
		ID:          createdNewsletter.ID,
		Name:        createdNewsletter.Name,
		Description: createdNewsletter.Description,
		OwnerId:     createdNewsletter.OwnerId,
		CreateAt:    createdNewsletter.CreatedAt,
		UpdatedAt:   createdNewsletter.UpdatedAt,
	}

	// Return the newly created newsletter object
	return newNewsletter, nil
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
		pgx.NamedArgs{"id": post.ID,
			"title":   post.Title,
			"content": post.Content,
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
		return message, fmt.Errorf("no rows affected\n")
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

func (r *NewsletterRepository) SubscribeNewsletter(ctx context.Context, newsletterId id.Newsletter, userId string) (*model.Subscription, error) {
	var subscription dbmodel.Subscription

	err := pgxscan.Get(
		ctx,
		r.pool,
		&subscription,
		query.SubscribeNewsletter,
		pgx.NamedArgs{
			"newsletter_id": newsletterId,
			"user_id":       userId,
		},
	)
	if err != nil {
		return nil, err
	}

	serviceSubscription := &model.Subscription{
		ID:           subscription.ID,
		UserId:       subscription.UserId,
		NewsletterId: subscription.NewsletterId,
		CreatedAt:    subscription.CreatedAt,
	}

	return serviceSubscription, nil
}

func (r *NewsletterRepository) UnsubscribeNewsletter(ctx context.Context, newsletterId id.Newsletter, userId string) (string, error) {
	result, err := r.pool.Exec(ctx, query.UnsubscribeNewsletter, pgx.NamedArgs{
		"newsletter_id": newsletterId,
		"user_id":       userId,
	})

	if err != nil {
		message := fmt.Sprintf("failed to cancel subscription to newsletter ID: %s", newsletterId)
		return message, err
	}

	if result.RowsAffected() == 0 {
		message := fmt.Sprintf("You are not subscribed to the newsletter ID: %s. There was no need to unsubscribe", newsletterId)
		return message, fmt.Errorf("no rows affected\n")
	}
	message := fmt.Sprintf("Successful newsletter unsubscription (ID: %s)", newsletterId)
	return message, nil
}
