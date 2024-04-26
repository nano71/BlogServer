package api

import (
	"blogServer/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sort"
)

type Tag struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Count   int64  `json:"count" gorm:"-"`
}

var tags []Tag
var shouldFetchData = true

func GetTagList(c *gin.Context) {

	preprocess(c, nil, func(db *gorm.DB) {
		if !shouldFetchData {
			response.Success(c, tags)
			return
		}
		db.Find(&tags)

		for i, tag := range tags {
			var count int64
			db.Model(&Article{}).Where("tags LIKE ?", fmt.Sprintf("%%%s%%", tag.Name)).Count(&count)
			tags[i].Count = count
		}

		sort.Slice(tags, func(i, j int) bool {
			return tags[i].Count > tags[j].Count
		})
		shouldFetchData = false
		response.Success(c, tags)
	})
}

func AddCategory(c *gin.Context) {
	p := &struct {
		Name    string `json:"name" binding:"required"`
		Content string `json:"content" binding:"required"`
	}{}
	preprocess(c, p, func(db *gorm.DB) {
		tag := &Tag{
			Name:    p.Name,
			Content: p.Content,
		}
		result := db.Create(tag)
		if result.RowsAffected == 1 {
			response.Success(c, true)
			shouldFetchData = true
		} else {
			response.Fail(c, "标签添加失败")
		}
	})
}

func DeleteCategory(c *gin.Context) {
	p := &struct {
		TagId int `json:"tagId" binding:"required"`
	}{}
	preprocess(c, p, func(db *gorm.DB) {
		tag := &Tag{
			Id: p.TagId,
		}
		result := db.Delete(tag)
		if result.RowsAffected == 1 {
			response.Success(c, true)
			shouldFetchData = true
		} else {
			response.Fail(c, "标签删除失败")
		}
	})
}

func UpdateCategory(c *gin.Context) {
	p := &struct {
		TagId   int    `json:"tagId" binding:"required"`
		Name    string `json:"name" binding:"required"`
		Content string `json:"content" binding:"required"`
	}{}
	preprocess(c, p, func(db *gorm.DB) {
		tag := &Tag{
			Id:      p.TagId,
			Name:    p.Name,
			Content: p.Content,
		}
		result := db.Updates(tag)
		if result.RowsAffected == 1 {
			response.Success(c, true)
			shouldFetchData = true
		} else {
			response.Fail(c, "标签更新失败")
		}
	})
}
