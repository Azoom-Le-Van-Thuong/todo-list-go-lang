package ginitem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"todo-api/common"
	"todo-api/modules/item/biz"
	"todo-api/modules/item/model"
	"todo-api/modules/item/storage"
)

func CreateTodo(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var item model.TodoCreation

		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewMySqlStore(db)
		business := biz.NewCreateTodoBiz(store)

		if err := business.CreateNewTodo(c.Request.Context(), &item); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(item.Id))
	}
}
