package ginitem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"todo-api/common"
	"todo-api/modules/item/biz"
	"todo-api/modules/item/model"
	"todo-api/modules/item/storage"
)

func UpdateTodoById(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var item model.TodoUpdate
		// PARSED ID AND BODY
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// END PARSED ID AND BODY
		// CALL BIZ
		store := storage.NewMySqlStore(db)
		business := biz.NewUpdateTodoByIdBiz(store)

		if err := business.UpdateTodoById(c.Request.Context(), id, &item); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(id))
	}

}
