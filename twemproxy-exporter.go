package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shokunin/contrib/ginrus"
	"github.com/sirupsen/logrus"
	"time"
	"twemproxy-exporter/handlers/healthcheck"
	"twemproxy-exporter/handlers/metrics"
	"os"
)

func main() {
	router := gin.New()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	router.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true, "twemproxy-exporter"))

	// Start routes
	router.GET("/health", healthcheck.HealthCheck)
	router.GET("/", metrics.Redirect)
	router.GET("/metrics", metrics.Metrics)

	twemproxyExporterPort := os.Getenv("TWEMPROXY_EXPORTER_PORT")
	if (twemproxyExporterPort == "") {
		twemproxyExporterPort = "9119"
	}

	// RUN rabit run
	router.Run("0.0.0.0:" + twemproxyExporterPort) // listen and serve on 0.0.0.0:8080
}
