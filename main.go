package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	foo           string
	bar           int
	port          string
	my_test_token string
)

func main() {
	err := loadConfig()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	if port == "" {
		port = "8000"
	}

	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, Response{
			Message: "OK",
			Data:    nil,
		})
	})

	go debugConfigField()

	err = router.Run(":" + port)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func loadConfig() error {
	foo = os.Getenv("foo")
	bar, _ = strconv.Atoi(os.Getenv("bar"))
	my_test_token = os.Getenv("my_test_token")

	return nil
}

func debugConfigField() {
	for {
		time.Sleep(30 * time.Second)
		log.Println("my_test_token: " + my_test_token)
	}
}
