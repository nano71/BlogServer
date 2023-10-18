package api

import (
	"blogServer/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Article struct {
	Id         int       `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	UpdateTime time.Time `json:"updateTime"`
	CreateTime time.Time `json:"createTime"`
	Labels     string    `json:"labels,omitempty"`
}
type Comment struct {
	ArticleId int    `json:"article_id,omitempty"`
	Email     string `json:"email,omitempty"`
	Content   string `json:"content,omitempty"`
	Site      string `json:"site,omitempty"`
}

func GetComments(c *gin.Context) {
	type Params struct {
		ArticleId int `json:"articleId"`
		Limit     int `json:"limit"`
		Page      int `json:"page"`
	}
	params := Params{}
	db := preprocess(c, params)
	var comments []Comment
	db.Where("article_id = ?", params.ArticleId).Limit(params.Limit).Offset(params.Page * params.Limit).Find(&comments)

	c.JSON(200, comments)
}

func preprocess(c *gin.Context, model interface{}) *gorm.DB {
	err := c.BindJSON(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return nil
	}
	return database.GetDB()
}

func GetArticleList(c *gin.Context) {
	type Params struct {
		Limit int `json:"limit"`
		Page  int `json:"page"`
	}
	var (
		params   = &Params{}
		articles []Article
	)

	db := preprocess(c, params)
	if db == nil {
		return
	}
	db.Limit(params.Limit).Offset(params.Page * params.Limit).Find(&articles)

	c.JSON(200, articles)
}
