package router

import (
	"blogServer/api"
	"github.com/gin-gonic/gin"
)

func Default() gin.HandlerFunc {
	r := &Router{
		routes: make(map[string]gin.HandlerFunc),
	}

	r.POST("/api/validateKey", api.ValidateKey)

	r.POST("/api/getArticleContent", api.GetArticleContent)

	r.POST("/api/getArticleList", api.GetArticleList)

	r.POST("/api/manage/getArticleList", api.ManagerGetArticleList)

	r.POST("/api/getMessageList", api.GetMessageList)

	r.POST("/api/searchArticles", api.SearchArticles)

	r.POST("/api/searchArticlesByTag", api.SearchArticlesByTag)

	r.POST("/api/getTagList", api.GetTagList)

	r.POST("/api/uploadImage", api.UploadImage)

	r.POST("/api/publishArticle", api.PublishArticle)

	r.POST("/api/updateArticle", api.UpdateArticle)

	r.POST("/api/leaveMessage", api.LeaveMessage)

	r.POST("/api/getPermission", api.GetPermission)

	r.POST("/api/updateArticleCommentCount", api.UpdateArticleCommentCount)

	return r.Preprocessor()
}
