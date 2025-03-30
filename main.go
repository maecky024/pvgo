package main

import (
	modbusclient "KostalReader/modbus"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type JsonResponse struct {
	Code    int32  `json:"id"`
	Message string `json:"message"`
	Url     string `json:"url"`
}

func readKostal() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		modbusclient.ReadModbusStatus(c)
	}
	return fn
}

func GetError(c *gin.Context) {

	resp := JsonResponse{
		Code:    0,
		Message: "Forbidden",
		Url:     c.Request.RequestURI}

	c.JSON(http.StatusForbidden, resp)
}

func main() {

	router := gin.Default()
	defaultConfig := cors.DefaultConfig()
	defaultConfig.AllowAllOrigins = true
	defaultConfig.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	defaultConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	defaultConfig.ExposeHeaders = []string{"Content-Length"}
	defaultConfig.AllowCredentials = true
	defaultConfig.MaxAge = 12 * time.Hour

	router.Use(cors.New(defaultConfig))

	router.GET("/readKostal", readKostal())

	router.NoRoute(GetError)

	err := router.Run("localhost:7777")
	if err != nil {
		return
	}
	fmt.Println("Shutdown")
}
