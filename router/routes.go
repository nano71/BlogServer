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

	r.POST("/api/manage/deleteMessage", api.DeleteMessage)

	r.POST("/api/manage/addCategory", api.AddCategory)

	r.POST("/api/manage/deleteCategory", api.DeleteCategory)

	r.POST("/api/manage/updateCategory", api.UpdateCategory)

	r.POST("/api/manage/publishArticle", api.PublishArticle)

	r.POST("/api/manage/deleteArticle", api.DeleteArticle)

	r.POST("/api/manage/updateArticle", api.UpdateArticle)

	r.POST("/api/manage/getArticleList", api.ManagerGetArticleList)

	r.POST("/api/manage/getDailyVisitorVolume", api.GetDailyVisitorVolume)
	r.POST("/api/manage/getDailyBannedCount", api.GetDailyBannedCount)

	r.POST("/api/getMessageList", api.GetMessageList)

	r.POST("/api/searchArticles", api.SearchArticles)

	r.POST("/api/searchArticlesByTag", api.SearchArticlesByTag)

	r.POST("/api/getTagList", api.GetTagList)

	r.POST("/api/uploadImage", api.UploadImage)

	r.POST("/api/leaveMessage", api.LeaveMessage)

	r.POST("/api/getPermission", api.GetPermission)

	r.POST("/api/updateArticleCommentCount", api.UpdateArticleCommentCount)

	return r.Preprocessor()
}
