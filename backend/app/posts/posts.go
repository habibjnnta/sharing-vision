package posts

import (
	"context"
	"sharing-vision/app/tools"

	"github.com/gin-gonic/gin"

	protoPosts "sharing-vision/proto/sharing-vision/app/posts"
)

type Form struct {
	ID       int    `json:"-"`
	Title    string `json:"title" binding:"required,min=20"`
	Content  string `json:"content" binding:"required,min=200"`
	Category string `json:"category" binding:"required,min=3"`
	Status   string `json:"status" binding:"required,oneof=publish draft trash"`
}

type FilterPosts struct {
	Status string `json:"status" form:"status"`
}

type PostUsecase interface {
	GetAllPosts(c *gin.Context) ([]*protoPosts.Posts, *tools.Pagination, int, error)
	GetDetailPosts(c *gin.Context) (*protoPosts.Posts, int, error)
	CreatePosts(c *gin.Context) (int, error)
	UpdatePosts(c *gin.Context) (int, error)
	DeletePosts(c *gin.Context) (int, error)
}

type PostRepository interface {
	GetAllPosts(ctx context.Context, pagination *tools.Pagination, filter *FilterPosts) ([]*protoPosts.Posts, *tools.Pagination, error)
	GetDetailPosts(ctx context.Context, ID int) (*protoPosts.Posts, error)
	CreatePosts(ctx context.Context, form *Form) error
	UpdatePosts(ctx context.Context, form *Form) error
	DeletePosts(ctx context.Context, ID int) error
}
