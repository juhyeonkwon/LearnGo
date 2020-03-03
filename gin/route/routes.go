package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello!",
	})
}

func Api1(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "api1!",
	})
}

func GetParam(c *gin.Context) {
	param := c.Param("name")
	c.String(http.StatusOK, "Hello %s", param)
}

func QueryStr(c *gin.Context) {
	first := c.DefaultQuery("first", "Guest")
	last := c.Query("last")

	c.String(http.StatusOK, "hello %s %s", first, last)

}
