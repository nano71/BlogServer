package api

import (
	"blogServer/response"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.MissingParameters(c)
		return
	}
	imagePath := "uploads/" + strconv.Itoa(int(time.Now().UnixNano()/1e6)) + "-" + file.Filename
	err = c.SaveUploadedFile(file, imagePath)
	if err != nil {
		response.Fail(c, "文件保存失败")
		return
	}
	response.Success(c, "/"+imagePath)
}
