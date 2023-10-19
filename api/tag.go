package api

import (
	"blogServer/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sort"
)

type Tag struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Count   int64  `json:"count"`
}

func GetTagList(c *gin.Context) {

	preprocess(c, nil, func(db *gorm.DB) {
		var tags []Tag
		db.Find(&tags)

		for i, tag := range tags {
			var count int64
			db.Debug().Model(&Article{}).Where("tags LIKE ?", fmt.Sprintf("%%%s%%", tag.Name)).Count(&count)
			tags[i].Count = count
		}

		sort.Slice(tags, func(i, j int) bool {
			return tags[i].Count > tags[j].Count
		})

		response.Success(c, tags)
	})
}
