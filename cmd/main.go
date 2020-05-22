package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/TorchPing/go-torch/pkg/ping"
	"github.com/TorchPing/go-torch/pkg/resolve"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	version = "dev"
)

func routePing(c *gin.Context) {
	host := c.Param("host")
	port := c.Param("port")
	newPort, err := strconv.Atoi(port)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Port parse error",
		})
		return
	}
	target := ping.Target{
		Host:     host,
		Port:     uint16(newPort),
		Counter:  3,
		Interval: time.Second,
		Timeout:  time.Second * 3,
	}

	pinger := ping.NewPing()
	pinger.SetTarget(&target)

	pingerDone := pinger.Start()

	select {
	case <-pingerDone:
		break
	}
	result := pinger.Result()
	var resTime float64

	if result.SuccessCounter == 0 {
		resTime = 0
	} else {
		resTime = float64(result.TotalDuration) / float64(time.Millisecond) / float64(result.SuccessCounter)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": result.SuccessCounter > 0,
		"time":   resTime,
	})
}

func routeResolve(c *gin.Context) {
	host := c.Param("host")
	target := resolve.Target{
		Host:     host,
		Counter:  3,
		Interval: time.Second,
		Timeout:  time.Second * 3,
	}

	resolver := resolve.NewResolve()
	resolver.SetTarget(&target)

	pingerDone := resolver.Start()

	select {
	case <-pingerDone:
		break
	}
	result := resolver.Result()
	var resTime float64

	if result.SuccessCounter == 0 {
		resTime = 0
	} else {
		resTime = float64(result.TotalDuration) / float64(time.Millisecond) / float64(result.SuccessCounter)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": result.SuccessCounter > 0,
		"time":   resTime,
		"result": result.Addrs,
	})
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Meow~",
			"version": "Golang Edition",
		})
	})

	router.GET("/ping/:host/:port", routePing)

	router.GET("/resolve/:host", routeResolve)

	router.Run()
}
