package api

import (
	"blogServer/database"
	"blogServer/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log/slog"
	"time"
)

type Article struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	UpdateTime  time.Time `json:"updateTime"`
	CreateTime  time.Time `json:"createTime"`
	Tags        string    `json:"tags"`
	CoverImage  string    `json:"coverImage"`
	ReadCount   int       `json:"readCount"`
}

func preprocess(c *gin.Context, p interface{}, callback func(*gorm.DB)) {
	c.Header("Content-Type", "application/json")
	if p != nil {
		err := c.BindJSON(&p)
		if err != nil {
			response.ParameterError(c)
			return
		}
	}
	callback(database.GetDB())
}

func GetArticleList(c *gin.Context) {
	p := &struct {
		Limit int `json:"limit"`
		Page  int `json:"page"`
	}{}
	preprocess(c, p, func(db *gorm.DB) {
		articles := &[]struct {
			Id          int       `json:"id"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			CreateTime  time.Time `json:"createTime"`
			Tags        string    `json:"tags"`
			CoverImage  string    `json:"coverImage"`
		}{}
		slog.Info("", p)
		db.Model(Article{}).Limit(p.Limit).Offset(p.Page * p.Limit).Order("create_time desc").Find(articles)
		var count int64
		db.Model(Article{}).Count(&count)
		data := gin.H{
			"count": count,
			"list":  articles,
		}
		response.Success(c, data)
	})

}

func SearchArticles(c *gin.Context) {
	p := &struct {
		Search string `json:"search"`
		Limit  int    `json:"limit"`
		Page   int    `json:"page"`
	}{}

	preprocess(c, p, func(db *gorm.DB) {
		articles := &[]Article{}
		search := "%" + p.Search + "%"
		where := db.Where("title like ?", search).Or("content like ?", search)
		where.Find(articles)
		var count int64
		where.Count(&count)
		data := gin.H{
			"count": count,
			"list":  articles,
		}
		response.Success(c, data)
	})

}

func SearchArticlesByTag(c *gin.Context) {
	p := &struct {
		Label string `json:"label"`
		Limit int    `json:"limit"`
		Page  int    `json:"page"`
	}{}

	preprocess(c, p, func(db *gorm.DB) {
		articles := &[]Article{}
		where := db.Where("labels like ?", "%"+p.Label+"%")
		where.Find(articles)

		var count int64
		where.Count(&count)
		data := gin.H{
			"count": count,
			"list":  articles,
		}
		response.Success(c, data)
	})
}

func GetArticleContent(c *gin.Context) {
	p := &struct {
		ArticleId int `json:"articleId"`
	}{}
	preprocess(c, p, func(db *gorm.DB) {
		article := &Article{
			Id: p.ArticleId,
		}
		db.First(article)
		response.Success(c, article)
	})
}
