package iputil

import (
	"net"
	"net/http"
)

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
	XClientIP     = "x-client-ip"
)

func GetLocalIP() string {
	addrs, _ := net.InterfaceAddrs()

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok {
			if ipv4 := ipnet.IP.To4(); ipv4 != nil && !ipv4.IsLoopback() {
				return ipv4.String()
			}
		}
	}

	return "127.0.0.1"
}

func RemoteIP(req *http.Request) string {
	remote := req.RemoteAddr

	if r := req.Header.Get(XForwardedFor); r != "" {
		return r
	} else if r := req.Header.Get(XRealIP); r != "" {
		return r
	} else if r := req.Header.Get(XClientIP); r != "" {
		return r
	} else {
		remote, _, _ = net.SplitHostPort(remote)
	}

	if remote == "::1" {
		return "127.0.0.1"
	}

	return remote
}
