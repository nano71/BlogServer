package api

import (
	"blogServer/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Comment struct {
	ArticleId int    `json:"articleId"`
	Email     string `json:"email"`
	Content   string `json:"content"`
	Site      string `json:"site"`
}

func GetArticleComments(c *gin.Context) {
	p := &struct {
		ArticleId int `json:"articleId"`
		Limit     int `json:"limit"`
		Page      int `json:"page"`
	}{}

	preprocess(c, p, func(db *gorm.DB) {
		comments := &[]Comment{}
		db.Where("article_id = ?", p.ArticleId).Limit(p.Limit).Offset(p.Page * p.Limit).Find(comments)
		response.Success(c, comments)
	})

}

func AddComment(c *gin.Context) {
	p := &Comment{}

	preprocess(c, p, func(db *gorm.DB) {
		result := db.Create(p)
		if result.Error != nil {
			response.Fail(c, "数据插入失败!")
			return
		}
		response.Success(c, nil)
	})

}
