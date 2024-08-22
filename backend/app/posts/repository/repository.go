package repository

import (
	"context"
	"sharing-vision/app/posts"
	"sharing-vision/app/tools"
	"sharing-vision/database/ent"

	entPosts "sharing-vision/database/ent/posts"
	protoPosts "sharing-vision/proto/sharing-vision/app/posts"
)

type PostRepository struct {
	db *ent.Client
}

func NewPostRepository(db *ent.Client) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) GetAllPosts(ctx context.Context, pagination *tools.Pagination, filter *posts.FilterPosts) ([]*protoPosts.Posts, *tools.Pagination, error) {
	postsQuery := r.db.Posts.Query()

	if filter.Status != "" {
		postsQuery.Where(
			entPosts.StatusEQ(entPosts.Status(filter.Status)),
		)
	}

	dbPosts, err := postsQuery.Offset(pagination.Offset).Limit(pagination.Limit).All(ctx)
	if err != nil {
		return nil, nil, err
	}

	// Pagination
	count, _ := postsQuery.Count(ctx)
	pagination.Count = int(count)
	pagination = tools.Paging(pagination)

	var articles []*protoPosts.Posts

	for _, v := range dbPosts {
		articles = append(articles, &protoPosts.Posts{
			Id:       int32(v.ID),
			Title:    v.Title,
			Content:  v.Content,
			Category: v.Category,
			Status:   string(v.Status),
		})
	}

	return articles, pagination, nil
}

func (r *PostRepository) GetDetailPosts(ctx context.Context, ID int) (*protoPosts.Posts, error) {
	postsQuery := r.db.Posts.Query().
		Where(
			entPosts.And(
				entPosts.StatusNotIn(
					entPosts.Status(entPosts.StatusThrash),
				),
				entPosts.IDEQ(ID),
			),
		)

	if count, _ := postsQuery.Count(ctx); count == 0 {
		return nil, nil
	}

	exec, _ := postsQuery.First(ctx)

	return &protoPosts.Posts{
		Id:       int32(exec.ID),
		Title:    exec.Title,
		Content:  exec.Content,
		Category: exec.Category,
		Status:   string(exec.Status),
	}, nil
}

func (r *PostRepository) CreatePosts(ctx context.Context, form *posts.Form) error {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Posts.Create().
		SetTitle(form.Title).
		SetContent(form.Content).
		SetCategory(form.Category).
		SetStatus(entPosts.Status(form.Status)).
		Save(ctx)

	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) UpdatePosts(ctx context.Context, form *posts.Form) error {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return err
	}

	err = tx.Posts.Update().
		Where(entPosts.IDEQ(form.ID)).
		SetTitle(form.Title).
		SetContent(form.Content).
		SetCategory(form.Category).
		SetStatus(entPosts.Status(form.Status)).
		Exec(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) DeletePosts(ctx context.Context, ID int) error {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return err
	}

	queryPosts, err := tx.Posts.Query().
		Where(
			entPosts.IDEQ(ID),
		).
		First(ctx)

	if err != nil {
		return err
	}

	err = tx.Posts.Update().
		Where(
			entPosts.IDEQ(queryPosts.ID),
		).
		SetStatus(entPosts.StatusThrash).
		Exec(ctx)

	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
