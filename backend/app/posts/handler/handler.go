package handler

import (
	"sharing-vision/app/posts"
	"sharing-vision/app/tools"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PostHandler struct {
	uc  posts.PostUsecase
	log *logrus.Entry
}

func PostRoute(uc posts.PostUsecase, r *gin.RouterGroup, log *logrus.Entry) {
	h := PostHandler{
		uc:  uc,
		log: log,
	}

	v2 := r.Group("article")

	v2.GET("", h.GetAllPosts)
	v2.GET("/:id", h.GetDetailPosts)
	v2.POST("", h.CreatePosts)
	v2.PUT("/:id", h.UpdatePosts)
	v2.DELETE("/:id", h.DeletePosts)
}

// @Tags Article
// @Summary Get All Article
// @Description Get All Article
// @Router /article [get]
// @Accept json
// @Produce json
// @param status query string false "Filter by Status" enums(publish,draft,trash)
// @param limit query integer false "Limit of pagination"
// @param page query integer false "Page of pagination"
func (h *PostHandler) GetAllPosts(c *gin.Context) {
	result, pagination, code, err := h.uc.GetAllPosts(c)
	if err != nil {
		h.log.Errorf("get all posts handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(code, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get All Articles",
		Meta:    pagination,
	})
}

// @Tags Article
// @Summary Get Detail Article
// @Description Get Detail Article by ID
// @Router /article/{id} [get]
// @Accept json
// @Produce json
// @param id path int true "Article ID"
func (h *PostHandler) GetDetailPosts(c *gin.Context) {
	result, code, err := h.uc.GetDetailPosts(c)
	if err != nil {
		h.log.Errorf("get detail posts handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(code, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get Detail Articles",
	})
}

// @Tags Article
// @Summary Create Article
// @Description Create Article
// @Router /article [post]
// @Accept json
// @Produce json
// @Param request body posts.Form true "Payload Body for Create Article [RAW]"
func (h *PostHandler) CreatePosts(c *gin.Context) {
	code, err := h.uc.CreatePosts(c)
	if err != nil {
		h.log.Errorf("create posts handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Create Articles",
	})
}

// @Tags Article
// @Summary Update Article
// @Description Update Article
// @Router /article/{id} [put]
// @Accept json
// @Produce json
// @param id path int true "Article ID"
// @Param request body posts.Form true "Payload Body for Update Article [RAW]"
func (h *PostHandler) UpdatePosts(c *gin.Context) {
	code, err := h.uc.UpdatePosts(c)
	if err != nil {
		h.log.Errorf("update posts handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Update Articles",
	})
}

// @Tags Article
// @Summary Delete Article
// @Description Delete Article by ID
// @Router /article/{id} [delete]
// @Accept json
// @Produce json
// @param id path int true "Article ID"
func (h *PostHandler) DeletePosts(c *gin.Context) {
	code, err := h.uc.DeletePosts(c)
	if err != nil {
		h.log.Errorf("delete posts handlers: %v", err)
		c.AbortWithStatusJSON(code, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(code, tools.Response{
		Status:  "success",
		Message: "Delete Articles",
	})
}
