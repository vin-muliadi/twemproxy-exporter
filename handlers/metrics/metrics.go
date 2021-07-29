package metrics

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
)

type TwemproxyStats struct {
	Service          string `json:"service"`
	Source           string `json:"source"`
	Version          string `json:"version"`
	Uptime           int    `json:"uptime"`
	Timestamp        int    `json:"timestamp"`
	TotalConnections int    `json:"total_connections"`
	CurrConnections  int    `json:"curr_connections"`
	Shared           struct {
		ClientEOF         int `json:"client_eof"`
		ClientErr         int `json:"client_err"`
		ClientConnections int `json:"client_connections"`
		ServerEjects      int `json:"server_ejects"`
		ForwardError      int `json:"forward_error"`
		Fragments         int `json:"fragments"`
	} `json:"shared"`
}

func Redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/metrics")
}

func handleError(c *gin.Context, errorString string) {
	c.JSON(500, gin.H{"message": errorString, "status": "handle_error"})
}

func Metrics(c *gin.Context) {
	var metricString bytes.Buffer
	twemproxyServerHost := os.Getenv("TWEMPROXY_TARGET_HOST")
	if (twemproxyServerHost == "") {
		twemproxyServerHost = "127.0.0.1"
	}

	twemproxyServerPort := os.Getenv("TWEMPROXY_TARGET_PORT")
	if (twemproxyServerPort == "") {
		twemproxyServerPort = "22222"
	}

	conn, _ := net.Dial("tcp", twemproxyServerHost + ":" + twemproxyServerPort)
	message, _ := bufio.NewReader(conn).ReadString('\n')
	stats := TwemproxyStats{}
	jsonErr := json.Unmarshal([]byte(message), &stats)
	if jsonErr != nil {
		handleError(c, "JSON decode failed")
	}
	metricString.WriteString(fmt.Sprintf("# HELP twemproxy_uptime\n# TYPE twemproxy_uptime counter\ntwemproxy_uptime %d\n", stats.Uptime))
	metricString.WriteString(fmt.Sprintf("# HELP twemproxy_total_connections\n# TYPE twemproxy_total_connections counter\ntwemproxy_total_connections %d\n", stats.TotalConnections))
	metricString.WriteString(fmt.Sprintf("# HELP twemproxy_current_connections\n# TYPE twemproxy_current_connections gauge\ntwemproxy_current_connections %d\n", stats.CurrConnections))
	metricString.WriteString(fmt.Sprintf("# HELP twemproxy_client_eof\n# TYPE twemproxy_client_eof counter\ntwemproxy_client_eof %d\n", stats.Shared.ClientEOF))
	metricString.WriteString(fmt.Sprintf("# HELP twemproxy_client_err\n# TYPE twemproxy_client_err counter\ntwemproxy_client_err %d\n", stats.Shared.ClientErr))
	metricString.WriteString(fmt.Sprintf("# HELP twemproxy_client_client_connections\n# TYPE twemproxy_client_client_connections counter\ntwemproxy_client_client_connections %d\n", stats.Shared.ClientConnections))
	metricString.WriteString(fmt.Sprintf("# HELP twemproxy_server_ejects\n# TYPE twemproxy_server_ejects counter\ntwemproxy_server_ejects %d\n", stats.Shared.ServerEjects))
	metricString.WriteString(fmt.Sprintf("# HELP twemproxy_forward_error\n# TYPE twemproxy_forward_error counter\ntwemproxy_forward_error %d\n", stats.Shared.ForwardError))
	metricString.WriteString(fmt.Sprintf("# HELP twemproxy_fragments\n# TYPE twemproxy_fragments counter\ntwemproxy_fragments %d\n", stats.Shared.Fragments))
	c.Data(200, "application/json; charset=utf-8", []byte(metricString.String()))
}
