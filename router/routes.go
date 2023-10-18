package router

import (
	"blogServer/api"
	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	r := &Router{
		routes: make(map[string]gin.HandlerFunc),
	}

	//r.POST("/", func(context *gin.Context) {
	//	slog.Info("api", context.Request.Body)
	//})
	//r.POST("/api/addComment", api.AddComment)
	//
	//r.POST("/api/validateKey", api.ValidateKey)
	//
	//r.POST("/api/getComments", api.GetComments)

	//r.POST("/api/getArticle",)
	r.POST("/api/getArticleList", api.GetArticleList)

	return r.Preprocessor()
}
