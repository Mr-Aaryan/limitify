package limitify

import (
	"net"
	"net/http"
	"strings"

	"github.com/oschwald/geoip2-golang"
)

func GetClientIP(r *http.Request) string {
	xff := r.Header.Get("X-Forwader-For")
	if xff != "" {
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0])
	}
	if realIp := r.Header.Get("X-Real-IP"); realIp != "" {
		return realIp
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "127.0.0.1"
	}
	return ip
}

func GetRequestMethod(r *http.Request) string {
	return r.Method
}

func GetRequestPath(r *http.Request) string {
	return r.URL.Path
}

func GetCountry(ip string) string {
	db, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		return "NA"
	}
	defer db.Close()

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return "NA"
	}
	record, err := db.Country(parsedIP)
	if err != nil || record.Country.Names["en"] == "" {
		return "NA"
	}
	return record.Country.Names["en"]
}
