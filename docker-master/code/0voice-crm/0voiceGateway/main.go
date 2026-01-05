package main

import (
	"0voiceGateway/config"
	order2 "0voiceGateway/order"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
)

func main() {
	r := gin.Default()
	initRouter(r)
	r.Run(":8081")
}

func initRouter(r *gin.Engine) {
	conf := config.GetConf()
	apiGroup := r.Group("/api")
	order := order2.NewOrderInstance(conf)
	apiGroup.GET("/order", order.GetOrderDetail)

	apiGroup.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	apiGroup.GET("/ping", func(c *gin.Context) {
		ip := getIP()
		c.JSON(http.StatusOK, gin.H{
			"ip": ip,
		})
	})
}

func getIP() string {
	var ip string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println(err)
		return ""
	}
	for _, v := range addrs {
		if ipnet, ok := v.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}
		}
	}
	return ip
}