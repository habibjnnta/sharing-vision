package usecase

import (
	"context"
	"fmt"
	"net/http"
	"sharing-vision/app/posts"
	"sharing-vision/app/tools"
	"strconv"

	protoPosts "sharing-vision/proto/sharing-vision/app/posts"

	"github.com/gin-gonic/gin"
)

type PostUsecase struct {
	repo posts.PostRepository
	ctx  context.Context
}

func NewPostUsecase(repo posts.PostRepository, ctx context.Context) *PostUsecase {
	return &PostUsecase{
		repo: repo,
		ctx:  ctx,
	}
}

func (uc *PostUsecase) GetAllPosts(c *gin.Context) ([]*protoPosts.Posts, *tools.Pagination, int, error) {
	pagination, err := tools.Paginate(c)
	if err != nil {
		return nil, nil, http.StatusBadRequest, err
	}

	filter := new(posts.FilterPosts)
	if err := c.ShouldBindQuery(filter); err != nil {
		return nil, nil, http.StatusBadRequest, err
	}

	result, pagination, err := uc.repo.GetAllPosts(uc.ctx, pagination, filter)
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	return result, pagination, http.StatusOK, nil
}

func (uc *PostUsecase) GetDetailPosts(c *gin.Context) (*protoPosts.Posts, int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	result, err := uc.repo.GetDetailPosts(uc.ctx, id)
	if result == nil {
		return nil, http.StatusNotFound, fmt.Errorf("articles not found")
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return result, http.StatusOK, nil
}

func (uc *PostUsecase) CreatePosts(c *gin.Context) (int, error) {
	var createArticles *posts.Form
	err := c.ShouldBindJSON(&createArticles)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = uc.repo.CreatePosts(uc.ctx, createArticles)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (uc *PostUsecase) UpdatePosts(c *gin.Context) (int, error) {
	var updateArticles *posts.Form
	err := c.ShouldBindJSON(&updateArticles)
	if err != nil {
		return http.StatusBadRequest, err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return http.StatusBadRequest, err
	}

	updateArticles.ID = id

	err = uc.repo.UpdatePosts(uc.ctx, updateArticles)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusAccepted, nil
}

func (uc *PostUsecase) DeletePosts(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = uc.repo.DeletePosts(uc.ctx, id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusNoContent, nil
}
