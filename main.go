package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChatMsgRequest struct {
	Prompt string `json:"prompt"`
	Data   string `json:"data"`
}

type ChatMsgResponse struct {
	Data ChatMedia `json:"data"`
	Msg  string    `json:"msg"`
}

type ChatMedia struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func main() {
	r := gin.Default()
	r.POST("/handle-prompt", func(c *gin.Context) {
		var json ChatMsgRequest
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.Prompt != "echo" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid command"})
			return
		}
		var resp ChatMsgResponse
		resp.Msg = json.Data
		c.JSON(http.StatusOK, resp)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
