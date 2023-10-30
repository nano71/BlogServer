package api

import (
	"blogServer/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log/slog"
	"strconv"
	"strings"
	"time"
)

type Article struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Content      string `json:"content"`
	UpdateTime   string `json:"updateTime"`
	CreateTime   string `json:"createTime"`
	Tags         string `json:"tags"`
	CoverImage   string `json:"coverImage"`
	ReadCount    int    `json:"readCount"`
	CommentCount int    `json:"commentCount"`
	Markdown     string `json:"markdown"`
}

func GetArticleList(c *gin.Context) {
	p := &struct {
		Limit int `json:"limit" binding:"required"`
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
		var total int64
		db.Model(Article{}).Count(&total).Limit(p.Limit).Offset(p.Page * p.Limit).Order("create_time desc").Find(articles)
		data := gin.H{
			"total": total,
			"list":  articles,
		}
		response.Success(c, data)
	})

}

func SearchArticles(c *gin.Context) {
	p := &struct {
		Query string `json:"query" binding:"required"`
		Limit int    `json:"limit" binding:"required"`
		Page  int    `json:"page"`
	}{}

	preprocess(c, p, func(db *gorm.DB) {
		articles := &[]Article{}
		search := "%" + p.Query + "%"
		where := db.Model(Article{}).Where("title like ?", search).Or("content like ?", search)
		var total int64
		where.Count(&total).Limit(p.Limit).Offset(p.Page * p.Limit).Order("create_time desc").Find(articles)
		data := gin.H{
			"total": total,
			"list":  articles,
		}
		response.Success(c, data)
	})

}

func SearchArticlesByTag(c *gin.Context) {
	p := &struct {
		Tag   string `json:"tag" binding:"required"`
		Limit int    `json:"limit" binding:"required"`
		Page  int    `json:"page"`
	}{}

	preprocess(c, p, func(db *gorm.DB) {
		articles := &[]Article{}
		where := db.Model(Article{}).Where("tags like ?", "%"+p.Tag+"%")
		var total int64
		where.Count(&total).Limit(p.Limit).Offset(p.Page * p.Limit).Order("create_time desc").Find(articles)
		data := gin.H{
			"total": total,
			"list":  articles,
		}
		response.Success(c, data)
	})
}

func GetArticleContent(c *gin.Context) {
	p := &struct {
		ArticleId int `json:"articleId" binding:"required"`
	}{}
	preprocess(c, p, func(db *gorm.DB) {
		article := &Article{
			Id: p.ArticleId,
		}
		db.First(article)

		cookie, err := c.Cookie("viewed_article_ids")
		slog.Info("cookie", cookie)
		articleIdStringify := strconv.Itoa(p.ArticleId)
		domain := ".nano71.com"
		canAddReadCount := false

		if err != nil {
			canAddReadCount = true
			c.SetCookie("viewed_article_ids", articleIdStringify, 3600, "/", domain, false, true)
		} else {
			viewedArticleIdList := strings.Split(cookie, ",")
			found := false
			for _, item := range viewedArticleIdList {
				if item == articleIdStringify {
					found = true
					break
				}
			}
			if !found {
				canAddReadCount = true
				viewedArticleIdList = append(viewedArticleIdList, articleIdStringify)
				cookie = strings.Join(viewedArticleIdList, ",")
				c.SetCookie("viewed_article_ids", cookie, 3600, "/", domain, false, true)
			}
		}
		if canAddReadCount {
			db.Model(article).Update("read_count", article.ReadCount+1)
		}

		response.Success(c, article)
	})
}

func PublishArticle(c *gin.Context) {
	p := &struct {
		Title       string `json:"title" binding:"required"`
		Content     string `json:"content" binding:"required"`
		Description string `json:"description" binding:"required"`
		Markdown    string `json:"markdown" binding:"required"`
		CreateTime  string `json:"createTime" binding:"required" `
		CoverImage  string `json:"coverImage"`
		Tags        string `json:"tags"`
	}{}
	preprocess(c, p, func(db *gorm.DB) {
		article := &Article{
			Title:       p.Title,
			Content:     p.Content,
			Description: p.Description,
			Markdown:    p.Markdown,
			CreateTime:  p.CreateTime,
			CoverImage:  p.CoverImage,
			Tags:        p.Tags,
		}
		result := db.Omit("UpdateTime").Create(article)
		if result.RowsAffected == 1 {
			response.Success(c, true)
			shouldFetchData = true
		} else {
			response.Fail(c, "文章发布失败")
		}
	})
}
