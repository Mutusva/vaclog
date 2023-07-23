package router

import (
	_ "Codenotary/docs/swag"
	"Codenotary/internal/api"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(c api.VacRecordController) {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	v1 := router.Group("/api/v1")
	v1.GET("/records", c.GetAllRecords)
	v1.PUT("/records", c.CreateRecord)
	v1.GET("/records/:documentID", c.SearchRecord)

	url := ginSwagger.URL(fmt.Sprintf("%s:%s", c.Config.Host, c.Config.Port) + "/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run(fmt.Sprintf("%s:%s", c.Config.Host, c.Config.Port))
}
