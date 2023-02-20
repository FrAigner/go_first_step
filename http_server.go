package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

var Status string = "green"

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/ping", ping)
	r.GET("/example/healthiness", getExampleHealthiness)
	r.POST("/example/healthiness", postExampleHealthiness)
	r.Run("localhost:8080")
}

func ping(c *gin.Context) {
	c.PureJSON(200, gin.H{
		"message": "pong",
	})
}

func getExampleHealthiness(c *gin.Context) {
	type HINTS struct {
		Key string `json:"key"`
		Value  string `json:"value"`
	}
	var hints = []HINTS{
		{
			Key: "Bezeichnung der Informationen",
			Value:  "weitere informationen",
		},
	}

	c.PureJSON(200, gin.H{
		"status":    Status,
		"timestamp": time.Now().Format(time.RFC3339),
		"hints":     hints,
	})
}

func postExampleHealthiness(c *gin.Context) {
	type STATUS struct {
		HEALTHINESSSTATUS string `json:"healthinessStatus"`
	}

	var status STATUS
	c.BindJSON(&status)
	c.JSON(200, gin.H{"status": status.HEALTHINESSSTATUS})
}
