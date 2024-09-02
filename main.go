package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
	"todo-api/common"
	ginitem "todo-api/modules/item/transport/gin"
)

type Todo struct {
	common.SQLModel
	Title       string `json:"title,omitempty" gorm:"column:title"`
	Description string `json:"description,omitempty" gorm:"column:description"`
	Status      string `json:"status,omitempty" gorm:"column:status"`
}

func (Todo) TableName() string {
	return "todo"
}

func main() {
	//now := time.Now().UTC()
	//item := Todo{Id: 1, Title: "Task 1", Description: "Description 1", Status: "Pending", CreatedAt: &now, UpdatedAt: &now}
	//itemJson, err := json.Marshal(item)
	//if err != nil {
	//	log.Fatalln(err)
	//	return
	//}
	//log.Println(string(itemJson))
	//itemString := "{\"id\":1,\"title\":\"Task 1\",\"description\":\"Description 1\",\"status\":\"Pending\",\"created_at\":\"2024-08-18T09:15:35.116342Z\",\"updated_at\":\"2024-08-18T09:15:35.116342Z\"}"
	//
	//var item2 Todo
	//if err := json.Unmarshal([]byte(itemString), &item2); err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println(item2)

	dsn := os.Getenv("DB_CONNECTION")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = db.Debug()
	if err != nil {
		log.Fatalln("failed to connect database")
	}
	log.Println("Connected to database ... ", db)

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		items := v1.Group("/items")
		items.GET("/", ListItem(db))
		items.POST("/", ginitem.CreateTodo(db))
		items.GET("/:id", ginitem.GetTodo(db))
		items.PATCH("/:id", func(c *gin.Context) {})
		items.DELETE("/:id", DeleteItem(db))

	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:5005") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.

func DeleteItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var item Todo
		if err := db.Where("id = ?", id).First(&item, id).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return

		}

		if err := db.Table(Todo{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
			"status": "DELETED",
		}).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func ListItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		// should by parse paging
		var paging common.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		paging.Process()
		if err := db.Table(Todo{}.TableName()).Select("id").Count(&paging.Total).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}
		var items []Todo
		offset := (paging.Page - 1) * paging.Limit
		if err := db.
			Table(Todo{}.TableName()).
			Order("id desc").
			Offset(offset).Limit(paging.Limit).
			Find(&items).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, common.NewSuccessResponse(items, nil, paging))
	}
}
