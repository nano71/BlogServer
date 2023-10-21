package router

import (
	"blogServer/api"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func Default() gin.HandlerFunc {
	r := &Router{
		routes: make(map[string]gin.HandlerFunc),
	}

	r.POST("/api/addComment", api.AddComment)

	r.POST("/api/validateKey", api.ValidateKey)

	r.POST("/api/getArticleComments", api.GetArticleComments)

	r.POST("/api/getArticleContent", api.GetArticleContent)

	r.POST("/api/getArticleList", api.GetArticleList)

	r.POST("/api/searchArticles", api.SearchArticles)

	r.POST("/api/searchArticlesByTag", api.SearchArticlesByTag)

	r.POST("/api/getTagList", api.GetTagList)

	r.POST("/api/uploadImage", api.UploadImage)

	r.POST("/api/publishArticle", api.PublishArticle)

	r.POST("/api/getPermission", api.GetPermission)
	return r.Preprocessor()
}
