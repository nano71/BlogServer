package api

import (
	"blogServer/response"
	"crypto/md5"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.MissingParameters(c)
		return
	}
	filename := []byte(file.Filename)
	hash := md5.Sum(filename)
	imagePath := "uploads/" + strconv.Itoa(int(time.Now().UnixNano()/1e6)) + "-" + fmt.Sprintf("%x", hash)
	err = c.SaveUploadedFile(file, imagePath)
	if err != nil {
		response.Fail(c, "文件保存失败")
		return
	}
	response.Success(c, "/"+imagePath)
}
