package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"os"
	"os/exec"
	"pvgo/internal/config"
	"pvgo/internal/controller"
	modbusclient "pvgo/internal/modbus"
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

var myConfig config.Config
var db *pgxpool.Pool

func main() {

	config.ReadFile(&myConfig, "config.yaml")

	log.Println("Debug mode: ", myConfig.Debug)

	if myConfig.Database.RunLiquibase {
		err := os.Chdir("liquibase")
		if err != nil {
			log.Println("Unable to change directory to liquibase:", err)
			return
		}

		cmd := exec.Command("liquibase", "update")
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Println("\n\b", string(output))
			return
		}
		log.Println(string(output))
		err = os.Chdir("..")
		if err != nil {
			log.Println(err)
			return
		}
	}

	if !myConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	db := controller.ConnectDB()

	if db == nil {
		panic("Error opening DB")
	}

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
	defer db.Close()
}
