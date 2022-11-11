package helpers

import (
	"net"
	"net/http"
)

func ClientIP(req *http.Request) string {
	clientIP := req.RemoteAddr
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		clientIP = ip
	} else if ip = req.Header.Get("X-Forward-For"); ip != "" {
		clientIP = ip
	} else {
		clientIP, _, _ = net.SplitHostPort(clientIP)
	}

	if clientIP == "::1" {
		clientIP = "127.0.0.1"
	}

	return clientIP
}
