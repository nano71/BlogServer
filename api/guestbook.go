package api

import (
	"blogServer/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	Id         int    `json:"id"`
	Nickname   string `json:"title"`
	Url        string `json:"url" `
	Face       string `json:"face"`
	Content    string `json:"content"`
	CreateTime string `json:"createTime"`
	Ip         string `json:"ip"`
	IsVisible  int    `json:"isVisible"`
}

func (Message) TableName() string {
	return "guestbook"
}

func LeaveMessage(c *gin.Context) {
	p := &struct {
		Nickname   string `json:"nickname"`
		Url        string `json:"url"`
		Face       string `json:"face" binding:"required"`
		Content    string `json:"content" binding:"required"`
		CreateTime string `json:"createTime" binding:"required" `
	}{}
	preprocess(c, p, func(db *gorm.DB) {
		message := &Message{
			Nickname:   p.Nickname,
			Url:        p.Url,
			Face:       p.Face,
			Content:    p.Content,
			CreateTime: p.CreateTime,
			Ip:         c.ClientIP(),
		}
		result := db.Model(Message{}).Create(message)
		if result.RowsAffected == 1 {
			response.Success(c, true)
		} else {
			response.Fail(c, "留言发布失败")
		}
	})
}

func GetMessageList(c *gin.Context) {
	p := &struct {
		Limit int `json:"limit" binding:"required"`
		Page  int `json:"page"`
	}{}
	preprocess(c, p, func(db *gorm.DB) {
		messageList := &[]struct {
			Id         int       `json:"id"`
			Nickname   string    `json:"nickname"`
			Url        string    `json:"url"`
			Face       string    `json:"face"`
			Content    string    `json:"content"`
			CreateTime time.Time `json:"createTime"`
		}{}
		var total int64
		db.Model(Message{}).Count(&total).Limit(p.Limit).Offset(p.Page * p.Limit).Order("create_time desc").Find(messageList)
		data := gin.H{
			"total": total,
			"list":  messageList,
		}
		response.Success(c, data)
	})
}

func DeleteMessage(c *gin.Context) {
	p := &struct {
		MessageId int `json:"messageId" binding:"required"`
	}{}
	preprocess(c, p, func(db *gorm.DB) {
		message := &Message{
			Id: p.MessageId,
		}
		result := db.Delete(&message)
		if result.RowsAffected == 1 {
			response.Success(c, true)
		} else {
			response.Fail(c, "留言删除失败")
		}
	})
}
