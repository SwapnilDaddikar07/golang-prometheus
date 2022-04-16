package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-prometheus/controllers"
	"go-prometheus/middleware"
)

func main() {
	r := gin.Default()

	//add top level monitoring middleware and add endpoints that need to be excluded from metrics collection.
	r.Use(middleware.MonitorMetrics([]string{"/healthz","/metrics"}))

	responseSimulator := controllers.ResponseSimulator{}

	//register /metrics endpoint to get metrics in prometheus format. Add promhttp handler.
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	//expose some other urls to which simulate different behaviours.
	//these urls will be included in monitoring.
	r.GET("/all-success-simulator", responseSimulator.AllSuccessSimulator)
	r.GET("/internal-server-error-simulator", responseSimulator.InternalServerErrorSimulator)
	r.GET("/slow-simulator", responseSimulator.SlowResponseSimulator)

	//healthz endpoint configured to check exclusion from monitoring.
	r.GET("/healthz", func(ctx *gin.Context) {
		fmt.Println("service is live.")
	})

	r.Run(":8080")
}
