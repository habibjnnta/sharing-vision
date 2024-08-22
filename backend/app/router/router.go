package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sharing-vision/app/middleware"
	"sharing-vision/app/tools"
	"sharing-vision/database/connection"
	"sharing-vision/database/ent"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	postsHandler "sharing-vision/app/posts/handler"
	postsRepo "sharing-vision/app/posts/repository"
	postsUC "sharing-vision/app/posts/usecase"
)

type Handlers struct {
	Ctx context.Context
	DB  *ent.Client
	R   *gin.Engine
	Log *logrus.Entry
}

func (h *Handlers) Routes() {
	middleware.Add(h.R, middleware.CORSMiddleware())
	v1 := h.R.Group(os.Getenv("PREFIX_API"))

	h.R.Use(func(c *gin.Context) {
		go routine()
		c.Next()
	})

	v1.GET("/check-connection", h.CheckConnection)

	// Swagger
	v1.GET("/documentation/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Repository
	PostsRepo := postsRepo.NewPostRepository(h.DB)

	// Usecase
	PostsUC := postsUC.NewPostUsecase(PostsRepo, h.Ctx)

	// Handler
	postsHandler.PostRoute(PostsUC, v1, h.Log)
}

func routine() {
	time.Sleep(1 * time.Second)
}

// @BasePath /api/v1
// @Router /check-connection [get]
// @Accept json
// @Produce json
func (h *Handlers) CheckConnection(c *gin.Context) {
	c.JSON(http.StatusOK, tools.Response{
		Status:  "success",
		Message: "Success Check Connect to API",
	})
}

func (h *Handlers) DeleteDatabaseTesting(c *gin.Context) {
	result := connection.DeleteDatabase(os.Getenv("DB_USER_TEST"), os.Getenv("DB_PASS_TEST"), os.Getenv("DB_HOST_TEST"), os.Getenv("DB_PORT_TEST"), os.Getenv("DB_NAME_TEST"))
	if !result {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Errorf("database not deleted"),
		})
		return
	}

	c.JSON(http.StatusOK, tools.Response{
		Status:  "success",
		Message: "Success Delete Database Test",
	})
}
