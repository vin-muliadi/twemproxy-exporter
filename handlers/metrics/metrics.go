package metrics

import (
	//"errors"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/metrics")
}

//func handleError(c *gin.Context, errorString error) {
//	c.JSON(500, gin.H{"message": errorString, "status": "error"})
//}
