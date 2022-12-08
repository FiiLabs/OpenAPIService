package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AccountHandler(c *gin.Context)  {
	name := c.PostForm("name")
	operation_id := c.PostForm("operation_id")

	fmt.Printf("name: %s; operation_id: %s", name, operation_id)
	datax := map[string]interface{}{
		"account": "bar",
		"name": "string",
		"operation_id": "string",
	}
	data := map[string]interface{}{
		"data": datax,
	}
	c.JSONP(http.StatusOK, data)
}