package api

import (
	"blogServer/response"
	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.MissingParameters(c)
		return
	}

	err = c.SaveUploadedFile(file, "uploads/"+file.Filename)
	if err != nil {
		response.Fail(c, "文件保存失败")
		return
	}
	response.Success(c, "/uploads/"+file.Filename)
}
