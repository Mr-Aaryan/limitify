package limitify

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type RateLimiter struct {
	ApiKey    string
	ServerURL string
}

type RateLimitRequest struct {
	APIKey    string `json:"api_key"`
	Path      string `json:"path"`
	Timestamp string `json:"timestamp"`
	Method    string `json:"method"`
	IP        string `json:"ip"`
	Country   string `json:"country_code"`
}

func NewRateLimiter(apiKey string) *RateLimiter {
	serverURL := "http://localhost:5000/rate-limit"
	return &RateLimiter{ApiKey: apiKey, ServerURL: serverURL}
}

func (rl *RateLimiter) CheckLimit(r *http.Request) (int, map[string]interface{}) {
	ip := GetClientIP(r)
	method := GetRequestMethod(r)
	path := GetRequestPath(r)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	country := GetCountry(ip)

	payload := RateLimitRequest{
		APIKey:    rl.ApiKey,
		Path:      path,
		Timestamp: timestamp,
		Method:    method,
		IP:        ip,
		Country:   country,
	}

	jsonData, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", rl.ServerURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return 500, map[string]interface{}{"error": err.Error()}
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return 500, map[string]interface{}{"error": err.Error()}
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return resp.StatusCode, result
}
