package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetForm(c *gin.Context) {
	config := Config()
	c.JSON(http.StatusOK, config)
}

func PostForm(c *gin.Context) {
	config := Config()
	err := c.ShouldBind(&config)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	UpdateConfig(config)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"code":    1,
	})
}

func GetSub(c *gin.Context) {
	config := Config()
	key := c.Query("i")
	if key == config.SubUrlShort {
		c.Redirect(301, config.SubUrl)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "error",
			"code":    0,
		})
	}
}
