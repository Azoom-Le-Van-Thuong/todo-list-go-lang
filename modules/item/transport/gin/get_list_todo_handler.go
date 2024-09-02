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

func GetListTodo(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		st := storage.NewMySqlStore(db)
		business := biz.NewGetListTodoBiz(st)
		var queryString struct {
			common.Paging
			model.Filter
		}

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		queryString.Paging.Process()
		todos, err := business.ListTodo(c.Request.Context(), &queryString.Filter, &queryString.Paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(todos, queryString.Filter, queryString.Paging))
	}
}
