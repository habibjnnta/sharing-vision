package test

import (
	"context"
	"os"
	"sharing-vision/database/connection"
	"sharing-vision/database/ent"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	godotenv.Load()
	os.Setenv("DB_USER_TEST", "root")
	os.Setenv("DB_PASS_TEST", "")
	os.Setenv("DB_HOST_TEST", "localhost")
	os.Setenv("DB_PORT_TEST", "3306")
	os.Setenv("DB_NAME_TEST", "article_test")
}

type HandlerTesting struct {
	Ctx   context.Context
	DB    *ent.Client
	Route *gin.Engine
	Log   *logrus.Entry
}

func SetUpRouter() HandlerTesting {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	ctx := context.Background()
	log := logrus.NewEntry(logrus.StandardLogger())
	db := connection.ConnectionDB(ctx, log, os.Getenv("DB_USER_TEST"), os.Getenv("DB_PASS_TEST"), os.Getenv("DB_HOST_TEST"), os.Getenv("DB_PORT_TEST"), os.Getenv("DB_NAME_TEST"))

	return HandlerTesting{
		Ctx:   ctx,
		DB:    db,
		Route: r,
		Log:   log,
	}
}
