package main

import (
	"twemproxy-exporter/handlers/healthcheck"
	"twemproxy-exporter/handlers/metrics"
	"github.com/gin-gonic/gin"
	"github.com/shokunin/contrib/ginrus"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	router := gin.New()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	router.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true, "twemproxy-exporter"))

	// Start routes
	router.GET("/health", healthcheck.HealthCheck)
	router.GET("/", metrics.Redirect)
	//router.GET("/metrics", metrics.Metrics)

	// RUN rabit run
	router.Run("0.0.0.0:9119") // listen and serve on 0.0.0.0:8080
}
