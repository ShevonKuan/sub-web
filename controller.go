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
		bResp := callBHandler(c, config.SubUrl)
		c.DataFromReader(bResp.StatusCode, bResp.ContentLength, bResp.Header.Get("Content-Type"), bResp.Body, nil)
		bResp.Body.Close()
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "error",
			"code":    0,
		})
	}
}

func callBHandler(c *gin.Context, bURL string) *http.Response {
	// Create a new request with the same method and body as the original request
	bReq, _ := http.NewRequest("GET", bURL, nil)
	bReq.Header = c.Request.Header

	// Send the request to the application and return the response
	client := http.DefaultClient
	bResp, _ := client.Do(bReq)
	return bResp
}
