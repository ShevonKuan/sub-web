package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetForm(c *gin.Context) {
	config := Config() // Get config data
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

type nameJSON struct {
	Name string `json:"name"`
}

func AddForm(c *gin.Context) {
	// get query string
	name := nameJSON{}
	err := c.ShouldBind(&name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Print(name)
	config := Config()
	newuuid := uuid.New().String()
	*config = append(*config,
		ConfigData{
			UUID: newuuid,
			Name: name.Name,
			Config: ConfigForm{
				Emoji:   true,
				Udp:     true,
				Scv:     true,
				NewName: true,
			},
		},
	)
	UpdateConfig(config)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"code":    1,
	})
}

func GetSub(c *gin.Context) {
	config := Config()
	key := c.Query("i")
	for _, v := range *config {
		if key == v.Config.SubUrlShort {
			bResp := callBHandler(c, v.Config.SubUrl)
			c.DataFromReader(bResp.StatusCode, bResp.ContentLength, bResp.Header.Get("Content-Type"), bResp.Body, nil)
			bResp.Body.Close()

			return // exit function
		} else {
			continue
		}
	}
	// If the key is not found, return an error message
	c.JSON(http.StatusOK, gin.H{
		"message": "error",
		"code":    0,
	})

}

func callBHandler(c *gin.Context, bURL string) *http.Response {
	// Create a new request with the same method and body as the original request
	bReq, _ := http.NewRequest("GET", bURL, nil)
	// bReq.Header = c.Request.Header

	// Send the request to the application and return the response
	client := http.DefaultClient
	bResp, _ := client.Do(bReq)
	return bResp
}
