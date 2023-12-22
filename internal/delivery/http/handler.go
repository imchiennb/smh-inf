package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/imchiennb/acmex/internal/app/model"
	"gorm.io/gorm"
)

func StartWebServer(db *gorm.DB) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.OPTIONS("/", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost")
		c.Header("Access-Control-Allow-Methods", "PUT, PATCH, GET, DELETE, POST")
		c.Header("Access-Control-Allow-Headers", "Origin")
		c.Header("Access-Control-Expose-Headers", "Content-Length")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "43200") // 12 hours

		c.Status(http.StatusOK)
	})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge:                    12 * time.Hour,
		OptionsResponseStatusCode: http.StatusOK,
	}))

	r.SetTrustedProxies([]string{"localhost"})

	r.GET("/entities", func(ctx *gin.Context) {
		query := struct {
			Size int    `json:"size" form:"size" binding:"required"`
			Page int    `json:"page" form:"page" binding:"required"`
			Name string `json:"name" form:"name"`
		}{}

		if err := ctx.ShouldBindQuery(&query); err != nil {
			query.Page = 1
			query.Size = 10
		}

		entities := []model.Entity{}

		db := db.Limit(query.Size).Offset((query.Page - 1) * query.Size).Order("created_at desc")

		fmt.Println("query:::", query)

		if query.Name != "" {
			db.Where("name like ?", "%"+query.Name+"%")
		}

		db.Find(&entities)

		ctx.JSON(200, entities)
	})

	r.GET("/entities/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		entity := model.Entity{}
		db.First(&entity, id)
		if entity.ID == 0 {
			ctx.JSON(404, gin.H{
				"message": "not found",
			})
			return
		}
		ctx.JSON(200, entity)
	})

	r.POST("/entities", func(ctx *gin.Context) {
		json := model.Entity{}
		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
		}

		body := model.Entity{
			Name: json.Name,
			Data: json.Data,
		}

		db.Create(&body)

		ctx.JSON(200, body)
	})

	r.PUT("/entities/:id", func(ctx *gin.Context) {
		json := model.Entity{}
		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
		}

		body := model.Entity{
			Name: json.Name,
			Data: json.Data,
		}

		id := ctx.Param("id")

		if affected := db.Model(&body).Where("id = ?", id).Updates(body).RowsAffected; affected == 0 {
			ctx.JSON(404, gin.H{
				"message": "not found",
			})
			return
		}
		ctx.JSON(200, gin.H{
			"message": "updated",
		})
	})

	r.DELETE("/entities/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		entity := model.Entity{}
		db.Delete(&entity, id)
		ctx.JSON(200, gin.H{
			"message": "deleted",
		})
	})

	r.Run(":8080")
}
