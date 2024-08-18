package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Todo struct {
	Id          int        `json:"id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Status      string     `json:"status,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
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

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		items := v1.Group("/items")
		items.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "GET items",
			})
		})
		items.POST("/", func(c *gin.Context) {})
		items.GET("/:id", func(c *gin.Context) {})
		items.PATCH("/:id", func(c *gin.Context) {})
		items.DELETE("/:id", func(c *gin.Context) {})

	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
