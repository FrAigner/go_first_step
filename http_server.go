package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/ping", ping)
	r.GET("/example/healthiness", exampleHealthiness)
	r.Run("localhost:8080")
}

func ping(c *gin.Context) {
	c.PureJSON(200, gin.H{
		"message": "pong",
	})
}

func exampleHealthiness(c *gin.Context) {
	type hint struct {
		Name_value string `json:"name_value"`
		Key_value  string `json:"key_value"`
	}
	var hints = []hint{
		{
			Name_value: "Bezeichnung der Informationen",
			Key_value:  "weitere informationen",
		},
	}
	var status string = "green"
	
	c.PureJSON(200, gin.H{
		"status":    status,
		"timestamp": time.Now().Format(time.RFC3339),
		"hints":     hints,
	})
}
