package api

import (
	"blogServer/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateArticleCommentCount(c *gin.Context) {
	p := &struct {
		ArticleId int `json:"articleId" binding:"required"`
		Count     int `json:"count" binding:"required"`
	}{}
	preprocess(c, p, func(db *gorm.DB) {
		db.Model(&Article{
			Id: p.ArticleId,
		}).Update("comment_count", p.Count)
		response.Success(c, nil)
	})
}
