package api

import (
	"blogServer/middleware"
	"blogServer/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sort"
	"time"
)

func GetDailyVisitorVolume(c *gin.Context) {
	sql := `SELECT * FROM (
                  SELECT *,
                         (@row_number := IF(@prev_date = DATE(create_time) AND @prev_ip = ip, @row_number + 1, 1)) AS row_num,
                         (@prev_date := DATE(create_time)) AS dummy_date,
                         (@prev_ip := ip) AS dummy_ip
                  FROM log
                           CROSS JOIN (SELECT @row_number := 0, @prev_date := NULL, @prev_ip := NULL) AS vars
                  WHERE status = 200
                    AND create_time >= DATE_SUB(NOW(), INTERVAL 1 YEAR)
                  ORDER BY ip, create_time
              			) AS subquery WHERE row_num = 1 order by create_time;`

	preprocess(c, nil, func(db *gorm.DB) {
		var logs []middleware.Log
		db.Raw(sql).Scan(&logs)

		type Data struct {
			Date          string `json:"date"`
			VisitorVolume int    `json:"visitorVolume"`
		}

		result := make(map[string]int)
		for _, log := range logs {
			date := log.CreateTime.Format("2006-01-02")
			result[date]++
		}
		if len(result) < 365 {
			var startDate time.Time = time.Now().AddDate(-1, 0, 0)
			var endDate time.Time = time.Now()
			dates := generateDateList(startDate, endDate)
			for _, date := range dates {
				dateStr := date.Format("2006-01-02")
				if _, ok := result[dateStr]; !ok {
					result[dateStr] = 0
				}
			}
		}
		sortedDates := make([]string, 0, len(result))
		for date := range result {
			sortedDates = append(sortedDates, date)
		}
		sort.Strings(sortedDates)
		var jsonData []Data
		for _, date := range sortedDates {
			jsonData = append(jsonData, Data{
				Date:          date,
				VisitorVolume: result[date],
			})
		}

		response.Success(c, jsonData)
	})
}

func GetDailyBannedCount(c *gin.Context) {
	sql := `SELECT * FROM (
                  SELECT *,
                         (@row_number := IF(@prev_date = DATE(create_time) AND @prev_ip = ip, @row_number + 1, 1)) AS row_num,
                         (@prev_date := DATE(create_time)) AS dummy_date,
                         (@prev_ip := ip) AS dummy_ip
                  FROM log
                           CROSS JOIN (SELECT @row_number := 0, @prev_date := NULL, @prev_ip := NULL) AS vars
                  WHERE status = 403
                    AND create_time >= DATE_SUB(NOW(), INTERVAL 1 YEAR)
                  ORDER BY ip, create_time
              			) AS subquery WHERE row_num = 1 order by create_time;`

	preprocess(c, nil, func(db *gorm.DB) {
		var logs []middleware.Log
		db.Raw(sql).Scan(&logs)

		type Data struct {
			Date        string `json:"date"`
			BannedCount int    `json:"bannedCount"`
		}

		result := make(map[string]int)
		for _, log := range logs {
			date := log.CreateTime.Format("2006-01-02")
			result[date]++
		}
		if len(result) < 365 {
			var startDate = time.Now().AddDate(-1, 0, 0)
			var endDate = time.Now()
			dates := generateDateList(startDate, endDate)
			for _, date := range dates {
				dateStr := date.Format("2006-01-02")
				if _, ok := result[dateStr]; !ok {
					result[dateStr] = 0
				}
			}
		}
		sortedDates := make([]string, 0, len(result))
		for date := range result {
			sortedDates = append(sortedDates, date)
		}
		sort.Strings(sortedDates)
		var jsonData []Data
		for _, date := range sortedDates {
			jsonData = append(jsonData, Data{
				Date:        date,
				BannedCount: result[date],
			})
		}

		response.Success(c, jsonData)
	})
}

func GetPopularArticles(c *gin.Context) {
	preprocess(c, nil, func(db *gorm.DB) {
		articles := &[]struct {
			Id           int       `json:"id"`
			Title        string    `json:"title"`
			CreateTime   time.Time `json:"createTime"`
			ReadCount    int       `json:"readCount"`
			CommentCount int       `json:"commentCount"`
		}{}
		var total int64
		db.Model(Article{}).Count(&total).Find(articles)
		data := gin.H{
			"total": total,
			"list":  articles,
		}
		response.Success(c, data)
	})
}

func generateDateList(startDate, endDate time.Time) []time.Time {
	var dates []time.Time
	for d := startDate; d.Before(endDate); d = d.AddDate(0, 0, 1) {
		dates = append(dates, d)
	}
	return dates
}
