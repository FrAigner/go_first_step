package main

import (
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

var Status string = "green"

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/ping", ping)
	r.GET("/example/healthiness", getExampleHealthiness)
	r.POST("/example/healthiness", postExampleHealthiness)
	r.Run("localhost:8080")
	// Definieren Sie den Swagger-Endpunkt
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}

func ping(c *gin.Context) {
	c.PureJSON(200, gin.H{
		"message": "pong",
	})
}

func getExampleHealthiness(c *gin.Context) {
	type HINTS struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	var hints = []HINTS{
		{
			Key:   "Bezeichnung der Informationen",
			Value: "weitere informationen",
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
