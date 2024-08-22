package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"sharing-vision/app/posts"
	"sharing-vision/app/test"
	"sharing-vision/database/ent"
	"strconv"
	"testing"

	"github.com/sirupsen/logrus"
	"gotest.tools/assert"

	postsRepo "sharing-vision/app/posts/repository"
	postsUC "sharing-vision/app/posts/usecase"
	protoPosts "sharing-vision/proto/sharing-vision/app/posts"
)

var ID int

type Result struct {
	Data []protoPosts.Posts `json:"data"`
}

func ReturnPostHandler(ctx context.Context, db *ent.Client, log *logrus.Entry) PostHandler {
	PostsRepo := postsRepo.NewPostRepository(db)
	PostsUC := postsUC.NewPostUsecase(PostsRepo, ctx)

	return PostHandler{
		uc:  PostsUC,
		log: log,
	}
}

func TestCreatePosts(t *testing.T) {
	ht := test.SetUpRouter()

	h := ReturnPostHandler(ht.Ctx, ht.DB, ht.Log)
	v1 := ht.Route.Group(os.Getenv("PREFIX_API"))
	v1.POST("article", h.CreatePosts)

	body := posts.Form{
		Title:    "testing sdsadsad article",
		Content:  "testing article content sdsadasnlkdlsandn kaslnrkaenkfnkasdksnakldnk lknksn dnasklnd ksankdnkalskld naslkndk askdnksankln askdnk nsakdk lsandknsakn klsankldnk ksandk lasnkl sdsadsadsb djsbdj sjabdjsbajs djbsja dbsajkdj sbjdb asjkjkd bsad",
		Category: "testing article category",
		Status:   "publish",
	}

	url := os.Getenv("PREFIX_API") + "/" + "article"

	reqBody, _ := json.Marshal(body)

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(reqBody))

	w := httptest.NewRecorder()
	ht.Route.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetAllPosts(t *testing.T) {
	ht := test.SetUpRouter()

	h := ReturnPostHandler(ht.Ctx, ht.DB, ht.Log)
	v1 := ht.Route.Group(os.Getenv("PREFIX_API"))
	v1.GET("article", h.GetAllPosts)

	url := os.Getenv("PREFIX_API") + "/" + "article" + "?" + "limit=1"

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	w := httptest.NewRecorder()
	ht.Route.ServeHTTP(w, req)

	body, _ := ioutil.ReadAll(w.Body)

	var result Result
	_ = json.Unmarshal(body, &result)

	ID = int(result.Data[0].Id)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetDetailPosts(t *testing.T) {
	ht := test.SetUpRouter()

	h := ReturnPostHandler(ht.Ctx, ht.DB, ht.Log)
	v1 := ht.Route.Group(os.Getenv("PREFIX_API"))
	v1.GET("article/:id", h.GetDetailPosts)

	url := os.Getenv("PREFIX_API") + "/" + "article" + "/" + strconv.Itoa(ID)

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	w := httptest.NewRecorder()
	ht.Route.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdatePosts(t *testing.T) {
	ht := test.SetUpRouter()

	h := ReturnPostHandler(ht.Ctx, ht.DB, ht.Log)
	v1 := ht.Route.Group(os.Getenv("PREFIX_API"))
	v1.PUT("article/:id", h.UpdatePosts)

	body := posts.Form{
		Title:    "testing article sdsadsdsa update",
		Content:  "testing article content sdsadasnlkdlsandn kaslnrkaenkfnkasdksnakldnk lknksn dnasklnd ksankdnkalskld naslkndk askdnksankln askdnk nsakdk lsandknsakn klsankldnk ksandk lasnkl sadkbsabdj busbd jks dbjsbjd bsjdabsj jsdb update",
		Category: "testing article category update",
		Status:   "draft",
	}

	reqBody, _ := json.Marshal(body)

	url := os.Getenv("PREFIX_API") + "/" + "article" + "/" + strconv.Itoa(ID)

	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewReader(reqBody))

	w := httptest.NewRecorder()
	ht.Route.ServeHTTP(w, req)

	assert.Equal(t, http.StatusAccepted, w.Code)
}

func TestDeletePosts(t *testing.T) {
	ht := test.SetUpRouter()

	h := ReturnPostHandler(ht.Ctx, ht.DB, ht.Log)
	v1 := ht.Route.Group(os.Getenv("PREFIX_API"))
	v1.DELETE("article/:id", h.DeletePosts)

	url := os.Getenv("PREFIX_API") + "/" + "article" + "/" + strconv.Itoa(ID)

	req, _ := http.NewRequest(http.MethodDelete, url, nil)

	w := httptest.NewRecorder()
	ht.Route.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
