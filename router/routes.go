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

	r.POST("/", func(context *gin.Context) {
		slog.Info("api", context.Request.Body)
	})
	r.POST("/api/addComment", api.AddComment)

	r.POST("/api/validateKey", api.ValidateKey)

	r.POST("/api/getArticleComments", api.GetArticleComments)

	r.POST("/api/getArticleContent", api.GetArticleContent)

	r.POST("/api/getArticleList", api.GetArticleList)

	r.POST("/api/searchArticles", api.SearchArticles)

	r.POST("/api/searchArticlesByLabel", api.SearchArticlesByLabel)

	return r.Preprocessor()
}
